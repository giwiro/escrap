package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func SetUpLogger(level log.Level) {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(level)
}
