package config

import (
	"log"
	"os"
)

type AppConfig struct {
	Name string
	Port string
}

func ConfigApp() *AppConfig {
	name := os.Getenv("APP_NAME")
	port := os.Getenv("APP_PORT")
	log.Printf("ConfigApp: name=%s, port=%s", name, port)

	// Membuat objek AppConfig dan mengembalikannya
	return &AppConfig{
		Name: name,
		Port: port,
	}
}
