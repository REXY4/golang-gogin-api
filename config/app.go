package config

import (
	"os"
)

type AppConfig struct {
	Author string `json:"author"`
	Name   string `json:"name"`
	Port   string `json:"port"`
}

func MainAppConfig() *AppConfig {
	cfg := &AppConfig{
		Author: os.Getenv("AUTHOR"),
		Name:   os.Getenv("APP_NAME"),
		Port:   os.Getenv("APP_PORT"),
	}
	return cfg
}
