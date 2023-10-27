package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Start(bind string) error {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	logrus.Info("Starting API on " + bind)

	return e.Start(bind)
}
