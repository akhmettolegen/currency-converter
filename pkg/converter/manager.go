package converter

import "github.com/akhmettolegen/currency-converter/pkg/config"

type Manager struct {
	Config *config.Config
}

func Get() *Manager{
	config := config.Get()

	return &Manager{
		Config: config,
	}
}