package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var loadedConfig *Config

type sessionConfig struct {
	Key    string
	Domain string
}

type hostConfig struct {
	ServerName string
	Port       string
}

type databaseConfig struct {
	User       string
	Pass       string
	Host       string
	Port       string
	Name       string
	SslMode    string
	ConnString string
	URL        string
}

type sentryConfig struct {
	Dsn string
	Env string
}

// config represents a categorized structure of environment
// variables that are loaded for the application
type Config struct {
	Host     hostConfig
	Database databaseConfig
	Session  sessionConfig
	Sentry   sentryConfig
}

func Get() *Config {
	load()
	return loadedConfig
}

// getDbConnString constructs the postgres db connection string
// from the loaded database environment variables and populates it
// in the config object exposed through Get()
func getDbConnString() string {
	load()

	// Heroku provides DATABASE_URL. If this is provided, just use it.
	url := loadedConfig.Database.URL
	if url != "" {
		return url
	}

	// Format: postgres://postgres:password@localhost:5432/nova?sslmode=disable
	url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		loadedConfig.Database.User,
		loadedConfig.Database.Pass,
		loadedConfig.Database.Host,
		loadedConfig.Database.Port,
		loadedConfig.Database.Name,
		loadedConfig.Database.SslMode,
	)

	return url
}

// isLoaded determines if the config object has been loaded into memory
func isLoaded() bool {
	return loadedConfig != nil
}

// Load loads environment variables into the configuration object.
// Calling load multiple times will not throw an error
func load() {
	if isLoaded() {
		// configuration has already been loaded
		return
	}

	// load .env if one is available
	err := loadDotEnv()
	if err != nil {
		log.Println("no .env found, skipping")
	}

	// load settings and construct loadedConfig
	loadedConfig = &Config{
		Host: hostConfig{
			ServerName: os.Getenv("SERVER_NAME"),
			Port:       os.Getenv("SERVER_PORT"),
		},
		Database: databaseConfig{
			User:    os.Getenv("DATABASE_USER"),
			Pass:    os.Getenv("DATABASE_PASS"),
			Host:    os.Getenv("DATABASE_HOST"),
			Port:    os.Getenv("DATABASE_PORT"),
			Name:    os.Getenv("DATABASE_NAME"),
			SslMode: os.Getenv("DATABASE_SSL_MODE"),
			URL:     os.Getenv("DATABASE_URL"),
		},
		Session: sessionConfig{
			Key:    os.Getenv("SESSION_KEY"),
			Domain: os.Getenv("SESSION_COOKIE_DOMAIN"),
		},
		Sentry: sentryConfig{
			Dsn: os.Getenv("SENTRY_DSN"),
			Env: os.Getenv("SENTRY_ENV"),
		},
	}

	if loadedConfig.Host.Port == "" {
		// Fallback: Heroku injects PORT env variable
		loadedConfig.Host.Port = os.Getenv("PORT")
	}

	loadedConfig.Database.ConnString = getDbConnString()
}

func loadDotEnv() error {
	wd, err := os.Getwd()
	if err != nil {
		// couldn't find root path of the application, just die
		log.Fatal(err)
	}

	// read .env and load it into environment variables
	err = godotenv.Load(wd + "/.env")
	if err != nil {
		return err
	}

	return nil
}
