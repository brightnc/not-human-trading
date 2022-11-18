package dto

import (
	"log"

	"github.com/brightnc/not-human-trading/internal/core/domain"
)

type BotConfigRequest struct {
	RSIConfig        RSIConfig        `json:"rsi_config"`
	STOConfig        STOConfig        `json:"sto_config"`
	MACDConfig       MACDConfig       `json:"macd_config"`
	EMAConfig        EMAConfig        `json:"ema_config"`
	SupertrendConfig SupertrendConfig `json:"supertrend_config"`
	OrderConfig      OrderConfig      `json:"order_config"`
	Timeframe        string           `json:"timeframe" validate:"oneof= 1m 3m 5m 15m 30m 1h 2h 4h 6h 8h 12h 1d 3d 1w 1mth"`
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
	var timeframe domain.Period
	switch d.Timeframe {
	case "1m":
		timeframe = domain.Min1
	case "3m":
		timeframe = domain.Min3
	case "5m":
		timeframe = domain.Min5
	case "15m":
		timeframe = domain.Min15
	case "30m":
		timeframe = domain.Min30
	case "1h":
		timeframe = domain.Min60
	case "2h":
		timeframe = domain.Hour2
	case "4h":
		timeframe = domain.Hour4
	case "6h":
		timeframe = domain.Hour6
	case "8h":
		timeframe = domain.Hour8
	case "12h":
		timeframe = domain.Hour12
	case "1d":
		timeframe = domain.Daily
	case "3d":
		timeframe = domain.Day3
	case "1w":
		timeframe = domain.Weekly
	case "1mth":
		timeframe = domain.Monthly
	default:
		//! should not be reached
		log.Println("unexpected timeframe ", d.Timeframe)
		timeframe = domain.Min5
	}
	return domain.BotConfig{
		RSIConfig:        (domain.RSIConfig)(d.RSIConfig),
		STOConfig:        (domain.STOConfig)(d.STOConfig),
		MACDConfig:       (domain.MACDConfig)(d.MACDConfig),
		EMAConfig:        (domain.EMAConfig)(d.EMAConfig),
		SupertrendConfig: (domain.SupertrendConfig)(d.SupertrendConfig),
		OrderConfig:      (domain.OrderConfig)(d.OrderConfig),
		Timeframe:        timeframe,
	}
}

type UpdateBotExchangeRequest struct {
	APIKey    string `json:"api_key"`
	SecretKey string `json:"secret_key"`
}

func (d UpdateBotExchangeRequest) ToBotExchangeDomain() domain.BotExchange {
	return (domain.BotExchange)(d)
}
