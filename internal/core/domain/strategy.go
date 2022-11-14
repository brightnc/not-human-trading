package domain

type RSI struct {
	period int
}

type Stratagy struct {
	EMA, RSI, STO, Supertrend, MACD bool
}
