package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/brightnc/not-human-trading/internal/core/domain"
)

type indicatorConfig struct {
	EMA        emaConfig        `json:"ema"`
	MACD       macdConfig       `json:"macd"`
	RSI        rsiConfig        `json:"rsi"`
	STO        stoConfig        `json:"sto"`
	Supertrend supertrendConfig `json:"supertrend"`
}

type emaConfig struct {
	IsActive   bool `json:"isActive"`
	FastPeriod int  `json:"fastPeriod"`
	SlowPeriod int  `json:"slowPeriod"`
}

type macdConfig struct {
	IsActive      bool `json:"isActive"`
	EMAFastPeriod int  `json:"emaFastPeriod"`
	EMASlowPeriod int  `json:"emaSlowPeriod"`
	SignalPeriod  int  `json:"signalPeriod"`
}

type rsiConfig struct {
	IsActive bool `json:"isActive"`
	Period   int  `json:"period"`
}

type stoConfig struct {
	IsActive bool `json:"isActive"`
	Length   int  `json:"kLength"`

	// multiplier
	D int `json:"d"`
	K int `json:"k"`
}

type supertrendConfig struct {
	IsActive   bool `json:"isActive"`
	ATRPeriod  int  `json:"atrPeriod"`
	Multiplier int  `json:"multiplier"`
}

const (
	indicatorConfigFileName = "config.json"
)

type Indicator struct{}

func NewIndicator() *Indicator {
	return &Indicator{}
}

func (ind *Indicator) UpdateIndicator(indicator domain.IndicatorConfig) error {
	_, fileName, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(fileName)
	fmt.Println(currentDir)
	// var d indicatorConfig
	// s, err := json.Marshal(d)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(s))

	configFile := fmt.Sprintf("%s/%s", currentDir, indicatorConfigFileName)
	f, err := os.Create(configFile)
	if err != nil {
		// TODO: handle proper error
		panic(err)
	}
	config := indicatorConfig{
		EMA:        (emaConfig)(indicator.EMAConfig),
		MACD:       (macdConfig)(indicator.MACDConfig),
		RSI:        (rsiConfig)(indicator.RSIConfig),
		STO:        (stoConfig)(indicator.STOConfig),
		Supertrend: (supertrendConfig)(indicator.SupertrendConfig),
	}
	configJSON, err := json.Marshal(&config)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(configJSON)
	if err != nil {
		// TODO: handle proper error
		panic(err)
	}
	defer f.Close()
	//	configFile := fmt.Println("%s", dir)
	// f, err := os.Open(configFile)
	// if err != nil {
	// 	// TODO: handle proper error
	// 	panic(err)
	// }
	// records, err := csv.NewReader(f).ReadAll()
	// if err != nil {
	// 	// TODO: handle proper error
	// 	panic(err)
	// }

	// for _, record := range records {
	// 	fmt.Println(record)
	// }
	// b, err := ioutil.ReadFile(configFile)
	// if err != nil {
	// 	// TODO: handle proper error
	// 	panic(err)
	// }
	// config := indicatorConfig{}
	// err = json.Unmarshal(b, &config)
	// if err != nil {
	// 	// TODO: handle proper error
	// 	panic(err)
	// }
	// fmt.Println("config >>> ", config)
	return nil
}
