package service

import "github.com/brightnc/not-human-trading/internal/core/domain"

func (svc *Service) UpdateBotOrder(in domain.BotOrder) error {
	return svc.botOrder.UpdateBotOrder(in)
}
