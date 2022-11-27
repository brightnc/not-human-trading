package service

import (
	"fmt"
	"math"
	"time"

	"github.com/brightnc/not-human-trading/internal/core/domain"
	"github.com/brightnc/not-human-trading/pkg/logger"
	"github.com/cinar/indicator"
	"github.com/markcheno/go-talib"
)

func (svc *Service) StartBot(in domain.BotConfig) error {
	err := svc.botRepo.UpdateBotConfig(in)
	if err != nil {
		logger.Error("error while update UpdateBotConfig... ", err)
		return err
	}
	botConfig, err := svc.botRepo.RetrieveBotConfig()
	if err != nil {
		logger.Error("error while retrieving bot Config... ", err)
		return err
	}
	botExchange, err := svc.botRepo.RetrieveBotExchange()
	if err != nil {
		logger.Error("error while retrieving bot Exchange... ", err)
		return err
	}
	svc.hasStopSignal = false
	go svc.startBot(botConfig, botExchange)
	return err
}

func (svc *Service) StopBot() {
	logger.Info("got stoping signal")
	svc.hasStopSignal = true
}

func (svc *Service) startBot(botConfig domain.BotConfig, botExchange domain.BotExchange) {
	lastExecutionTime := time.Now()
	turnTimer := time.Now().Add(time.Second * 60)
	turnCounter := 0
	var placedQuantity string
	defer func() {
		svc.broacast(domain.WsMessage{
			Time:    time.Now(),
			Message: "bot has been stopped",
			Type:    domain.MessageTypeFeed,
		})
	}()
	for !svc.hasStopSignal {

		if time.Now().After(turnTimer) {
			turnCounter = 0
			logger.Info("Clear turn counter : ", turnCounter)
			turnTimer = time.Now().Add(time.Second * 60)
		}
		if turnCounter == 12 {
			logger.Info("maximun turn reached ")
			time.Sleep(time.Second)
			continue
		}
		// Cooldown excution process...
		if time.Since(lastExecutionTime) < time.Duration(time.Second*5) {
			time.Sleep(time.Millisecond * 500)
			continue
		}
		last1Month := time.Now().AddDate(0, -1, 0)
		goneDaysOfMonth := last1Month.Day()
		startDate := last1Month.AddDate(0, 0, -goneDaysOfMonth+1).Format("2006-01-02")
		endDate := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
		// TODO: retrive klines information...
		quote, err := svc.exchange.RetrieveKLines(
			botConfig.OrderConfig.Symbol,
			startDate,
			endDate,
			domain.Min5,
		)
		if err != nil {
			logger.Info("error while retrieving klines...")
			svc.hasStopSignal = true
		}
		// take an action buy/sell depend on svc.hasCreatedOrder
		hasSignal := false
		if botConfig.EMAConfig.IsActive {
			logger.Info("Checking for EMA Signal")
			svc.broacast(domain.WsMessage{
				Time:    time.Now(),
				Message: "Checking for EMA Signal...",
				Type:    domain.MessageTypeFeed,
			})
			emaFast := emaConfig{
				values:   quote.Close,
				period:   botConfig.EMAConfig.FastPeriod,
				priority: 1,
			}
			emaSlow := emaConfig{
				values:   quote.Close,
				period:   botConfig.EMAConfig.SlowPeriod,
				priority: 2,
			}
			hasSignal = svc.emaSignal(emaFast, emaSlow)
		}

		if botConfig.MACDConfig.IsActive {
			logger.Info("Checking for MACD Signal")
			svc.broacast(domain.WsMessage{
				Time:    time.Now(),
				Message: "Checking for MACD Signal...",
				Type:    domain.MessageTypeFeed,
			})
			macdConfig := macDConfig{
				values:       quote.Close,
				emaFast:      botConfig.MACDConfig.EMAFastPeriod,
				emaSlow:      botConfig.MACDConfig.EMASlowPeriod,
				signalPeriod: botConfig.MACDConfig.SignalPeriod,
			}
			hasSignal = svc.macdSignal(macdConfig)
		}
		if botConfig.RSIConfig.IsActive {
			logger.Info("Checking for RSI Signal")

			svc.broacast(domain.WsMessage{
				Time:    time.Now(),
				Message: "Checking for RSI Signal...",
				Type:    domain.MessageTypeFeed,
			})

			rsiConfig := RSIConfig{
				values: quote.Close,
				period: botConfig.RSIConfig.Period,
			}
			hasSignal = svc.rsiSignal(rsiConfig)
		}
		if botConfig.STOConfig.IsActive {
			logger.Info("Checking for STO Signal")

			svc.broacast(domain.WsMessage{
				Time:    time.Now(),
				Message: "Checking for STO Signal...",
				Type:    domain.MessageTypeFeed,
			})

			stoConfig := STOConfig{
				High:  quote.High,
				Low:   quote.Low,
				Close: quote.Close,
			}
			hasSignal = svc.stoSignal(stoConfig)
		}
		if botConfig.SupertrendConfig.IsActive {
			logger.Info("Checking Supertrend Signal")

			svc.broacast(domain.WsMessage{
				Time:    time.Now(),
				Message: "Checking for Supertrend Signal...",
				Type:    domain.MessageTypeFeed,
			})

			sptConfig := SPTConfig{
				High:       quote.High,
				Low:        quote.Low,
				Close:      quote.Close,
				ATRPeriod:  botConfig.SupertrendConfig.ATRPeriod,
				Multiplier: float64(botConfig.SupertrendConfig.Multiplier),
			}

			_, _, _, trend, _ := superTrendDetail(sptConfig)
			logger.Info("TREND >> ", trend[len(trend)-10:])
			hasSignal = svc.superTrendSignal(trend)
		}
		if hasSignal {
			order := domain.PlaceOrder{
				Symbol:    botConfig.OrderConfig.Symbol,
				Quantity:  botConfig.OrderConfig.Quantity,
				OrderType: domain.OrderTypeMarket,
			}
			exchangeKey := domain.BotExchange{
				APIKey:    botExchange.APIKey,
				SecretKey: botExchange.SecretKey,
			}
			if svc.hasCreatedOrder {
				//TODO: selling
				logger.Info("Selling... symbol: %s quantity:  %v @%v\n", botConfig.OrderConfig.Symbol, botConfig.OrderConfig.Quantity, time.Now().Format(time.RFC3339Nano))
				order.Quantity = placedQuantity
				result, err := svc.exchange.PlaceAsk(order, exchangeKey)
				if err != nil {
					logger.Error("cannot PlaceAsk order ", err)
					return
				}
				svc.broacast(placedOrderMessage(result))
				logger.Info("created sell order : ", result)
				svc.hasCreatedOrder = false
			} else {
				// TODO:: buying
				fmt.Printf("Buying... symbol: %s quantity:  %v @%s", botConfig.OrderConfig.Symbol, botConfig.OrderConfig.Quantity, time.Now().Format(time.RFC3339Nano))

				result, err := svc.exchange.PlaceBid(order, exchangeKey)
				if err != nil {
					logger.Error("cannot PlaceAsk order ", err)
					return
				}
				svc.broacast(placedOrderMessage(result))
				svc.broacast(placedOrderMessage(result))
				logger.Info("created buy order : ", result)
				placedQuantity = result.ActualQuantity
				svc.hasCreatedOrder = true
			}

		}
		lastExecutionTime = time.Now()
		turnCounter++
	}
	logger.Info("bot has been stopped")
}

func placedOrderMessage(result domain.PlaceOrderResult) domain.WsMessage {
	message := fmt.Sprintf(
		`symbol: %s, 
	origin_quantity: %s, 
	actual_quantity: %s, 
	side: %s, 
	order_type: %s, 
	price: %s`, result.Symbol, result.OriginQuantity, result.ActualQuantity, result.Side, result.OrderType, result.Price)

	return domain.WsMessage{
		Time:    time.Now(),
		Message: message,
		Type:    domain.MessageTypeTradingReport,
	}
}

func (svc *Service) broacast(message domain.WsMessage) {
	for i := range svc.subscribers {

		svc.subscribers[i].Message <- message
	}
}

func (svc *Service) RegisterWebsocketClient(conn *domain.Connection) error {
	svc.subscribers = append(svc.subscribers, conn)
	return nil
}

type emaConfig struct {
	values   []float64
	period   int
	priority int
}

// return sell signal if has created order
// return buy signal if not has created order
func (svc *Service) emaSignal(configs ...emaConfig) bool {
	type indicator struct {
		indicatorValue []float64
	}
	mapper := make(map[int]indicator, len(configs))
	isOK := false
	for i := range configs {
		ema := talib.Ema(configs[i].values, configs[i].period)
		mapper[configs[i].priority-1] = indicator{
			indicatorValue: ema,
		}
	}

	for i := range configs {
		if i == len(configs)-1 {
			break
		}
		logger.Info("EMA first priority = ", mapper[i].indicatorValue[len(mapper[i].indicatorValue)-1])
		logger.Info("EMA second priority = ", mapper[i+1].indicatorValue[len(mapper[i+1].indicatorValue)-1])

		if svc.hasCreatedOrder {
			// looking for sell signal
			isOK = talib.Crossunder(mapper[i].indicatorValue, mapper[i+1].indicatorValue)
		} else {
			// looking for buy signal
			isOK = talib.Crossover(mapper[i].indicatorValue, mapper[i+1].indicatorValue)
		}

	}

	return isOK
}

type macDConfig struct {
	values       []float64
	emaFast      int
	emaSlow      int
	signalPeriod int
}

func (svc *Service) macdSignal(configs macDConfig) bool {
	emaFast := ValuesToEMA(configs.values, configs.emaFast)
	emaSlow := ValuesToEMA(configs.values, configs.emaSlow)

	macD := make([]float64, len(emaSlow))
	// macD = append(macD, emaFast[len(emaFast)-1]-emaSlow[len(emaSlow)-1])
	// macD = append(macD, emaFast[len(emaFast)-2]-emaSlow[len(emaSlow)-2])

	for i := range macD {
		macD[i] = emaFast[i] - emaSlow[i]
	}
	emaSignal := ValuesToEMA(macD, configs.signalPeriod)
	lastValueOfMACD := macD[len(macD)-1]
	lastValueOfSignal := emaSignal[len(emaSignal)-1]
	isOK := false
	v := (lastValueOfMACD - lastValueOfSignal) / lastValueOfSignal
	logger.Info("Fast >>> ", len(emaFast), " Slow >>> ", len(emaSlow))
	logger.Info("current value >> ", lastValueOfMACD, " ", lastValueOfSignal, " Signal >>> ", emaSignal[len(emaSignal)-2:])
	if v >= -0.1 && v <= 0.1 {
		logger.Info("almost cross >> ", lastValueOfMACD, " ", lastValueOfSignal)
	}

	if svc.hasCreatedOrder {
		// looking for sell signal
		isOK = talib.Crossunder(macD, emaSignal)
	} else {
		// looking for buy signal
		isOK = talib.Crossover(macD, emaSignal)
	}

	return isOK
}
func ValuesToEMA(values []float64, period int) []float64 {
	return talib.Ema(values, period)
}

type RSIConfig struct {
	values []float64
	period int
}

func (svc *Service) rsiSignal(configs RSIConfig) bool {
	isOK := false
	rsi := talib.Rsi(configs.values, configs.period)
	logger.Info(rsi[len(rsi)-1])
	if svc.hasCreatedOrder {
		// looking for sell signal
		isOK = rsi[len(rsi)-1] > 70
	} else {
		// looking for buy signal
		isOK = rsi[len(rsi)-1] < 30
	}

	return isOK

}

type STOConfig struct {
	High  []float64
	Low   []float64
	Close []float64
}

func (svc *Service) stoSignal(configs STOConfig) bool {
	isOK := false
	k, d := indicator.StochasticOscillator(configs.High, configs.Low, configs.Close)
	fmt.Printf("k >>> %v  d >>> %v \n", k[len(k)-1], d[len(d)-1])
	if svc.hasCreatedOrder {
		// looking for sell signal
		isOK = talib.Crossunder(k, d) && k[len(k)-1] > 80 && d[len(d)-1] > 80
	} else {
		// looking for buy signal
		isOK = talib.Crossover(k, d) && k[len(k)-1] < 20 && d[len(d)-1] < 20
	}
	return isOK

}

type SPTConfig struct {
	High       []float64
	Low        []float64
	Close      []float64
	ATRPeriod  int
	Multiplier float64
}

func (svc *Service) superTrendSignal(trend []bool) bool {
	prev_trend := trend[len(trend)-2]
	current_trend := trend[len(trend)-1]

	isOK := false
	if svc.hasCreatedOrder {
		// looking for sell signal
		isOK = prev_trend && !current_trend
	} else {
		// looking for buy signal
		isOK = !prev_trend && current_trend
	}
	return isOK
}

func superTrendDetail(configs SPTConfig) ([]float64, []float64, []float64, []bool, []time.Time) {
	l := len(configs.High)
	hl2 := talib.MedPrice(configs.High, configs.Low)
	atr := talib.Atr(configs.High, configs.Low, configs.Close, configs.ATRPeriod)

	up := make([]float64, l)
	down := make([]float64, l)
	trendUp := make([]float64, l)
	trendDown := make([]float64, l)
	trend := make([]bool, l)
	tsl := make([]float64, l)
	times := make([]time.Time, 0)
	for i := 0; i < l; i++ {
		up[i] = hl2[i] - atr[i]*configs.Multiplier
		down[i] = hl2[i] + atr[i]*configs.Multiplier
		if i == 0 {
			trendUp[i] = up[i]
			trendDown[i] = down[i]
			trend[i] = true
			tsl[i] = trendUp[i]
			continue
		}
		if configs.Close[i-1] > trendUp[i-1] {
			trendUp[i] = math.Max(up[i], trendUp[i-1])
		} else {
			trendUp[i] = up[i]
		}
		if configs.Close[i-1] < trendDown[i-1] {
			trendDown[i] = math.Min(down[i], trendDown[i-1])
		} else {
			trendDown[i] = down[i]
		}

		if configs.Close[i] > trendDown[i-1] {

			trend[i] = true
		} else if configs.Close[i] < trendUp[i-1] {

			trend[i] = false
		} else {

			trend[i] = trend[i-1]
		}

		if trend[i] {
			tsl[i] = trendUp[i]
		} else {
			tsl[i] = trendDown[i]
		}
	}

	return trendUp, trendDown, tsl, trend, times
}
