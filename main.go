package main

import (
	"github.com/labstack/echo/v4"
	"github.com/saintox/go-basic-auth/controllers"
	"github.com/saintox/go-basic-auth/middlewares"
	"github.com/saintox/go-basic-auth/pkg/databases"
	log "github.com/saintox/go-basic-auth/pkg/logger"
	"github.com/saintox/go-basic-auth/repositories"
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

	// db initialization
	db, err := databases.CreateMySqlClient()
	if err != nil {
		logger.Fatal().Err(err).Msg("Database connection failed")
		panic(app)
	}

	// repository, service, controller initialization
	repo := repositories.NewRepository(db)
	srv := services.NewService(repo)
	validator := middlewares.NewValidator()
	ctl := controllers.NewController(srv, validator, logger)

	auth := app.Group("auth")
	auth.POST("/login", ctl.Auth.Login)
	auth.POST("/register", ctl.Auth.Register)

	// ping route
	app.GET("ping", func(c echo.Context) error {
		return c.NoContent(200)
	})

	err = app.Start(":8000")
	if err != nil {
		return
	}
}
