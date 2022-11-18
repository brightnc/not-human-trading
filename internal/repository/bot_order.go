package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/brightnc/not-human-trading/internal/core/domain"
)

const (
	botOrderConfig string = "bot_order.json"
)

type botOrder struct {
	Symbol   string  `json:"sym"`
	Quantity float64 `json:"qty"`
}

type BotOrder struct{}

func NewBotOrder() *BotOrder {
	return &BotOrder{}
}

// UpdateBotOrder ...
// bot configurations

func (r *BotOrder) UpdateBotOrder(in domain.BotOrder) error {
	_, fileName, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(fileName)
	fmt.Println(currentDir)

	configFile := fmt.Sprintf("%s/%s", currentDir, botOrderConfig)
	f, err := os.Create(configFile)
	if err != nil {
		// TODO: handle proper error
		panic(err)
	}
	config := (botOrder)(in)
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
