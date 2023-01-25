package config

import (
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Environment string
	Database    struct {
		User     string
		Password string
		Address  string
		DBName   string
	}
	Server struct {
		Address string
	}
}

var Conf Config

func ReadConfig(path string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		return err
	}

	log.Infof("Loading config from: %s", path)
	spew.Dump(Conf)

	return nil
}
