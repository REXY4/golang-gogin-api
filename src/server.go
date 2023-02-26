package src

import (
	"encoding/json"
	"log"

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
	// routes := r.Router()
	for _, route := range routers.Router() {
		log.Println("ini route", route)
	}
	port := ":" + config["app"]["port"]
	r.Run(port)
}

// // Get the Gin router from the routers package
// router := routers.Router()

// // Register the routes defined in the router with the Gin engine
// routes := r.Routes()
// for _, route := range router.Routes() {
// 	r.Handle(route.Method, route.Path, route.HandlerFuncs...)
// }
