package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("SERVER_NAME", "localhost")
	os.Setenv("SERVER_PORT", "8080")

	os.Setenv("DATABASE_HOST", "localhost")
	os.Setenv("DATABASE_PORT", "5432")
	os.Setenv("DATABASE_USER", "postgres")
	os.Setenv("DATABASE_PASS", "password")
	os.Setenv("DATABASE_NAME", "nova")
	os.Setenv("DATABASE_SSL_MODE", "disable")

	os.Setenv("SESSION_KEY", "super-secret")

	cfg := Get()

	if cfg.Host.ServerName != "localhost" {
		t.Error("expected host server name to be servername")
	}

	if cfg.Host.Port != "8080" {
		t.Error("expected host server port to be serverport")
	}

	connString := "postgres://postgres:password@localhost:5432/nova?sslmode=disable"
	if cfg.Database.ConnString != connString {
		t.Errorf("database connection string environment variables not loaded properly")
	}

	if cfg.Session.Key != "super-secret" {
		t.Error("expected session key to be super-secret")
	}
}
