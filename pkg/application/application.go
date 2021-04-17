package application

import (
	"github.com/akhmettolegen/currency-converter/pkg/config"
	"github.com/akhmettolegen/currency-converter/pkg/converter"
)

type Application struct {
	Config *config.Config
	Manager *converter.Manager
}

func Get() (*Application, error) {
	config := config.Get()
	manager := converter.Get()

	return &Application{
		Config: config,
		Manager: manager,
	}, nil
}
