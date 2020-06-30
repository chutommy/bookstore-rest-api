package config

import (
	"net/http"
	"os"
	"time"
)

// Config is a application config struct.
type Config struct {
	Log *LogConfig
	Srv *ServerConfig
	DB  *DBConfig
}

// set config
var cfg = &Config{

	// log
	Log: &LogConfig{
		Output: os.Stdout,
	},

	// srv
	Srv: &ServerConfig{
		Addr:              ":8081",
		ReadTimeout:       3 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       10 * time.Second,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	},

	// db
	DB: &DBConfig{
		Host:     "localhost",
		Port:     5432,
		DBName:   "project_01",
		User:     "user_00",
		Password: "159258",
	},
}

// GetConfig returns application config struct.
func GetConfig() *Config {
	return cfg
}