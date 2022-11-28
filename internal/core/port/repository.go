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
	PlaceBid(domain.PlaceOrder, domain.BotExchange) (domain.PlaceOrderResult, error)
	PlaceAsk(domain.PlaceOrder, domain.BotExchange) (domain.PlaceOrderResult, error)
	Cancel() error
	RetrieveKLines(symbol, startDate, endDate string, period domain.Period) (domain.Quote, error)
}

type ExchangeWithoutFeeFromResp interface {
	// RetriveMyTrades ...
	// retrieving my trading history
	RetriveMyTrades(filter domain.MyTradeFilter)
}

type BotConfig interface {
	UpdateBotConfig(domain.BotConfig) error
	RetrieveBotConfig() (domain.BotConfig, error)
	RetrieveBotExchange() (domain.BotExchange, error)
	UpdateBotExchange(in domain.BotExchange) error
}

type User interface {
	RetrieveUserByID(userID string) (domain.User, error)
}
