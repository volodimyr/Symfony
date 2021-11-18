package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/volodimyr/Symphony/api/http"
	"github.com/volodimyr/Symphony/app"

	_ "github.com/lib/pq"
)

//go:generate sqlboiler --add-global-variants psql

func main() {
	app.SetUpLogger(os.Getenv("LOG_LEVEL"))

	server := app.NewEchoServer()

	c := app.NewConfig()
	d, err := app.NewDependencies(c)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init dependencies")
	}
	routes := http.NewRouter(d)
	routes.Attach(server)

	server.Logger.Fatal(server.Start(":" + c.Server.Port))
}
