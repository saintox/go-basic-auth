package main

import (
	"github.com/labstack/echo/v4"
	"github.com/saintox/go-basic-auth/controllers"
	"github.com/saintox/go-basic-auth/services"
)

func main() {
	app := echo.New()
	ctl := controllers.NewController(services.NewService())

	app.GET("/", ctl.Auth.Login)

	app.Logger.Fatal(app.Start(":8000"))
}
