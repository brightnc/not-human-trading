package service

import "github.com/brightnc/not-human-trading/internal/core/port"

type Service struct {
	// apply multiple stratagies ...
	stratagies       map[string][]interface{}
	openOrderQuantiy float64
	exchange         port.Exchange
	indicator        port.Indicator
	botConfig        port.BotConfig
	botOrder         port.BotOrder
}

func NewService(
	exchangeRepo port.Exchange,
	indicatorRepo port.Indicator,
	botConfigRepo port.BotConfig,
	botOrderRepo port.BotOrder,
) *Service {
	return &Service{
		exchange:  exchangeRepo,
		indicator: indicatorRepo,
		botConfig: botConfigRepo,
		botOrder:  botOrderRepo,
	}
}
