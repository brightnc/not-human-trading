package service

import (
	"fmt"
	"time"

	"github.com/brightnc/not-human-trading/internal/core/domain"
	"github.com/markcheno/go-talib"
)

func (svc *Service) StartBot(in domain.BotConfig) error {
	err := svc.botRepo.UpdateBotConfig(in)
	if err != nil {
		fmt.Println("error while update bot config...")
		return err
	}
	botConfig, err := svc.botRepo.RetrieveBotConfig()
	if err != nil {
		fmt.Println("error while retrieving bot config...")
		return err
	}
	svc.hasStopSignal = false
	go svc.startBot(botConfig)
	return err
}

func (svc *Service) StopBot() {
	fmt.Println("got stop bot signal")
	svc.hasStopSignal = true
	return
}

// func (svc *Service) botHandler(startSignal chan bool, botConfig domain.BotConfig) {
// 	for {
// 		select {
// 		case <-startSignal:
// 			go svc.startBot(botConfig)
// 		case <-svc.stopSignal:
// 			fmt.Println("bot stopping")
// 			return
// 		}
// 	}
// }

func (svc *Service) startBot(botConfig domain.BotConfig) {
	lastExecutionTime := time.Now()
	for svc.hasStopSignal {
		// Cooldown excution process...
		if time.Now().Sub(lastExecutionTime) < time.Duration(time.Second) {
			time.Sleep(time.Millisecond * 500)
			continue
		}
		last1Month := time.Now().AddDate(0, -1, 0)
		goneDaysOfMonth := last1Month.Day()
		startDate := last1Month.AddDate(0, 0, -goneDaysOfMonth+1).Format("2006-01-02")
		endDate := time.Now().Format("2006-01-02")
		// TODO: retrive klines information...
		quote, err := svc.exchange.RetrieveKLines(
			botConfig.OrderConfig.Symbol,
			startDate,
			endDate,
			domain.Min5,
		)
		if err != nil {
			fmt.Println("error while retrieving klines...")
			svc.stopSignal <- true
		}
		// take an action buy/sell depend on svc.hasCreatedOrder
		hasSignal := false
		if botConfig.EMAConfig.IsActive {
			fmt.Println("Checking for EMA Signal")
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
			fmt.Println("Checking for MACD Signal")
		}
		if botConfig.RSIConfig.IsActive {
			fmt.Println("Checking for RSI Signal")
		}
		if botConfig.STOConfig.IsActive {
			fmt.Println("Checking for STO Signal")
		}
		if botConfig.SupertrendConfig.IsActive {
			fmt.Println("Checking Supertrend Signal")
		}
		if hasSignal && svc.hasCreatedOrder {
			//TODO: selling
			fmt.Printf("Selling symbol: %s quantity:  %v @%v", botConfig.OrderConfig.Symbol, botConfig.OrderConfig.Quantity, time.Now().Format(time.RFC3339Nano))
			svc.hasCreatedOrder = false
		}
		if hasSignal && !svc.hasCreatedOrder {
			// TODO:: buying
			fmt.Printf("Buying symbol: %s quantity:  %v @%s", botConfig.OrderConfig.Symbol, botConfig.OrderConfig.Quantity, time.Now().Format(time.RFC3339Nano))
			svc.hasCreatedOrder = true
		}
		lastExecutionTime = time.Now()
	}
	fmt.Println("bot has been stopped")

}

type emaConfig struct {
	values   []float64
	period   int
	priority int
	//operaton bool
}

// return sell signal if has created order
// return buy signal if not has created order
func (svc *Service) emaSignal(configs ...emaConfig) bool {
	type indicator struct {
		indicatorValue float64
	}
	mapper := make(map[int]indicator, len(configs))
	isOK := false
	for i := range configs {
		ema := talib.Ema(configs[i].values, configs[i].period)
		mapper[configs[i].priority-1] = indicator{
			indicatorValue: ema[len(ema)-1],
		}
	}

	for i := range configs {
		if i == len(configs)-1 {
			break
		}
		fmt.Println("EMA first priority = ", mapper[i].indicatorValue)
		fmt.Println("EMA second priority = ", mapper[i+1].indicatorValue)

		if svc.hasCreatedOrder {
			// looking for sell signal
			isOK = mapper[i].indicatorValue < mapper[i+1].indicatorValue
		} else {
			// looking for buy signal
			isOK = mapper[i].indicatorValue > mapper[i+1].indicatorValue
		}

	}

	return isOK
}
