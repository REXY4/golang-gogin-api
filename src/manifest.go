package src

import "github.com/REXY4/golang-gogin-api/config"

func setup() {
	app := config.Configs().App
	Server(app.Author, app.Name, app.Port)
}

func Start() {
	setup()
}
