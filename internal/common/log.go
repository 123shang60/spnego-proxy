package common

import "github.com/sirupsen/logrus"

func SetLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, DisableColors: false})
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
}
