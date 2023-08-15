package config

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var keys = []string{"DATABASE_URL", "PORT", "APP_NAME"}

type Config struct {
	DbUrl   string `mapstructure:"DATABASE_URL"`
	Port    string `mapsctructure:"PORT"`
	AppName string `mapstructure:"APP_NAME"`
}

func Get() (Config, error) {
	var result Config
	path, err := os.Getwd()
	if err != nil {
		return result, err
	}

	log.Info().Str("path", path).Msg("config path")

	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()
	viper.ReadInConfig()

	for _, key := range keys {
		viper.BindEnv(key)
	}

	if err := viper.Unmarshal(&result); err != nil {
		return result, err
	}

	return result, nil
}
