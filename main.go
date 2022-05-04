package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest-api/api"
	"rest-api/app/modules"
	"rest-api/config"
	"rest-api/util"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetConfig()
	dbCon := util.NewConnectionDatabase(config)

	controllers := modules.RegisterModules(dbCon, config)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	api.RegisterRoutes(e, &controllers)

	go func() {
		addres := fmt.Sprintf(":%d", config.App.Port)
		if err := e.Start(addres); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

}
