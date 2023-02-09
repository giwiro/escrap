package config

import (
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Environment string
	Database    struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}
	Server struct {
		Address string
		Ttl     int64
	}
	Scrapper struct {
		Amazon struct {
			AssociateTag string `mapstructure:"associate_tag"`
			AccessKey    string `mapstructure:"access_key"`
			SecretKey    string `mapstructure:"secret_key"`
		}
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
