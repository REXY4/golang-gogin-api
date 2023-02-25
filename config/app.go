package config

import (
	"log"
	"os"
)

type AppConfig struct {
	Author string
	Name   string
	Port   string
}

func MainAppConfig() *AppConfig {
	author := os.Getenv("AUTHOR")
	name := os.Getenv("APP_NAME")
	port := os.Getenv("APP_PORT")
	log.Printf("ConfigApp: name=%s, port=%s", name, port)

	// Membuat objek AppConfig dan mengembalikannya
	return &AppConfig{
		Author: author,
		Name:   name,
		Port:   port,
	}
}
