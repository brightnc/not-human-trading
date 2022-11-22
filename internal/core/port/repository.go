package port

import "github.com/brightnc/not-human-trading/internal/core/domain"

/*
	|--------------------------------------------------------------------------
	| Application Port
	|--------------------------------------------------------------------------
	|
	| Here you can define an interface which will allow foreign actors to
	| communicate with the Application
	|
*/

type Exchange interface {
	PlaceBid() error
	PlaceAsk() error
	Cancel() error
	RetrieveKLines(symbol, startDate, endDate string, period domain.Period) (domain.Quote, error)
}

type BotConfig interface {
	UpdateBotConfig(domain.BotConfig) error
	RetrieveBotConfig() (domain.BotConfig, error)
	RetrieveBotExchange() (domain.BotExchange, error)
	UpdateBotExchange(in domain.BotExchange) error
}
