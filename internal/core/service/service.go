package service

import (
	"github.com/brightnc/not-human-trading/internal/core/port"
)

type Service struct {
	exchange        port.Exchange
	botRepo         port.BotConfig
	hasCreatedOrder bool
	isBotRunning    bool
	stopSignal      chan bool
}

func NewService(
	exchangeRepo port.Exchange,
	botRepo port.BotConfig,
) *Service {
	return &Service{
		exchange:        exchangeRepo,
		botRepo:         botRepo,
		hasCreatedOrder: false,
		isBotRunning:    false,
		stopSignal:      make(chan bool),
	}
}
