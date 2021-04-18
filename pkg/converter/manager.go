package converter

import (
	"github.com/akhmettolegen/currency-converter/pkg/config"
	"github.com/akhmettolegen/currency-converter/pkg/model"
)

type Manager struct {
	Config *config.Config
	CurrencyRate *map[model.Currency]float64
}

func Get() *Manager{
	config := config.Get()
	rates := GetCurrencyRate()

	return &Manager{
		Config: config,
		CurrencyRate: rates.Rate,
	}
}