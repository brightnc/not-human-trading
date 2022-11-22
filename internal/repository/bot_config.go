package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/brightnc/not-human-trading/internal/core/domain"
)

type botConfig struct {
	EMA        emaConfig        `json:"ema"`
	MACD       macdConfig       `json:"macd"`
	RSI        rsiConfig        `json:"rsi"`
	STO        stoConfig        `json:"sto"`
	Supertrend supertrendConfig `json:"supertrend"`
	Order      botOrder         `json:"botOrder"`
	Timeframe  string           `json:"timeframe"`
}

func (cfg botConfig) ToBotConfigDomain() domain.BotConfig {
	return domain.BotConfig{
		RSIConfig:        (domain.RSIConfig)(cfg.RSI),
		STOConfig:        (domain.STOConfig)(cfg.STO),
		MACDConfig:       (domain.MACDConfig)(cfg.MACD),
		EMAConfig:        (domain.EMAConfig)(cfg.EMA),
		SupertrendConfig: (domain.SupertrendConfig)(cfg.Supertrend),
		OrderConfig:      (domain.OrderConfig)(cfg.Order),
		Timeframe:        domain.Period(cfg.Timeframe),
	}
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

type botOrder struct {
	Symbol   string `json:"sym"`
	Quantity string `json:"qty"`
}

// -----

type botExchangeConfig struct {
	APIKey    string `json:"apiKey"`
	SecretKey string `json:"secretKey"`
}

func (r botExchangeConfig) ToBotConfigExchangeDomain() domain.BotExchange {
	return domain.BotExchange{
		APIKey:    r.APIKey,
		SecretKey: r.SecretKey,
	}
}

const (
	botExchangeCofigFileName = "botKeys.json"
	botConfigFileName        = "config.json"
)

type BotConfig struct{}

func NewBotConfig() *BotConfig {
	return &BotConfig{}
}

func (ind *BotConfig) UpdateBotConfig(in domain.BotConfig) error {
	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configFile := fmt.Sprintf("%s/%s", rootDir, botConfigFileName)
	f, err := os.Create(configFile)
	if err != nil {
		// TODO: handle proper error
		panic(err)
	}
	config := botConfig{
		EMA:        (emaConfig)(in.EMAConfig),
		MACD:       (macdConfig)(in.MACDConfig),
		RSI:        (rsiConfig)(in.RSIConfig),
		STO:        (stoConfig)(in.STOConfig),
		Supertrend: (supertrendConfig)(in.SupertrendConfig),
		Order:      (botOrder)(in.OrderConfig),
		Timeframe:  string(in.Timeframe),
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
	return nil
}

func (ind *BotConfig) RetrieveBotConfig() (domain.BotConfig, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configFile := fmt.Sprintf("%s/%s", rootDir, botConfigFileName)
	f, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("cannot read file from ", configFile)
		return domain.BotConfig{}, err
	}
	var cfg botConfig
	err = json.Unmarshal(f, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg.ToBotConfigDomain(), err
}

func (ind *BotConfig) RetrieveBotExchange() (domain.BotExchange, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configFile := fmt.Sprintf("%s/%s", rootDir, botExchangeCofigFileName)
	f, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("cannot read file from ", configFile)
		return domain.BotExchange{}, err
	}
	var cfg botExchangeConfig
	err = json.Unmarshal(f, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg.ToBotConfigExchangeDomain(), err
}

func (ind *BotConfig) UpdateBotExchange(in domain.BotExchange) error {
	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configFile := fmt.Sprintf("%s/%s", rootDir, botExchangeCofigFileName)
	f, err := os.Create(configFile)
	if err != nil {
		// TODO: handle proper error
		panic(err)
	}
	config := botExchangeConfig{
		APIKey:    in.APIKey,
		SecretKey: in.SecretKey,
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
	return nil
}
