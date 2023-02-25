package config

import (
	"log"

	"github.com/subosito/gotenv"
)

type Configuration struct {
	App *AppConfig
}

func init() {
	err := gotenv.Load()
	if err != nil {
		log.Fatal("Failed to Load env")
	}
	//app config
	MainAppConfig()
}

func Configs() *Configuration {
	app := MainAppConfig()
	return &Configuration{App: app}
}
