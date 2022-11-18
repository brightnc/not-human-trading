package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/brightnc/not-human-trading/internal/core/domain"
	"github.com/markcheno/go-quote"
)

type Binance struct {
	mutext *sync.Mutex
}

func NewBinanceExchange() *Binance {
	return &Binance{
		mutext: &sync.Mutex{},
	}
}

func (r *Binance) PlaceBid() error { return nil }
func (r *Binance) PlaceAsk() error { return nil }
func (r *Binance) Cancel() error   { return nil }

// RetrieveKLines...
// Binance historical prices for a symbol
func (r *Binance) RetrieveKLines(symbol, startDate, endDate string, period domain.Period) (domain.Quote, error) {

	start := quote.ParseDateString(startDate)
	end := quote.ParseDateString(endDate)

	var interval string
	var granularity int // seconds

	switch period {
	case domain.Min1:
		interval = "1m"
		granularity = 60
	case domain.Min3:
		interval = "3m"
		granularity = 3 * 60
	case domain.Min5:
		interval = "5m"
		granularity = 5 * 60
	case domain.Min15:
		interval = "15m"
		granularity = 15 * 60
	case domain.Min30:
		interval = "30m"
		granularity = 30 * 60
	case domain.Min60:
		interval = "1h"
		granularity = 60 * 60
	case domain.Hour2:
		interval = "2h"
		granularity = 2 * 60 * 60
	case domain.Hour4:
		interval = "4h"
		granularity = 4 * 60 * 60
	case domain.Hour8:
		interval = "8h"
		granularity = 8 * 60 * 60
	case domain.Hour12:
		interval = "12h"
		granularity = 12 * 60 * 60
	case domain.Daily:
		interval = "1d"
		granularity = 24 * 60 * 60
	case domain.Day3:
		interval = "3d"
		granularity = 3 * 24 * 60 * 60
	case domain.Weekly:
		interval = "1w"
		granularity = 7 * 24 * 60 * 60
	case domain.Monthly:
		interval = "1M"
		granularity = 30 * 24 * 60 * 60
	default:
		interval = "1d"
		granularity = 24 * 60 * 60
	}

	var appQuote domain.Quote
	appQuote.Symbol = symbol

	maxBars := 500
	var step time.Duration
	step = time.Second * time.Duration(granularity)

	startBar := start
	endBar := startBar.Add(time.Duration(maxBars) * step)

	if endBar.After(end) {
		endBar = end
	}
	quoteMapper := make(map[int]quote.Quote)
	fetchingRound := -1
	var errs error
	var wg sync.WaitGroup
	for startBar.Before(end) && (errs == nil) {
		wg.Add(1)
		fetchingRound++
		go func(sequenceNumber int, wg *sync.WaitGroup) {
			q, err := r.retrieveKlines(symbol, interval, startBar, endBar)
			if err != nil {
				fmt.Println("error while featching kilines from Binance")
				errs = err
				return
			}
			fmt.Println("q from binanace ->>> ", q)
			r.mutext.Lock()
			quoteMapper[sequenceNumber] = quote.Quote{
				Date:   q.Date,
				Open:   q.Open,
				High:   q.High,
				Low:    q.Low,
				Close:  q.Close,
				Volume: q.Volume,
			}
			r.mutext.Unlock()
			defer wg.Done()
		}(fetchingRound, &wg)

		startBar = endBar.Add(step)
		endBar = startBar.Add(time.Duration(maxBars) * step)
	}
	// block until done
	wg.Wait()
	for i := 0; i < fetchingRound; i++ {
		appQuote.Date = append(appQuote.Date, quoteMapper[i].Date...)
		appQuote.Open = append(appQuote.Open, quoteMapper[i].Open...)
		appQuote.High = append(appQuote.High, quoteMapper[i].High...)
		appQuote.Low = append(appQuote.Low, quoteMapper[i].Low...)
		appQuote.Close = append(appQuote.Close, quoteMapper[i].Close...)
		appQuote.Volume = append(appQuote.Volume, quoteMapper[i].Volume...)
	}
	fmt.Println("appQuotes ->>> ", appQuote)
	return appQuote, errs
}

func (r *Binance) retrieveKlines(symbol, interval string, startBar, endBar time.Time) (quote.Quote, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v1/klines?symbol=%s&interval=%s&startTime=%d&endTime=%d",
		strings.ToUpper(symbol),
		interval,
		startBar.UnixNano()/1000000,
		endBar.UnixNano()/1000000)
	client := &http.Client{Timeout: quote.ClientTimeout}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("binance error: %v\n", err)
		return quote.Quote{}, err
	}
	defer resp.Body.Close()

	contents, _ := ioutil.ReadAll(resp.Body)
	type binance [12]interface{}
	var bars []binance
	err = json.Unmarshal(contents, &bars)
	if err != nil {
		fmt.Printf("binance error: %v\n", err)
		return quote.Quote{}, err
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
	for bar := 0; bar < numrows; bar++ {
		q.Date[bar] = time.Unix(int64(bars[bar][6].(float64))/1000, 0)
		q.Open[bar], _ = strconv.ParseFloat(bars[bar][1].(string), 64)
		q.High[bar], _ = strconv.ParseFloat(bars[bar][2].(string), 64)
		q.Low[bar], _ = strconv.ParseFloat(bars[bar][3].(string), 64)
		q.Close[bar], _ = strconv.ParseFloat(bars[bar][4].(string), 64)
		q.Volume[bar], _ = strconv.ParseFloat(bars[bar][5].(string), 64)
	}
	return q, err
}
