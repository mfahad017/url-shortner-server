// config/config.go
package config

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

var allowedEnvs = []string{"dev", "prod"}

type conf struct {
	ENV                          string
	Port                         string `validate:"-"`
	DBHost                       string `validate:"required"`
	DBUser                       string `validate:"required"`
	DBPass                       string `validate:"required"`
	DBPort                       string `validate:"required"`
	DBName                       string `validate:"required"`
	JWT_SECRET                   string `validate:"required"`
	AUTH_TOKEN_EXPIRY_IN_MINUTES int    `validate:"required"`
	DOMAIN_PREFIX                string `validate:"required"`
}

var config *conf

func verifyEnv(env string) error {

	if env == "" {
		return errors.New("ENV not set")
	}

	for _, allowedEnv := range allowedEnvs {
		if env == allowedEnv {
			return nil
		}
	}
	return errors.New("ENV not valid. Allowed values are " + fmt.Sprint(allowedEnvs))
}

// Load env variables from .env file
// and validate them. Exit if validation fails
func LoadConfig() {

	env := os.Getenv("ENV")

	envError := verifyEnv(env)

	if envError != nil {
		log.Fatal(envError) // default to dev if APP_ENV is not set
	}

	fmt.Println("Loading env: ", env)

	envFile := ".env." + env

	slog.Info("Loading config from " + envFile)

	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatalf("Error loading .env.%s file", env)
	}

	authTokenExpiry, err := strconv.Atoi(os.Getenv("AUTH_TOKEN_EXPIRY_IN_MINUTES"))
	if err != nil {
		log.Fatal(err)
	}

	config = &conf{
		Port:                         getEnv(os.Getenv("PORT"), "8080"),
		DBHost:                       os.Getenv("DB_HOST"),
		DBUser:                       os.Getenv("DB_USER"),
		DBPass:                       os.Getenv("DB_PASS"),
		DBPort:                       getEnv(os.Getenv("DB_PORT"), "5432"),
		DBName:                       os.Getenv("DB_NAME"),
		JWT_SECRET:                   os.Getenv("JWT_SECRET"),
		AUTH_TOKEN_EXPIRY_IN_MINUTES: authTokenExpiry,
		DOMAIN_PREFIX:                os.Getenv("DOMAIN_PREFIX"),
	}

	validate := validator.New()
	err = validate.Struct(*config)

	if err != nil {
		log.Fatal(err)
	}
}

// get copy of the env file variable to avoid direct mutation
func GetConfig() conf {
	return *config
}

func getEnv(val string, defaultValue string) string {

	if val == "" {
		return defaultValue
	}
	return val
}
