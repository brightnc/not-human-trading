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
	UpdateBotConfig(in domain.BotConfig) error
	UpdateBotExchange(in domain.BotExchange) error
	StartBot() error
	StopBot() error
}
