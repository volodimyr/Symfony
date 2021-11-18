package app

import (
	"github.com/labstack/echo/v4"
)

func MiddlewareLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			LogRequest(c, nil)
		}

		return err
	}
}
