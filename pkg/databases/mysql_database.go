package databases

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/saintox/go-basic-auth/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func CreateMySqlClient() (*gorm.DB, error) {
	var (
		connection *gorm.DB
		err        error
		ctx        context.Context
	)

	err = godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	database := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASS"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_NAME"),
	)

	if connection, err = gorm.Open(mysql.Open(database), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	}); err != nil {
		return nil, err
	}

	// migrate model
	migrateModels := []interface{}{
		&entities.User{},
	}

	tx := connection.WithContext(ctx)
	if err = tx.AutoMigrate(migrateModels...); err != nil {
		return connection, err
	}

	return connection, err
}
