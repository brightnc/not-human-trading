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
	IndicatorAdjustment(in domain.IndicatorConfig) error
	BotAdjusment() error
	BotStarting() error
	BotStoping() error
}
