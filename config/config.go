package config

import (
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

const DEVELOPMENT = "development"
const PRODUCTION = "production"
const LOCAL = "local"
const STAGING = "staging"
const DEFAULT_PORT = "9035"
const APP_NAME = "go-template"

type ConfigVariables struct {
	AppPort string
	AppName string
	Env     string
	DbUrl   string
	Redis   struct {
		Host     string
		Password string
		Db       int
	}
}

// the config variable to be used appwide
var AppConfig = ConfigVariables{
	AppPort: DEFAULT_PORT,
	Env:     DEVELOPMENT,
	AppName: APP_NAME,
}

func LoadConfig() {
	if envPort := os.Getenv("PORT"); envPort != "" {
		AppConfig.AppPort = envPort
	}

	if os.Getenv("APP_ENV") != "" {
		AppConfig.Env = strings.ToLower(os.Getenv("APP_ENV"))
	}

	if os.Getenv("APP_NAME") != "" {
		AppConfig.AppName = strings.ToLower(os.Getenv("APP_NAME"))
	}

	setDatabaseConfig()
	setRedisConfig()
}

func setDatabaseConfig() {
	if os.Getenv("DB_URL") != "" {
		AppConfig.DbUrl = strings.ToLower(os.Getenv("DB_URL"))
	}
}

func IsDev() bool {
	if AppConfig.Env == DEVELOPMENT || AppConfig.Env == LOCAL {
		return true
	}
	return false
}

func IsProd() bool {
	return AppConfig.Env == PRODUCTION
}

func IsStaging() bool {
	return AppConfig.Env == STAGING
}

func setRedisConfig() {
	AppConfig.Redis.Host = os.Getenv("REDIS_HOST")
	AppConfig.Redis.Password = os.Getenv("REDIS_PASSWORD")
	AppConfig.Redis.Db = 0
}

/*
 * Load app config at Package init
 */
func init() {
	LoadConfig()
}
