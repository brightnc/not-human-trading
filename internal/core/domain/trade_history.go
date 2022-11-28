package domain

type MyTradeFilter struct {
	Symbol  *string
	OrderID *string
	Limit   *int
}

type MyTradeResult struct {
	Symbol  string
	OrderID int
	Qty     float64
	Fee     float64
	IsBuyer bool
	Time    int64
}
