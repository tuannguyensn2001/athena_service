package infra

import (
	"athena_service/config"
	"github.com/dgraph-io/badger/v3"
	"github.com/pusher/pusher-http-go/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Infra struct {
	Db     *gorm.DB
	Pusher *pusher.Client
	Badger *badger.DB
}

func Get(config config.Config) (Infra, error) {
	result := Infra{}
	db, err := gorm.Open(postgres.Open(config.DbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return result, err
	}

	pusher := &pusher.Client{
		AppID:   "1651032",
		Key:     "f870f614902b26101ff8",
		Secret:  "41cf0c48b5f254b3c51d",
		Cluster: "ap1",
		Secure:  true,
	}
	bd, err := badger.Open(badger.DefaultOptions("tmp/badger"))
	if err != nil {
		return result, err
	}

	result.Db = db
	result.Pusher = pusher
	result.Badger = bd

	return result, nil
}
