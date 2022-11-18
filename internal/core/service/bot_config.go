package service

import (
	"github.com/brightnc/not-human-trading/internal/core/domain"
)

/*
	|--------------------------------------------------------------------------
	| Application's Business Logic
	|--------------------------------------------------------------------------
	|
	| Here you can implement a business logic  for your application
	|
*/

// type botOption func(*Service)

// func WithEMA(period int) botOption {
// 	return func(s *Service) {
// 		s.stratagies = append(s.stratagies, domain.EMAConfig{Period: period})
// 	}
// }

// func WithMACD(fastPeriod, slowPeriod, signalPeriod int) botOption {
// 	return func(s *Service) {
// 		s.stratagies = append(s.stratagies, domain.MACDConfig{EMAFastPeriod: fastPeriod, EMASlowPeriod: slowPeriod, SignalPeriod: signalPeriod})
// 	}
// }

// func WithRSI(period int) botOption {
// 	return func(s *Service) {
// 		s.stratagies = append(s.stratagies, domain.RSIConfig{Period: period})
// 	}
// }

// func WithSTO(length int) botOption {
// 	return func(s *Service) {
// 		s.stratagies = append(s.stratagies, domain.STOConfig{Length: length, D: 1, K: 3})
// 	}
// }

// func WithSupertrend(atrPeriod, multiplier int) botOption {
// 	return func(s *Service) {
// 		s.stratagies = append(s.stratagies, domain.SupertrendConfig{ATRPeriod: atrPeriod, Multiplier: multiplier})
// 	}
// }

func (svc *Service) UpdateBotConfig(in domain.BotConfig) error {
	err := svc.botRepo.UpdateBotConfig(in)
	if err != nil {
		// TODO: handle error
		panic(err)
	}
	return err
}

func (svc *Service) UpdateBotExchange(in domain.BotExchange) error {
	err := svc.botRepo.UpdateBotExchange(in)
	if err != nil {
		// TODO: handle error
		panic(err)
	}
	return err
}

func (svc *Service) StartBot() error {

	// quote, err := svc.repository.RetrieveKLines()
	// if err != nil {
	// 	// TODO: handle error
	// 	panic(err)
	// }

	// quo
	return nil

}

func (svc *Service) StopBot() error { return nil }

func (svc *Service) UpdateBot() error { return nil }

// func emaStratagy(){
// 	func EMASignal(isBuyAction bool, configs ...emaConfig) bool {
// 		// q := quote.NewQuote("BTCUSDT", numrows)

// 		type indicator struct {
// 			indicatorValue float64
// 		}
// 		mapper := make(map[int]indicator, len(configs))
// 		isOK := false
// 		for i := range configs {
// 			ema := talib.Ema(configs[i].values, configs[i].period)
// 			mapper[configs[i].priority-1] = indicator{
// 				indicatorValue: ema[len(ema)-1],
// 			}
// 		}

// 		for i := range configs {
// 			if i == len(configs)-1 {
// 				break
// 			}
// 			fmt.Println("EMA first priority = ", mapper[i].indicatorValue)
// 			fmt.Println("EMA second priority = ", mapper[i+1].indicatorValue)
// 			if isBuyAction {
// 				isOK = mapper[i].indicatorValue > mapper[i+1].indicatorValue
// 			} else {
// 				isOK = mapper[i].indicatorValue < mapper[i+1].indicatorValue
// 			}
// 		}

// 		return isOK
// 		//fmt.Print(spy.CSV())
// 		// dema := talib.Ema(spy.Close, 10)
// 		// dema2 := talib.Ema(spy.Close, 20)
// 		// interestedDema := dema[len(dema)-1]
// 		// interestedDema2 := dema2[len(dema2)-1]
// 		// fmt.Println("dema -> ", dema[len(dema)-6:])
// 		// fmt.Println("dema2 -> ", dema2[len(dema2)-6:])
// 		//fmt.Println(interestedDema > interestedDema2)
// 		//fmt.Println(interestedDema < interestedDema2)
// 		//fmt.Println(talib.Crossover(interestedDema, interestedDema2))
// 	}
// }
