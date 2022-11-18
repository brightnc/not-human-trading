package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/brightnc/not-human-trading/internal/core/domain"
)

type botConfig struct {
	ApiKey    string `json:"api_key"`
	SecretKey string `json:"secret_key"`
}

const (
	botConfiggFileName = "botKeys.json"
)

type BotConfigKey struct{}

func NewBotConfigKey() *BotConfigKey {
	return &BotConfigKey{}
}

func (r *BotConfigKey) UpdateBotConfig(bot domain.BotConfig) error {
	_, fileName, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(fileName)
	fmt.Println(currentDir)

	configFile := fmt.Sprintf("%s/%s", currentDir, botConfiggFileName)
	f, err := os.Create(configFile)
	if err != nil {
		// TODO: handle proper error
		panic(err)
	}
	config := botConfig{
		ApiKey:    bot.ApiKey,
		SecretKey: bot.SecretKey,
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
