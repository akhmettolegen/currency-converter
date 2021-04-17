package converter

import (
	"fmt"
	"github.com/akhmettolegen/currency-converter/pkg/model"
)

func (m *Manager)Converter(body *model.RequestBody) *model.ResponseBody {

	result := body.Amount * 100
	response := &model.ResponseBody{
		Result: fmt.Sprintf("%.2f %s", result, body.To),
	}

	return response
}
