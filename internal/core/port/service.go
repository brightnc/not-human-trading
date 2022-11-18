package port

import (
	"github.com/brightnc/not-human-trading/internal/core/domain"
)

/*
	|--------------------------------------------------------------------------
	| Application Port
	|--------------------------------------------------------------------------
	|
	| Here you can define an interface which will allow foreign actors to
	| communicate with the Application
	|
*/

type Service interface {
	UpdateIndicator(in domain.IndicatorConfig) error
	UpdateBotConfig(domain.BotConfig) error
	UpdateBotOrder(in domain.BotOrder) error
	StartBot() error
	StopBot() error
}
