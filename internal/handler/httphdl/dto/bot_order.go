package dto

type BotOrderRequest struct {
	Symbol   string  `json:"sym"`
	Quantity float64 `json:"qty"`
}
