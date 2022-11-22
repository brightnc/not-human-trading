package domain

type BotConfig struct {
	RSIConfig        RSIConfig
	STOConfig        STOConfig
	MACDConfig       MACDConfig
	EMAConfig        EMAConfig
	SupertrendConfig SupertrendConfig
	OrderConfig      OrderConfig
	Timeframe        Period
}

type RSIConfig struct {
	IsActive bool
	Period   int
}

type STOConfig struct {
	IsActive bool
	Length   int
	// we fixed this value to be 1, improvement is on our plan to support client to customization of this field in the future
	D int
	// we fixed this value to be 3, improvement is on our plan to support client to customization of this field in the future
	K int
}

type MACDConfig struct {
	IsActive      bool
	EMAFastPeriod int
	EMASlowPeriod int
	SignalPeriod  int
}

type EMAConfig struct {
	IsActive   bool
	FastPeriod int
	SlowPeriod int
}

type SupertrendConfig struct {
	IsActive   bool
	ATRPeriod  int
	Multiplier int
}

type OrderConfig struct {
	Symbol   string
	Quantity string
}

type BotExchange struct {
	APIKey    string
	SecretKey string
}
