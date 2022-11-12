package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/markcheno/go-quote"
)

var (
	openPosition = false
	createdOrder *futures.CreateOrderResponse
	// percentage
	expectedProfit = -0.0015
	// percentage
	cutLoss = -0.0015
)

func main() {
	fmt.Println(` __    _  _______  _______    __   __  __   __  __   __  _______  __    _                   
	|  |  | ||       ||       |  |  | |  ||  | |  ||  |_|  ||   _   ||  |  | |                  
	|   |_| ||   _   ||_     _|  |  |_|  ||  | |  ||       ||  |_|  ||   |_| |                  
	|       ||  | |  |  |   |    |       ||  |_|  ||       ||       ||       |                  
	|  _    ||  |_|  |  |   |    |       ||       ||       ||       ||  _    | ___   ___   ___  
	| | |   ||       |  |   |    |   _   ||       || ||_|| ||   _   || | |   ||   | |   | |   | 
	|_|  |__||_______|  |___|    |__| |__||_______||_|   |_||__| |__||_|  |__||___| |___| |___| 
	`)
	// staregyEMA("BTCUSDT", "0.001")
	// staregyMacD("BTCUSDT", "0.001")
	// strategyRSI("BTCUSDT", "0.001")
	// staregySTO("BTCUSDT", "0.001")
	superTrend("BTCUSDT", "0.001")
	// binanceKline()

	// err = godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// binanceTradeWs()

}

func convertUnixToDateTime(n int64) string {
	t := time.UnixMilli(n).UTC()

	return t.Format(time.RubyDate)
}

func convertDateTimeToUnix(d string) int64 {
	t, err := time.Parse(time.RubyDate, d)
	if err != nil {
		panic(err)
	}

	return t.Unix()
}

func binanceKline() {
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsKlineServe("BTCUSDT", "5m", wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

func binanceTradeWs() {

	var df dataframe.DataFrame
	//srs := series.Series{}
	isFirst := true
	columnSymbol := series.Series{}
	columnPrice := series.Series{}
	columnTime := series.Series{}
	columnUnixTime := series.Series{}
	currentTime := time.Now()
	f, err := os.Create("df.csv")
	if err != nil {
		panic(err)
	}
	wsTradeHandler := func(event *binance.WsTradeEvent) {
		now := time.Now()
		if time.Duration(currentTime.Second()) < time.Duration(now.Second()) {
			if isFirst {
				columnSymbol = series.New(event.Symbol, series.String, "Symbol")
				columnTime = series.New(convertUnixToDateTime(event.Time), series.String, "DateTime")
				columnUnixTime = series.New(strconv.FormatInt(event.Time, 16), series.String, "UnixTime")
				fmt.Println("event Time : ", event.Time)
				columnPrice = series.New(event.Price, series.Float, "Price")
				isFirst = false
			} else {
				columnSymbol.Append(event.Symbol)
				columnTime.Append(convertUnixToDateTime(event.Time))
				columnUnixTime.Append(strconv.FormatInt(event.Time, 16))
				columnPrice.Append(event.Price)
			}

			df = dataframe.New(columnSymbol, columnTime, columnUnixTime, columnPrice)

			staregy(-0.0000001, df)
			// fmt.Println(df)
			df.WriteCSV(f)

		}
		currentTime = now
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsTradeServe("ETHBUSD", wsTradeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}
func retrievePriceRecords(df dataframe.DataFrame) []float64 {
	return df.Col("Price").Float()
}

func strategyRSI(symbol, qty string) {
	fmt.Println("RSI staregy starting...")
	currentTime := time.Now()
	for {
		now := time.Now()
		if time.Duration(currentTime.Second()) < time.Duration(now.Second()) {
			last1Month := time.Now().AddDate(0, -1, 0)
			goneDaysOfMonth := last1Month.Day()

			start := last1Month.AddDate(0, 0, -goneDaysOfMonth+1).Format("2006-01-02")
			// start := time.Unix(time.Now().Unix()-int64(720*time.Hour), 0).Format("2006-01-02")
			end := time.Now().Format("2006-01-02")
			fmt.Println(start, end)
			spy, err := quote.NewQuoteFromBinance(symbol, start, end, quote.Min5)
			if err != nil {
				panic(err)
			}
			fmt.Println("took ", time.Since(now).Seconds(), "seconds to retrieve trading history")

			config := RSIConfig{
				values: spy.Close,
				period: 14,
			}

			if !openPosition {

				fmt.Println("buysignal>>>")
				if RSIBuySignal(config) {
					createOrder(symbol, futures.OrderTypeMarket, futures.SideTypeBuy, qty)
					openPosition = true
					fmt.Println("creteadBuyOrder ", createdOrder)

				}
			}

			if openPosition {

				fmt.Println("sellsignal>>>")
				if RSISellSignal(config) {
					createOrder(symbol, futures.OrderTypeMarket, futures.SideTypeSell, qty)
					openPosition = false
					fmt.Println("creteadSellOrder ", createdOrder)
				}
			}

		}
		currentTime = time.Now()
	}
}

func superTrend(symbol, qty string) {
	fmt.Println("SPT staregy starting...")
	currentTime := time.Now()
	defer recover()
	for {
		now := time.Now()
		if time.Duration(currentTime.Second()) < time.Duration(now.Second()) {
			last2Month := time.Now().AddDate(0, -2, 0)
			goneDaysOfMonth := last2Month.Day()

			start := last2Month.AddDate(0, 0, -goneDaysOfMonth+1).Format("2006-01-02")
			// start := time.Unix(time.Now().Unix()-int64(720*time.Hour), 0).Format("2006-01-02")
			end := time.Now().Format("2006-01-02")
			fmt.Println(start, end, symbol, quote.Min5)
			spy, err := quote.NewQuoteFromBinance(symbol, start, end, quote.Min5)
			if err != nil {
				panic(err)
			}
			fmt.Println("took ", time.Since(now).Seconds(), "seconds to retrieve trading history")
			fmt.Printf("Close >>> %v , H >>> %v , L >>> %v \n", spy.Close[len(spy.Close)-1], spy.High[len(spy.High)-1], spy.Low[len(spy.Low)-1])

			config := SPTConfig{
				High:       spy.High,
				Low:        spy.Low,
				Close:      spy.Close,
				ATRPeriod:  14,
				Multiplier: 3,
				Time:       spy.Date,
			}
			up, down, _, trend, times := SuperTrendDetail(config)

			prev_trend := trend[len(trend)-2]
			current_trend := trend[len(trend)-1]
			fmt.Println("----------------------------------------------------------------------")
			fmt.Printf("prev >> %v, current >> %v last100 >> %v\n", prev_trend, current_trend, trend[len(trend)-100:])
			fmt.Printf("time >>> %v\n", times[len(times)-100:])
			fmt.Printf("UP >> %v, DOWN >> %v\n", up[len(up)-1], down[len(down)-1])
			fmt.Println("----------------------------------------------------------------------")
			fmt.Println("configsignal>>>")
			if !openPosition {

				fmt.Println("buysignal>>>")

				if !prev_trend && current_trend {
					createOrder(symbol, futures.OrderTypeMarket, futures.SideTypeBuy, qty)
					openPosition = true
					fmt.Println("creteadBuyOrder ", createdOrder)
				}

			}

			if openPosition {

				fmt.Println("sellsignal>>>")
				if prev_trend && !current_trend {
					createOrder(symbol, futures.OrderTypeMarket, futures.SideTypeSell, qty)
					openPosition = false
					fmt.Println("creteadSellOrder ", createdOrder)
				}
			}
		}
		currentTime = time.Now()
	}
}

func staregySTO(symbol, qty string) {
	fmt.Println("STO staregy starting...")
	currentTime := time.Now()
	defer recover()
	for {
		now := time.Now()
		if time.Duration(currentTime.Second()) < time.Duration(now.Second()) {
			last2Month := time.Now().AddDate(0, -2, 0)
			goneDaysOfMonth := last2Month.Day()

			start := last2Month.AddDate(0, 0, -goneDaysOfMonth+1).Format("2006-01-02")
			// start := time.Unix(time.Now().Unix()-int64(720*time.Hour), 0).Format("2006-01-02")
			end := time.Now().Format("2006-01-02")
			fmt.Println(start, end, symbol, quote.Min5)
			spy, err := quote.NewQuoteFromBinance(symbol, start, end, quote.Min5)
			if err != nil {
				panic(err)
			}
			fmt.Println("took ", time.Since(now).Seconds(), "seconds to retrieve trading history")
			fmt.Printf("Close >>> %v , H >>> %v , L >>> %v \n", spy.Close[len(spy.Close)-1], spy.High[len(spy.High)-1], spy.Low[len(spy.Low)-1])

			config := STOConfig{
				High:  spy.High,
				Low:   spy.Low,
				Close: spy.Close,
			}
			fmt.Println("configsignal>>>")
			if !openPosition {

				fmt.Println("buysignal>>>")

				if STOBuySignal(config) {
					createOrder(symbol, futures.OrderTypeMarket, futures.SideTypeBuy, qty)
					openPosition = true
					fmt.Println("creteadBuyOrder ", createdOrder)

				}
			}

			if openPosition {

				fmt.Println("sellsignal>>>")
				if STOSellSignal(config) {
					createOrder(symbol, futures.OrderTypeMarket, futures.SideTypeSell, qty)
					openPosition = false
					fmt.Println("creteadSellOrder ", createdOrder)
				}
			}
		}
		currentTime = time.Now()
	}
}
func staregyMacD(symbol, qty string) {
	fmt.Println("MACD staregy starting...")
	currentTime := time.Now()
	defer recover()
	for {
		now := time.Now()
		if time.Duration(currentTime.Second()) < time.Duration(now.Second()) {
			last1Month := time.Now().AddDate(0, -1, 0)
			goneDaysOfMonth := last1Month.Day()

			start := last1Month.AddDate(0, 0, -goneDaysOfMonth+1).Format("2006-01-02")
			// start := time.Unix(time.Now().Unix()-int64(720*time.Hour), 0).Format("2006-01-02")
			end := time.Now().Format("2006-01-02")
			fmt.Println(start, end)
			spy, err := quote.NewQuoteFromBinance(symbol, start, end, quote.Min5)
			if err != nil {
				panic(err)
			}
			fmt.Println("took ", time.Since(now).Seconds(), "seconds to retrieve trading history")
			fmt.Printf("Close >>> %v , H >>> %v , L >>> %v \n", spy.Close[len(spy.Close)-1], spy.High[len(spy.High)-1], spy.Low[len(spy.Low)-1])

			config := macDConfig{
				values:       spy.Close,
				emaFast:      12,
				emaSlow:      26,
				signalPeriod: 9,
			}
			fmt.Println("configsignal>>>")
			if !openPosition {

				fmt.Println("buysignal>>>")
				if MACDBuySignal(config) {
					createOrder(symbol, futures.OrderTypeMarket, futures.SideTypeBuy, qty)
					openPosition = true
					fmt.Println("creteadBuyOrder ", createdOrder)

				}
			}

			if openPosition {

				fmt.Println("sellsignal>>>")
				if MACDSellSignal(config) {
					createOrder(symbol, futures.OrderTypeMarket, futures.SideTypeSell, qty)
					openPosition = false
					fmt.Println("creteadSellOrder ", createdOrder)
				}
			}
		}
		currentTime = time.Now()
	}

}

func staregyEMA(symbol, qty string) {
	fmt.Println("EMA staregy starting...")
	timeToStopProcess := time.Now().Add(time.Hour * 8)
	currentTime := time.Now()
	for {

		now := time.Now()
		if time.Duration(currentTime.Second()) < time.Duration(now.Second()) {
			fmt.Println("watching EMA ...")
			// stop process after 8 hours ago
			if time.Now().After(timeToStopProcess) {
				break
			}
			last1Month := time.Now().AddDate(0, -1, 0)
			goneDaysOfMonth := last1Month.Day()

			start := last1Month.AddDate(0, 0, -goneDaysOfMonth+1).Format("2006-01-02")
			end := time.Now().Format("2006-01-02")
			fmt.Println("period of information from ", start, "to ", end)
			spy, err := quote.NewQuoteFromBinance(symbol, start, end, quote.Min5)
			if err != nil {
				panic(err)
			}
			fmt.Println("took ", time.Since(now).Seconds(), "seconds to retrieve trading history")
			// waiting for buy signal
			if !openPosition {
				isOK := EMASignal(
					true,
					emaConfig{
						values:   spy.Close,
						period:   10,
						priority: 1,
					},
					emaConfig{
						values:   spy.Close,
						period:   20,
						priority: 2,
					},
				)
				if isOK {
					createOrder(symbol, futures.OrderTypeMarket, futures.SideTypeBuy, qty)
					openPosition = true
					fmt.Println("creteadBuyOrder ", createdOrder)
				}
			}

			// waiting for sell signal
			if openPosition {
				isOK := EMASignal(
					false,
					emaConfig{
						values:   spy.Close,
						period:   10,
						priority: 1,
					},
					emaConfig{
						values:   spy.Close,
						period:   20,
						priority: 2,
					},
				)
				if isOK {
					createOrder(symbol, futures.OrderTypeMarket, futures.SideTypeSell, qty)
					openPosition = false
					fmt.Println("creteadSellOrder ", createdOrder)
				}
			}
		}
		currentTime = now

	}
}

func staregy(entry float64, df dataframe.DataFrame) {
	targetSymbol := "ETHBUSD"
	qty := "0.01"
	log.Println("df view => ", df.String())
	result := calCumulative(df)
	if !openPosition {
		fmt.Println("result : ", result, " entry: ", entry)
		if result > entry {
			createOrder(targetSymbol, futures.OrderTypeMarket, futures.SideTypeBuy, qty)
			fmt.Println("creteadBuyOrder ", createdOrder)
			openPosition = true
		}
	}
	if openPosition {
		// startIndex := 0
		// endIndex := 0
		// timeRecords := df.Col("Time")
		// for i := range timeRecords.Records() {

		// }
		afterCreatedOrderPriceRecords := df.Filter(
			dataframe.F{
				Colidx:     2,
				Colname:    "UnixTime",
				Comparator: series.Greater,
				Comparando: strconv.FormatInt(createdOrder.UpdateTime, 16),
			},
		)
		// return if there are not any records after we created order
		if afterCreatedOrderPriceRecords.Error() != nil {
			fmt.Println("error afterCreatedOrderPriceRecords ", afterCreatedOrderPriceRecords.Error().Error())
			return
		}
		if len(afterCreatedOrderPriceRecords.Records()) <= 0 {
			fmt.Println("afterCreatedOrderPriceRecords is empty")
			return
		}
		if afterCreatedOrderPriceRecords.String() == "Empty DataFrame" {
			fmt.Println("afterCreatedOrderPriceRecords are empty ")
			return
		}
		fmt.Println("afterCreatedOrderPriceRecords => ", afterCreatedOrderPriceRecords.String())
		result = calCumulative(afterCreatedOrderPriceRecords)
		if result >= expectedProfit || result <= cutLoss {
			createOrder(targetSymbol, futures.OrderTypeMarket, futures.SideTypeSell, qty)
			fmt.Println("creteadSellOrder ", createdOrder)
			openPosition = false
		}
	}
}

func calCumulative(df dataframe.DataFrame) float64 {
	priceRecords := retrievePriceRecords(df)
	fPrice := priceRecords[0]
	lPrice := priceRecords[len(priceRecords)-1]
	return (lPrice / fPrice) - 1
}

func createOrder(symbol string, orderType futures.OrderType, side futures.SideType, qty string) {
	var err error
	futures.UseTestnet = true
	client := futures.NewClient("35ab0f236f179c4e3aa9126e9f6a5943fb7d3181117a23312c07d4d94ba14c56", "35f23c0014b0af922ba645d4efd6b8ce13003334ae6077e8342f7c317278cf4b")
	createdOrder, err = client.NewCreateOrderService().Symbol(symbol).Type(orderType).Side(side).Quantity(qty).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("created order : ", createdOrder)

}
