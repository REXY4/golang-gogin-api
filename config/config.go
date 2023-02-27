package config

import (
	"encoding/json"
	"log"

	"github.com/subosito/gotenv"
)

type Configuration struct {
	App *AppConfig `json:"app"`
	Sql *SqlConfig `json:"sql"`
}

func init() {
	err := gotenv.Load()
	if err != nil {
		log.Fatal("Failed to Load env")
	}
}

func Configs() (string, error) {
	data := &Configuration{
		App: MainAppConfig(),
		Sql: ConfigSql(),
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Panic(err)
	}
	return string(jsonData), nil
}
