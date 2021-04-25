package gorm

import (
	"fmt"

	"github.com/akhmettolegen/currency-converter/pkg/config"
	//"github.com/akhmettolegen/currency-converter/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBManager struct {
	DB *gorm.DB
}

func NewDB(config *config.Config) *DBManager {
	connectionString := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v TimeZone=Asia/Almaty", config.DB.Host, config.DB.Port, config.DB.User, config.DB.Name, config.DB.Pass, config.DB.Mode)

	// db, err := gorm.Open("postgres", connectionString)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		panic(err)
	}

	return &DBManager{
		DB: db,
	}
}

func (m *DBManager) AutoMigrate() {
	fmt.Println("auto migrating...")
	m.DB.AutoMigrate(
	)
}
