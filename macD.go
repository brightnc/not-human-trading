package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/cinar/indicator"
	"github.com/go-gota/gota/dataframe"
	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

type BitkubTradHistoryResponse struct {
	Close  []float64 `json:"c"`
	High   []float64 `json:"h"`
	Low    []float64 `json:"l"`
	Open   []float64 `json:"o"`
	status string    `json:"s"`
	Time   []int64   `json:"t"`
	Volume []float64 `json:"v"`
}

type Trader struct {
	datasource dataframe.DataFrame
}

type strategy interface {
	EMABuySignal(...[]emaConfig)
	MACDSignal(...macDConfig)
	Cumulative(entry float64)
}

type emaConfig struct {
	values   []float64
	period   int
	priority int
	//operaton bool
}
type macDConfig struct {
	values       []float64
	emaFast      int
	emaSlow      int
	signalPeriod int
}

type RSIConfig struct {
	values []float64
	period int
}
type STOConfig struct {
	High  []float64
	Low   []float64
	Close []float64
}

type SPTConfig struct {
	High       []float64
	Low        []float64
	Close      []float64
	ATRPeriod  int
	Multiplier float64
	Time       []time.Time
}

func RSIBuySignal(configs RSIConfig) bool {
	b := false
	rsi := talib.Rsi(configs.values, configs.period)
	fmt.Println(rsi[len(rsi)-1])
	if rsi[len(rsi)-1] < 30 {
		b = true
	}

	return b

}

func RSISellSignal(configs RSIConfig) bool {
	b := false
	rsi := talib.Rsi(configs.values, configs.period)
	fmt.Println(rsi[len(rsi)-1])
	if rsi[len(rsi)-1] > 70 {
		b = true
	}

	return b

}

func SuperTrendDetail(configs SPTConfig) ([]float64, []float64, []float64, []bool, []time.Time) {
	l := len(configs.High)
	hl2 := talib.MedPrice(configs.High, configs.Low)
	atr := talib.Atr(configs.High, configs.Low, configs.Close, configs.ATRPeriod)
	fmt.Println(atr[len(atr)-5:])

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
			times = append(times, configs.Time[i])
			trend[i] = true
		} else if configs.Close[i] < trendUp[i-1] {
			times = append(times, configs.Time[i])
			trend[i] = false
		} else {
			times = append(times, configs.Time[i])
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

func SuperTrend(configs SPTConfig) {
	atr := talib.Atr(configs.High, configs.Low, configs.Close, configs.ATRPeriod)
	_, atr2 := indicator.Atr(configs.ATRPeriod, configs.High, configs.Low, configs.Close)
	upper := (configs.High[len(configs.High)-1]+configs.Low[len(configs.Low)-1])/2 + float64(configs.Multiplier)*atr[len(atr)-1]
	lower := (configs.High[len(configs.High)-1]+configs.Low[len(configs.Low)-1])/2 - float64(configs.Multiplier)*atr[len(atr)-1]
	fmt.Printf("RED upper >> %v , GREEN lower >> %v \n", upper, lower)
	fmt.Printf("ATR >> %v, ATR2 >> %v \n", atr[len(atr)-3:], atr2[len(atr2)-3:])

	upper2 := (configs.High[len(configs.High)-1]+configs.Low[len(configs.Low)-1])/2 + float64(configs.Multiplier)*atr2[len(atr2)-1]
	lower2 := (configs.High[len(configs.High)-1]+configs.Low[len(configs.Low)-1])/2 - float64(configs.Multiplier)*atr2[len(atr2)-1]
	fmt.Printf("RED upper 2 >> %v , GREEN lower 2 >> %v \n", upper2, lower2)
}

func STOBuySignal(configs STOConfig) bool {
	now := time.Now()
	i := false
	k, d := indicator.StochasticOscillator(configs.High, configs.Low, configs.Close)
	fmt.Printf("Time >>> %v , K >>> %v , D >>> %v\n", now, k[len(k)-1], d[len(d)-1])
	if talib.Crossover(k, d) && k[len(k)-1] < 20 && d[len(d)-1] < 20 {
		i = true
	}
	return i

}

func STOSellSignal(configs STOConfig) bool {
	now := time.Now()
	i := false
	k, d := indicator.StochasticOscillator(configs.High, configs.Low, configs.Close)
	fmt.Printf("Time >>> %v , K >>> %v , D >>> %v\n", now, k[len(k)-1], d[len(d)-1])
	if talib.Crossunder(k, d) && k[len(k)-1] > 80 && d[len(d)-1] > 80 {
		i = true
	}
	return i

}
func MACDBuySignal(configs macDConfig) bool {

	// now := time.Now()
	// macD, macDSignal, _ := talib.Macd(configs.values, configs.emaFast, configs.emaSlow, configs.signalPeriod)
	// fmt.Println("Time >>> ", now, " ", macD[len(macDSignal)-5:], macDSignal[len(macDSignal)-5:])
	// return talib.Crossover(macD, macDSignal)

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
	v := (lastValueOfMACD - lastValueOfSignal) / lastValueOfSignal
	// fmt.Println("Fast >>> ", len(emaFast), " Slow >>> ", len(emaSlow))
	// fmt.Println("current value >> ", lastValueOfMACD, " ", lastValueOfSignal, " Signal >>> ", emaSignal[len(emaSignal)-2:])
	if v >= -0.1 && v <= 0.1 {
		fmt.Println("almost cross >> ", lastValueOfMACD, " ", lastValueOfSignal)
	}

	return talib.Crossover(macD, emaSignal)

}

func MACDSellSignal(configs macDConfig) bool {
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
	v := (lastValueOfMACD - lastValueOfSignal) / lastValueOfSignal
	fmt.Println("Fast >>> ", len(emaFast), " Slow >>> ", len(emaSlow))
	fmt.Println("current value >> ", lastValueOfMACD, " ", lastValueOfSignal, " Signal >>> ", emaSignal[len(emaSignal)-2:])
	if v >= -0.1 && v <= 0.1 {
		fmt.Println("almost cross >> ", lastValueOfMACD, " ", lastValueOfSignal)
	}

	return talib.Crossunder(macD, emaSignal)
}

func ValuesToEMA(values []float64, period int) []float64 {
	return talib.Ema(values, period)
}

func EMASignal(isBuyAction bool, configs ...emaConfig) bool {
	// q := quote.NewQuote("BTCUSDT", numrows)

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
		if isBuyAction {
			isOK = mapper[i].indicatorValue > mapper[i+1].indicatorValue
		} else {
			isOK = mapper[i].indicatorValue < mapper[i+1].indicatorValue
		}
	}

	return isOK
	//fmt.Print(spy.CSV())
	// dema := talib.Ema(spy.Close, 10)
	// dema2 := talib.Ema(spy.Close, 20)
	// interestedDema := dema[len(dema)-1]
	// interestedDema2 := dema2[len(dema2)-1]
	// fmt.Println("dema -> ", dema[len(dema)-6:])
	// fmt.Println("dema2 -> ", dema2[len(dema2)-6:])
	//fmt.Println(interestedDema > interestedDema2)
	//fmt.Println(interestedDema < interestedDema2)
	//fmt.Println(talib.Crossover(interestedDema, interestedDema2))
}

func NewQuoteFromBitkub(symbol string, startDate, endDate string, period quote.Period) (quote.Quote, error) {

	start := quote.ParseDateString(startDate)
	end := quote.ParseDateString(endDate)

	var interval string
	var granularity int // seconds

	switch period {
	case quote.Min1:
		interval = "1"
		granularity = 60
	case quote.Min5:
		interval = "5"
		granularity = 5 * 60
	case quote.Min15:
		interval = "15"
		granularity = 15 * 60
	case quote.Min60:
		interval = "60"
		granularity = 60 * 60
	case quote.Hour4:
		interval = "240"
		granularity = 4 * 60 * 60
	case quote.Daily:
		interval = "1D"
		granularity = 24 * 60 * 60
	default:
		interval = "1D"
		granularity = 24 * 60 * 60
	}

	var newQ quote.Quote
	newQ.Symbol = symbol

	maxBars := 500
	step := time.Second * time.Duration(granularity)

	startBar := start
	endBar := startBar.Add(time.Duration(maxBars) * step)

	if endBar.After(end) {
		endBar = end
	}

	// for startBar.Before(end) {

	url := fmt.Sprintf(
		"https://api.bitkub.com/tradingview/history?symbol=%s&resolution=%s&from=%d&to=%d",
		strings.ToUpper(symbol),
		interval,
		start.Unix(),
		end.Unix(),
	)
	// startBar.Unix(),
	// endBar.Unix())
	log.Println(url)
	client := &http.Client{Timeout: quote.ClientTimeout}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("binance error: %v\n", err)
		return quote.NewQuote("", 0), err
	}
	defer resp.Body.Close()

	contents, _ := ioutil.ReadAll(resp.Body)

	var bars BitkubTradHistoryResponse
	err = json.Unmarshal(contents, &bars)
	if err != nil {
		log.Printf("binance error: %v\n", err)
	}

	numrows := len(bars.Volume)
	q := quote.NewQuote(symbol, numrows)
	//fmt.Printf("numrows=%d, bars=%v\n", numrows, bars)

	/*
		0       OpenTime                 int64
		1 			Open                     float64
		2 			High                     float64
		3		 	Low                      float64
		4 			Close                    float64
		5 			Volume                   float64
		6 			CloseTime                int64
		7 			QuoteAssetVolume         float64
		8 			NumTrades                int64
		9 			TakerBuyBaseAssetVolume  float64
		10 			TakerBuyQuoteAssetVolume float64
		11 			Ignore                   float64
	*/

	for bar := 0; bar < numrows; bar++ {
		q.Date[bar] = time.Unix(bars.Time[bar], 0)
		q.Open[bar] = bars.Open[bar]
		q.High[bar] = bars.High[bar]
		q.Low[bar] = bars.Low[bar]
		q.Close[bar] = bars.Close[bar]
		q.Volume[bar] = bars.Volume[bar]
	}
	newQ.Date = append(newQ.Date, q.Date...)
	newQ.Open = append(newQ.Open, q.Open...)
	newQ.High = append(newQ.High, q.High...)
	newQ.Low = append(newQ.Low, q.Low...)
	newQ.Close = append(newQ.Close, q.Close...)
	newQ.Volume = append(newQ.Volume, q.Volume...)

	// time.Sleep(time.Second)
	// startBar = endBar.Add(step)
	// endBar = startBar.Add(time.Duration(maxBars) * step)

	//}
	return newQ, nil
}

func NewQuoteFromBinance(symbol string, startDate, endDate string, period quote.Period) (quote.Quote, error) {

	start := quote.ParseDateString(startDate)
	end := quote.ParseDateString(endDate)

	var interval string
	var granularity int // seconds

	switch period {
	case quote.Min1:
		interval = "1m"
		granularity = 60
	case quote.Min3:
		interval = "3m"
		granularity = 3 * 60
	case quote.Min5:
		interval = "5m"
		granularity = 5 * 60
	case quote.Min15:
		interval = "15m"
		granularity = 15 * 60
	case quote.Min30:
		interval = "30m"
		granularity = 30 * 60
	case quote.Min60:
		interval = "1h"
		granularity = 60 * 60
	case quote.Hour2:
		interval = "2h"
		granularity = 2 * 60 * 60
	case quote.Hour4:
		interval = "4h"
		granularity = 4 * 60 * 60
	case quote.Hour8:
		interval = "8h"
		granularity = 8 * 60 * 60
	case quote.Hour12:
		interval = "12h"
		granularity = 12 * 60 * 60
	case quote.Daily:
		interval = "1d"
		granularity = 24 * 60 * 60
	case quote.Day3:
		interval = "3d"
		granularity = 3 * 24 * 60 * 60
	case quote.Weekly:
		interval = "1w"
		granularity = 7 * 24 * 60 * 60
	case quote.Monthly:
		interval = "1M"
		granularity = 30 * 24 * 60 * 60
	default:
		interval = "1d"
		granularity = 24 * 60 * 60
	}

	var quotes quote.Quote
	quotes.Symbol = symbol

	maxBars := 500
	var step time.Duration
	step = time.Second * time.Duration(granularity)

	startBar := start
	endBar := startBar.Add(time.Duration(maxBars) * step)

	if endBar.After(end) {
		endBar = end
	}
	datasetMapper := make(map[int]quote.Quote)
	datasetCapacity := []int{}
	roundCounter := 0
	for startBar.Before(end) {

		url := fmt.Sprintf(
			"https://api.binance.com/api/v1/klines?symbol=%s&interval=%s&startTime=%d&endTime=%d",
			strings.ToUpper(symbol),
			interval,
			startBar.UnixNano()/1000000,
			endBar.UnixNano()/1000000)
		//log.Println(url)
		client := &http.Client{Timeout: quote.ClientTimeout}
		req, _ := http.NewRequest("GET", url, nil)
		resp, err := client.Do(req)

		if err != nil {
			fmt.Printf("binance error: %v\n", err)
			return quote.NewQuote("", 0), err
		}
		defer resp.Body.Close()

		contents, _ := ioutil.ReadAll(resp.Body)

		type binance [12]interface{}
		var bars []binance
		err = json.Unmarshal(contents, &bars)
		if err != nil {
			fmt.Printf("binance error: %v\n", err)
		}

		numrows := len(bars)
		q := quote.NewQuote(symbol, numrows)
		//fmt.Printf("numrows=%d, bars=%v\n", numrows, bars)

		/*
			0       OpenTime                 int64
			1 			Open                     float64
			2 			High                     float64
			3		 	Low                      float64
			4 			Close                    float64
			5 			Volume                   float64
			6 			CloseTime                int64
			7 			QuoteAssetVolume         float64
			8 			NumTrades                int64
			9 			TakerBuyBaseAssetVolume  float64
			10 			TakerBuyQuoteAssetVolume float64
			11 			Ignore                   float64
		*/
		go func(i int) {
			for bar := 0; bar < numrows; bar++ {
				q.Date[bar] = time.Unix(int64(bars[bar][6].(float64))/1000, 0)
				q.Open[bar], _ = strconv.ParseFloat(bars[bar][1].(string), 64)
				q.High[bar], _ = strconv.ParseFloat(bars[bar][2].(string), 64)
				q.Low[bar], _ = strconv.ParseFloat(bars[bar][3].(string), 64)
				q.Close[bar], _ = strconv.ParseFloat(bars[bar][4].(string), 64)
				q.Volume[bar], _ = strconv.ParseFloat(bars[bar][5].(string), 64)
			}
			datasetMapper[i] = q
			datasetCapacity = append(datasetCapacity, i)
			// quotes.Date = append(quotes.Date, q.Date...)
			// quotes.Open = append(quotes.Open, q.Open...)
			// quotes.High = append(quotes.High, q.High...)
			// quotes.Low = append(quotes.Low, q.Low...)
			// quotes.Close = append(quotes.Close, q.Close...)
			// quotes.Volume = append(quotes.Volume, q.Volume...)

		}(roundCounter)

		time.Sleep(time.Second)
		startBar = endBar.Add(step)
		endBar = startBar.Add(time.Duration(maxBars) * step)
		roundCounter++
	}

	for i := range datasetCapacity {
		quotes.Date = append(quotes.Date, datasetMapper[i].Date...)
		quotes.Open = append(quotes.Open, datasetMapper[i].Open...)
		quotes.High = append(quotes.High, datasetMapper[i].High...)
		quotes.Low = append(quotes.Low, datasetMapper[i].Low...)
		quotes.Close = append(quotes.Close, datasetMapper[i].Close...)
		quotes.Volume = append(quotes.Volume, datasetMapper[i].Volume...)
	}

	return quotes, nil
}
