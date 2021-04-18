package converter

import (
	"encoding/json"
	"github.com/akhmettolegen/currency-converter/pkg/model"
	"io/ioutil"
	"os"
)

type CurrencyRate struct {
	Rate *map[model.Currency]float64
}

func GetCurrencyRate() *CurrencyRate {
	codePath := os.Getenv("GOPATH") + "/src/github.com/akhmettolegen/currency-converter/pkg/converter/currency-rate.json"
	jsonFile, err := os.Open(codePath)
	rates := map[model.Currency]float64{}
	if err == nil {
		defer jsonFile.Close()
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			return nil
		}

		if err = json.Unmarshal(byteValue, &rates); err != nil {
			return nil
		}
	}
	return &CurrencyRate {
		Rate: &rates,
	}
}
