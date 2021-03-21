package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/theflyingcodr/things/data"
	"github.com/theflyingcodr/things/data/inmemory"
	"github.com/theflyingcodr/things/transports/rest"
	"github.com/theflyingcodr/things/data/mysql"
	"github.com/theflyingcodr/things/service"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover(),
		middleware.Logger(),
		middleware.TimeoutWithConfig(
			middleware.TimeoutConfig{
				ErrorMessage: "request timed out",
				Timeout:      10 * time.Second,
			}))
	g := e.Group("api/v1")
	// TODO: you would init and pass an actual db to this.
	rest.NewEndpoint(
		service.NewThing(
			data.NewThingFacade(inmemory.NewThing(), mysql.NewThing(&sql.DB{})))).
		RegisterRoutes(g)
	log.Fatal(e.Start(":3131"))
}
