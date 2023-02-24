package main

import (
	"log"

	server "github.com/REXY4/golang-gogin-api/src"
	"github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load()
	if err != nil {
		log.Fatalln("Failed to load env file")
	}
}

func main() {
	server.Start()
}
