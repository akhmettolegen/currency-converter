package converter

import (
	"fmt"
	"github.com/akhmettolegen/currency-converter/pkg/model"
)

func (m *Manager) Converter(body *model.RequestBody) *model.ResponseBody {

	codes := *m.CurrencyRate

	response := &model.ResponseBody{
		EUR: fmt.Sprintf("%.4f %s", body.Amount * codes[body.Currency] / codes[model.CurrencyEUR], model.CurrencyEUR),
		USD: fmt.Sprintf("%.4f %s", body.Amount * codes[body.Currency] / codes[model.CurrencyUSD], model.CurrencyUSD),
		RUB: fmt.Sprintf("%.4f %s", body.Amount * codes[body.Currency] / codes[model.CurrencyRUB], model.CurrencyRUB),
		KZT: fmt.Sprintf("%.4f %s", body.Amount * codes[body.Currency] / codes[model.CurrencyKZT], model.CurrencyKZT),
	}

	return response
}
