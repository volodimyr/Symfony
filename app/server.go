package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func NewEchoServer() *echo.Echo {
	server := echo.New()
	server.HideBanner = true
	server.Use(middleware.RequestID())
	server.Use(middleware.Recover())
	server.Use(middleware.Secure())

	server.Use(MiddlewareLogger)
	server.HTTPErrorHandler = errHandler

	return server
}

func errHandler(err error, c echo.Context) {
	defer LogRequest(c, err)

	cErr, ok := err.(interface {
		Code() int
		Message() string
	})

	switch {
	case ok:
		err = c.JSON(cErr.Code(), map[string]interface{}{
			"message": cErr.Message(),
		})
		if err != nil {
			log.Err(err).Msg("failed to marshal JSON as a custom HTTP error")
		}

	default:
		c.Echo().DefaultHTTPErrorHandler(err, c)
	}
}
