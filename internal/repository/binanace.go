package repository

import "github.com/brightnc/not-human-trading/internal/core/domain"

type Binance struct{}

func NewBinanceExchange() *Binance {
	return &Binance{}
}

func (r *Binance) PlaceBid() error                       { return nil }
func (r *Binance) PlaceAsk() error                       { return nil }
func (r *Binance) Cancel() error                         { return nil }
func (r *Binance) RetrieveKLines() (domain.Quote, error) { return domain.Quote{}, nil }
