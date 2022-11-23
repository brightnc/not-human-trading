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
	Symbol         string
	Side           OrderSide
	OrderType      OrderType
	OriginQuantity string
	ActualQuantity string
	Price          string
}

type OrderSide string

const (
	OrderSideBuy    OrderSide = "BUY"
	OrderSideSell   OrderSide = "SELL"
	OrderSideUnknow OrderSide = "UNKNOWN"
)
