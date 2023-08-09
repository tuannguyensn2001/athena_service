package infra

import (
	"athena_service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Infra struct {
	Db *gorm.DB
}

func Get(config config.Config) (Infra, error) {
	result := Infra{}
	db, err := gorm.Open(postgres.Open(config.DbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return result, err
	}

	result.Db = db

	return result, nil
}
