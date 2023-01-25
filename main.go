package main

import (
	"github.com/gin-gonic/gin"
	"github.com/giwiro/escrap/config"
	"github.com/giwiro/escrap/logger"
	"github.com/giwiro/escrap/modules/version"
	versionWeb "github.com/giwiro/escrap/modules/version/web"
	log "github.com/sirupsen/logrus"
)

func setUpEngine() *gin.Engine {
	logger.SetUpLogger(log.DebugLevel)

	configErr := config.ReadConfig(".")

	// Production settings
	if config.Conf.Environment == "production" {
		logger.SetUpLogger(log.InfoLevel)
		gin.SetMode(gin.ReleaseMode)
	}

	if configErr != nil {
		log.Fatalln(configErr)
	}

	engine := gin.New()
	setTrustedProxiesErr := engine.SetTrustedProxies(nil)

	if setTrustedProxiesErr != nil {
		log.Fatalln(setTrustedProxiesErr)
	}

	return engine
}

func main() {
	engine := setUpEngine()

	mainRouter := engine.Group("/api")

	versionController := versionWeb.NewVersionController()
	versionController.RegisterRoutes(mainRouter)

	log.Infof("escrap v%s", version.Version)
	log.Infof("Listening on: %s", config.Conf.Server.Address)
	runErr := engine.Run(config.Conf.Server.Address)

	if runErr != nil {
		log.Fatalln(runErr)
		return
	}
}
