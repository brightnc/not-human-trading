package service

import (
	"github.com/brightnc/not-human-trading/internal/core/domain"
	"github.com/brightnc/not-human-trading/internal/core/port"
)

/*
	|--------------------------------------------------------------------------
	| Application's Business Logic
	|--------------------------------------------------------------------------
	|
	| Here you can implement a business logic  for your application
	|
*/

type Trader struct {
	// apply multiple stratagies ...
	stratagies       []interface{}
	openOrderQuantiy float64
}

type Service struct {
	repository port.Repository
}

func New(repository port.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (svc *Service) IndicatorAdjustment(request domain.IndicatorConfig) error {

	return svc.repository.SomeFunction()
}

func (svc *Service) BotStarting() error {

	quote, err := svc.repository.RetrieveKLines()
	if err != nil {
		// TODO: handle error
		panic(err)
	}

	quo

}



func emaStratagy(){
	func EMASignal(isBuyAction bool, configs ...emaConfig) bool {
		// q := quote.NewQuote("BTCUSDT", numrows)
	
		type indicator struct {
			indicatorValue float64
		}
		mapper := make(map[int]indicator, len(configs))
		isOK := false
		for i := range configs {
			ema := talib.Ema(configs[i].values, configs[i].period)
			mapper[configs[i].priority-1] = indicator{
				indicatorValue: ema[len(ema)-1],
			}
		}
	
		for i := range configs {
			if i == len(configs)-1 {
				break
			}
			fmt.Println("EMA first priority = ", mapper[i].indicatorValue)
			fmt.Println("EMA second priority = ", mapper[i+1].indicatorValue)
			if isBuyAction {
				isOK = mapper[i].indicatorValue > mapper[i+1].indicatorValue
			} else {
				isOK = mapper[i].indicatorValue < mapper[i+1].indicatorValue
			}
		}
	
		return isOK
		//fmt.Print(spy.CSV())
		// dema := talib.Ema(spy.Close, 10)
		// dema2 := talib.Ema(spy.Close, 20)
		// interestedDema := dema[len(dema)-1]
		// interestedDema2 := dema2[len(dema2)-1]
		// fmt.Println("dema -> ", dema[len(dema)-6:])
		// fmt.Println("dema2 -> ", dema2[len(dema2)-6:])
		//fmt.Println(interestedDema > interestedDema2)
		//fmt.Println(interestedDema < interestedDema2)
		//fmt.Println(talib.Crossover(interestedDema, interestedDema2))
	}
}