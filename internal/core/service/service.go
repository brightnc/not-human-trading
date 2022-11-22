package service

import (
	"github.com/brightnc/not-human-trading/internal/core/domain"
	"github.com/brightnc/not-human-trading/internal/core/port"
)

type Service struct {
	exchange        port.Exchange
	botRepo         port.BotConfig
	hasCreatedOrder bool
	hasStopSignal   bool
	subscribers     []*domain.Connection
}

func NewService(
	exchangeRepo port.Exchange,
	botRepo port.BotConfig,
) *Service {
	return &Service{
		exchange:        exchangeRepo,
		botRepo:         botRepo,
		hasCreatedOrder: false,
		hasStopSignal:   false,
	}
}
