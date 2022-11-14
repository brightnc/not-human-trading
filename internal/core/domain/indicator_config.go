package domain

type IndicatorConfig struct {
	RSIConfig        RSIConfig        `json:"rsi_config"`
	STOConfig        STOConfig        `json:"sto_config"`
	MACDConfig       MACDConfig       `json:"macd_config"`
	EMAConfig        EMAConfig        `json:"ema_config"`
	SupertrendConfig SupertrendConfig `json:"supertrend_config"`
}

type RSIConfig struct {
	Period int `json:"period"`
}

type STOConfig struct {
	Length int `json:"length"`
	// we fixed this value to be 1, improvement is on our plan to support client to customization of this field in the future
	D int `json:"d"`
	// we fixed this value to be 3, improvement is on our plan to support client to customization of this field in the future
	K int `json:"k"`
}

type MACDConfig struct {
	EMAFastPeriod int `json:"ema_fast_period"`
	EMASlowPeriod int `json:"ema_slow_period"`
	SignalPeriod  int `json:"signal_period"`
}

type EMAConfig struct {
	Period int `json:"period"`
}

type SupertrendConfig struct {
	ATRPeriod  int `json:"atr_period"`
	Multiplier int `json:"multiplier"`
}
