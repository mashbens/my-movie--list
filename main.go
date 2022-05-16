package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest-api/api"
	"rest-api/app/modules"
	"rest-api/config"
	"rest-api/util"
	"time"

	// for ex
	dbMov "rest-api/business/movie/entity"
	dbUser "rest-api/business/user/entity"

	_ "rest-api/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Swagger Movie API
// @version 1.0
// @description Brikut API yang digunakan untuk memanage movie data
func main() {
	InitDB()
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

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	defer dbCon.CloseConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

}

var DB *gorm.DB

func InitDB() {

	dsn := "host=ec2-3-229-11-55.compute-1.amazonaws.com user=zzrjjedyarnucl password=2706b8a43ee46e5f83d361b0bfd931bbd80a42ade11cfbe90d45cd6362f790a9 dbname=dd29tj5s6ptpvg port=5432 sslmode=require TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&dbUser.User{}, &dbMov.Movie{})
}
