package src

import (
	"log"

	"github.com/REXY4/golang-gogin-api/config"
)

func setup(config string) {
	SetupServer(config)
}

func Start() {
	configStr, err := config.Configs()
	if err != nil {
		log.Panic(err)
	}
	setup(configStr)
}
