package application

import "github.com/akhmettolegen/currency-converter/pkg/config"

type Application struct {
	Config *config.Config
}

func Get() (*Application, error) {
	config := config.Get()
	return &Application{
		Config: config,
	}, nil
}
