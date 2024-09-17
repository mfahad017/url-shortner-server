package main

import (
	"log/slog"
	"server/src/config"
	"server/src/database"
	"server/src/models"
)

func init() {

	config.LoadConfig()

	database.ConnectToDB()
}

func main() {

	slog.Info("Running migrations")

	slog.Info("Migrating User model")
	err := database.DB.AutoMigrate(&models.User{})

	if err != nil {
		slog.Error("Failed to migrate User model")
	}

	slog.Info("Migrating URL model")
	err = database.DB.AutoMigrate(&models.URL{})

	if err != nil {
		slog.Error("Failed to migrate URL model")
	}

	slog.Info("Migrations complete")

}
