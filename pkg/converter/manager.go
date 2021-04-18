package converter

import (
	"encoding/json"
	"github.com/akhmettolegen/currency-converter/pkg/config"
	"github.com/akhmettolegen/currency-converter/pkg/model"
	"io/ioutil"
	"os"
)

type Manager struct {
	Config *config.Config
	CurrencyRate *map[model.Currency]float64
}

func Get() *Manager{
	config := config.Get()

	codePath := os.Getenv("GOPATH") + "/src/github.com/akhmettolegen/currency-converter/pkg/converter/currency-rate.json"
	jsonFile, err := os.Open(codePath)
	codes := map[model.Currency]float64{}
	if err == nil {
		defer jsonFile.Close()
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			return nil
		}

		if err = json.Unmarshal(byteValue, &codes); err != nil {
			return nil
		}
	}

	return &Manager{
		Config: config,
		CurrencyRate: &codes,
	}
}