package domain

type PlaceOrder struct {
	Symbol    string
	Quantity  string
	OrderType OrderType
}

type OrderType string

const (
	OrderTypeMarket OrderType = "MARKET"
	OrderTypeUnknow OrderType = "UNKNOWN"
)

type PlaceOrderResult struct {
	Symbol         string    `json:"symbol"`
	OrderID        string    `json:"order_id"`
	Side           OrderSide `json:"side"`
	OrderType      OrderType `json:"order_type"`
	OriginQuantity string    `json:"origin_quantity"`
	ActualQuantity string    `json:"actual_quantity"`
	Price          string    `json:"price"`
}

type OrderSide string

const (
	OrderSideBuy    OrderSide = "BUY"
	OrderSideSell   OrderSide = "SELL"
	OrderSideUnknow OrderSide = "UNKNOWN"
)
