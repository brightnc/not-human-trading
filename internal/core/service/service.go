package service

import "github.com/brightnc/not-human-trading/internal/core/port"

type Service struct {
	// apply multiple stratagies ...
	stratagies       map[string][]interface{}
	openOrderQuantiy float64
	exchange         port.Exchange
	botRepo          port.BotConfig
}

func NewService(
	exchangeRepo port.Exchange,
	botRepo port.BotConfig,
) *Service {
	return &Service{
		exchange: exchangeRepo,
		botRepo:  botRepo,
	}
}
