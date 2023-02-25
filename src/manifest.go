package src

import (
	"log"

	"github.com/REXY4/golang-gogin-api/config"
)

func setup() {
	configStr, err := config.Configs()
	if err != nil {
		log.Panic(err)
	}
	Server(configStr)
}

func Start() {
	setup()
}
