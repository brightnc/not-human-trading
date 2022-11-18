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
	RetrieveKLines() (domain.Quote, error)
}

type Indicator interface {
	UpdateIndicator(domain.IndicatorConfig) error
}

type BotConfig interface {
	UpdateBotConfig(domain.BotConfig) error
}

type BotOrder interface {
	UpdateBotOrder(domain.BotOrder) error
}
