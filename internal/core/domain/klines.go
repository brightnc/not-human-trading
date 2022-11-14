package domain

import "time"

type Quote struct {
	Symbol    string
	Precision int64
	Date      []time.Time
	Open      []float64
	High      []float64
	Low       []float64
	Close     []float64
	Volume    []float64
}
