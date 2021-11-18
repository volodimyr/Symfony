package app

import (
	"os"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetUpLogger(level string) {
	lvl, err := strconv.Atoi(level)
	if err != nil {
		lvl = -1
	}
	log.Logger = zerolog.New(os.Stdout).Level(zerolog.Level(lvl)).With().Timestamp().Logger()
	log.Info().Str("chosen level", log.Logger.GetLevel().String()).Msg("log level set")
}

func LogRequest(c echo.Context, err error) {
	log.Err(err).
		Str("path", c.Request().Method+" "+c.Request().RequestURI).
		Int("status", c.Response().Status).
		Str("rid", c.Response().Header().Get(echo.HeaderXRequestID)).
		Msg("incoming request")
}
