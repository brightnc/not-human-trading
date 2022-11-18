package dto

import "github.com/brightnc/not-human-trading/internal/core/domain"

type BotOrderRequest struct {
	Symbol   string  `json:"sym"`
	Quantity float64 `json:"qty"`
}

func (d BotOrderRequest) ToBotOrderDomain() domain.BotOrder {
	return (domain.BotOrder)(d)
}
