package dto

import "github.com/brightnc/not-human-trading/internal/core/domain"

type BotConfigRequest struct {
	ApiKey    string `json:"api_key"`
	SecretKey string `json:"secret_key"`
}

func (d BotConfigRequest) ToBotConfigDomain() domain.BotConfig {
	return domain.BotConfig{
		ApiKey:    d.ApiKey,
		SecretKey: d.SecretKey,
	}
}
