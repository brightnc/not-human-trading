package dto

import "github.com/brightnc/not-human-trading/internal/core/domain"

type BotConfigRequest struct {
	RSIConfig        RSIConfig        `json:"rsi_config"`
	STOConfig        STOConfig        `json:"sto_config"`
	MACDConfig       MACDConfig       `json:"macd_config"`
	EMAConfig        EMAConfig        `json:"ema_config"`
	SupertrendConfig SupertrendConfig `json:"supertrend_config"`
	OrderConfig      OrderConfig      `json:"order_config"`
	Timeframe        string           `json:"timeframe"`
}

type RSIConfig struct {
	IsActive bool `json:"is_active"`
	Period   int  `json:"period"`
}

type STOConfig struct {
	IsActive bool `json:"is_active"`
	Length   int  `json:"length"`
	// we fixed this value to be 1, improvement is on our plan to support client to customization of this field in the future
	D int `json:"d"`
	// we fixed this value to be 3, improvement is on our plan to support client to customization of this field in the future
	K int `json:"k"`
}

type MACDConfig struct {
	IsActive      bool `json:"is_active"`
	EMAFastPeriod int  `json:"ema_fast_period"`
	EMASlowPeriod int  `json:"ema_slow_period"`
	SignalPeriod  int  `json:"signal_period"`
}

type EMAConfig struct {
	IsActive   bool `json:"is_active"`
	FastPeriod int  `json:"fast_period"`
	SlowPeriod int  `json:"slow_period"`
}

type SupertrendConfig struct {
	IsActive   bool `json:"is_active"`
	ATRPeriod  int  `json:"atr_period"`
	Multiplier int  `json:"multiplier"`
}

type OrderConfig struct {
	Symbol   string  `json:"sym"`
	Quantity float64 `json:"qty"`
}

// ToBotConfigDomain ...
// convert protocol payload to be domain object
func (d BotConfigRequest) ToBotConfigDomain() domain.BotConfig {
	return domain.BotConfig{
		RSIConfig:        (domain.RSIConfig)(d.RSIConfig),
		STOConfig:        (domain.STOConfig)(d.STOConfig),
		MACDConfig:       (domain.MACDConfig)(d.MACDConfig),
		EMAConfig:        (domain.EMAConfig)(d.EMAConfig),
		SupertrendConfig: (domain.SupertrendConfig)(d.SupertrendConfig),
		OrderConfig:      (domain.OrderConfig)(d.OrderConfig),
		Timeframe:        d.Timeframe,
	}
}

type UpdateBotExchangeRequest struct {
	APIKey    string `json:"api_key"`
	SecretKey string `json:"secret_key"`
}

func (d UpdateBotExchangeRequest) ToBotExchangeDomain() domain.BotExchange {
	return (domain.BotExchange)(d)
}
