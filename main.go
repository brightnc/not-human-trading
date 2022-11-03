package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/joho/godotenv"
)

var openPosition = false

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	binanceTradeWs()

}

func convertUnixToDateTime(n int64) string {
	t := time.UnixMilli(n).UTC()

	return t.Format(time.RubyDate)
}

func binanceTradeWs() {

	var df dataframe.DataFrame
	//srs := series.Series{}
	isFirst := true
	columnSymbol := series.Series{}
	columnPrice := series.Series{}
	columnTime := series.Series{}
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
				columnTime = series.New(convertUnixToDateTime(event.Time), series.String, "Time")
				columnPrice = series.New(event.Price, series.Float, "Price")
				isFirst = false
			} else {
				columnSymbol.Append(event.Symbol)
				columnTime.Append(convertUnixToDateTime(event.Time))
				columnPrice.Append(event.Price)
			}
			df = dataframe.New(columnSymbol, columnTime, columnPrice)
			priceRecords := df.Col("Price").Float()

			staregy(priceRecords, 0.0000001)
			// fmt.Println(df)
			df.WriteCSV(f)
			currentTime = now
		}

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

func staregy(priceRecords []float64, entry float64) {
	fPrice := priceRecords[0]
	lPrice := priceRecords[len(priceRecords)-1]
	result := (lPrice / fPrice) - 1
	// fmt.Println(result)

	if !openPosition {
		if result > entry {
			createOrder()
			openPosition = true
		}
	}
	if openPosition {

	}
}

func createOrder() {
	futures.UseTestnet = true
	client := futures.NewClient("35ab0f236f179c4e3aa9126e9f6a5943fb7d3181117a23312c07d4d94ba14c56", "35f23c0014b0af922ba645d4efd6b8ce13003334ae6077e8342f7c317278cf4b")
	createdOrder, err := client.NewCreateOrderService().Symbol("ETHBUSD").Type(futures.OrderTypeMarket).Side(futures.SideTypeBuy).Quantity("0.01").Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("created order : ", createdOrder.UpdateTime)

}

func sellOrder() {
	futures.UseTestnet = true
	client := futures.NewClient("35ab0f236f179c4e3aa9126e9f6a5943fb7d3181117a23312c07d4d94ba14c56", "35f23c0014b0af922ba645d4efd6b8ce13003334ae6077e8342f7c317278cf4b")
	selldOrder, err := client.NewCreateOrderService().Symbol("ETHBUSD").Type(futures.OrderTypeMarket).Side(futures.SideTypeSell).Quantity("0.01").Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("sell order : ", selldOrder)
}
