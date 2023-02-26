package src

import (
	"encoding/json"

	"github.com/REXY4/golang-gogin-api/src/routers"
	"github.com/gin-gonic/gin"
)

func Server(manifest string) {
	r := gin.Default()
	var config map[string]map[string]string
	err := json.Unmarshal([]byte(manifest), &config)
	if err != nil {
		panic(err.Error())
	}
	routes := routers.Router()
	for _, route := range routes() {
		for _, routes := range route.Action {
			r.Handle(route.Method, route.Path, routes)
		}
	}
	port := ":" + config["app"]["port"]
	r.Run(port)
}
