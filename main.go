package main

import (
	"github.com/sirupsen/logrus"
	"github.com/vcoltfre/godo/src/api"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {
	if err := api.Start(":8080"); err != nil {
		logrus.WithError(err).Fatal("Failed to start API")
	}
}
