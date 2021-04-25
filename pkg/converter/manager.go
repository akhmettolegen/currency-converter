package converter

import (
	"fmt"
	"github.com/akhmettolegen/currency-converter/pkg/config"
	"github.com/akhmettolegen/currency-converter/pkg/gorm"
	"github.com/akhmettolegen/currency-converter/pkg/model"
	"github.com/robfig/cron/v3"
)

type Manager struct {
	Config *config.Config
	CurrencyRate *map[model.Currency]float64
	DB *gorm.DBManager
}

func Get() *Manager{
	config := config.Get()
	rates, err := GetCurrencyRate()
	if err != nil {
		return nil
	}
	go func() {
		c := cron.New(cron.WithSeconds())
		spec := "0 */1 * * * *"
		c.AddFunc(spec, func() {
			fmt.Println("cron running")
			rates, err = GetCurrencyRate()
			if err != nil {
				return
			}
		})
		// c.Start()
		c.Run()
		select {}
	}()

	db := gorm.NewDB(config)
	return &Manager{
		Config: config,
		CurrencyRate: rates.Rate,
		DB: db,
	}
}