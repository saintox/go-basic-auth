package main

import (
	"github.com/labstack/echo/v4"
	"github.com/saintox/go-basic-auth/controllers"
	"github.com/saintox/go-basic-auth/pkg/databases"
	log "github.com/saintox/go-basic-auth/pkg/logger"
	"github.com/saintox/go-basic-auth/services"
)

func main() {
	app := echo.New()

	// log configuration
	logger := log.Configure(log.Config{
		EnableDebugMode:      false,
		EnableConsoleLogging: true,
		EncodeLogsAsJson:     true,
		EnableFileLogging:    false,
		FileDir:              "logs",
		FileName:             "system-log",
		MaxAge:               1,
		MaxLogSize:           0,
	})

	_, err := databases.CreateMySqlClient()
	if err != nil {
		logger.Fatal().Err(err).Msg("Database connection failed")
		panic(app)
	}

	ctl := controllers.NewController(services.NewService())

	app.GET("/", ctl.Auth.Login)

	err = app.Start(":8000")
	if err != nil {
		return
	}
}
