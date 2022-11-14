package trader

type Trader interface {
	PlaceBid() error
	PlaceAsk() error
	Cancel() error
}

type trader struct {
}
