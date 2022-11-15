package repository

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/brightnc/not-human-trading/internal/core/domain"
)

const (
	indicatorConfigFileName = "config.csv"
)

type Indicator struct{}

func NewIndicator() *Indicator {
	return &Indicator{}
}

func (ind *Indicator) UpdateIndicator(indicator domain.IndicatorConfig) error {
	f, err := os.Open(indicatorConfigFileName)
	if err != nil {
		// TODO: handle proper error
		panic(err)
	}
	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		// TODO: handle proper error
		panic(err)
	}

	for _, record := range records {
		fmt.Println(record)
	}
	return nil
}
