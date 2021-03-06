package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"rest-api/api"
	"rest-api/app/modules"
	"rest-api/config"
	"rest-api/util"
	"time"

	_ "rest-api/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Movie API
// @version 1.0
// @description Brikut API yang digunakan untuk memanage movie data
func main() {
	config := config.GetConfig()
	dbCon := util.NewConnectionDatabase(config)

	controllers := modules.RegisterModules(dbCon, config)

	e := echo.New()
	handleSwag := echoSwagger.WrapHandler

	// error heroku
	// e.GET("/swagger/*", handleSwag)

	// api.RegisterRoutes(e, &controllers)

	// go func() {
	// 	addres := fmt.Sprintf(":%d", config.App.Port)
	// 	if err := e.Start(addres); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	e.GET("/swagger/*", handleSwag)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK!!")
	})

	api.RegisterRoutes(e, &controllers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8084"
	}
	e.Logger.Fatal(e.Start(":" + port))

	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	// <-quit

	defer dbCon.CloseConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

}
