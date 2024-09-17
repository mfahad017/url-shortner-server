// contains go code for database connection and operations

package database

import (
	"fmt"
	"log"
	"log/slog"
	"server/src/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() error {

	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.GetConfig().DBHost, config.GetConfig().DBUser, config.GetConfig().DBPass, config.GetConfig().DBName, config.GetConfig().DBPort)

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect to database")
		return err
	}
	slog.Info("Connected to database")
	return nil
}

func DisconnectDB() {
	// Close the connection
	pg, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to disconnect from database")
	}
	pg.Close()
}
