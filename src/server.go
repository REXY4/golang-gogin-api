package src

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func Server(manifest string) {
	r := gin.Default()
	var config map[string]map[string]string
	err := json.Unmarshal([]byte(manifest), &config)
	if err != nil {
		panic(err.Error())
	}
	port := ":" + config["app"]["port"]
	r.Run(port)
}
