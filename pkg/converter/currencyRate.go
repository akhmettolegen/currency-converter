package converter

import (
	"encoding/xml"
	"github.com/akhmettolegen/currency-converter/pkg/model"
	"io/ioutil"
	"net/http"
)

type CurrencyRate struct {
	Rate *map[model.Currency]float64
}

func GetCurrencyRate() (*CurrencyRate, error) {

	rates := map[model.Currency]float64{"KZT": 1}
	xmlPath := "https://nationalbank.kz/rss/rates_all.xml?switch=kazakh"

	resp, err := http.Get(xmlPath)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)
	var xmlRates model.XMLRates
	xml.Unmarshal(byteValue, &xmlRates)

	for i := 0; i < len(xmlRates.Channel.Item); i++ {
		rates[xmlRates.Channel.Item[i].Title] = xmlRates.Channel.Item[i].Description
	}

	return &CurrencyRate {
		Rate: &rates,
	}, nil
}
