package cmd

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "inventory/docs"
	"inventory/server"
	"inventory/server/errorhandler"
	"os"
)

func Start() {
	if err := godotenv.Load(); err != nil {
		panic(".env file not found")
	}
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.HTTPErrorHandler = errorhandler.HttpEchoCustomError

	e.GET("/swagger/*any", echoSwagger.WrapHandler)

	server.InitService(e)

	log.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
