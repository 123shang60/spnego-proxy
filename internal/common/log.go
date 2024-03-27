package common

import (
	"fmt"

	"github.com/123shang60/spnego-proxy/internal/config"
	"github.com/sirupsen/logrus"
)

func SetLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, DisableColors: false})

	level, err := logrus.ParseLevel(config.C.Log.Level)
	if err != nil {
		fmt.Printf("Couldn't parse log level: %s \n", config.C.Log.Level)
		level = logrus.InfoLevel
	}

	if level == logrus.DebugLevel {
		logrus.SetReportCaller(true)
	}

	logrus.SetLevel(level)
}
