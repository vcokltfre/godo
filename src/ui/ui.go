package ui

import (
	_ "embed"

	"github.com/labstack/echo/v4"
)

//go:embed index.html
var data string

func GetIndex(c echo.Context) error {
	return c.HTML(200, data)
}
