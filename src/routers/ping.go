package routers

import (
	"github.com/REXY4/golang-gogin-api/src/controllers"
	"github.com/gin-gonic/gin"
)

var (
	getPing = controllers.GetPing
)

func PingDefenition() []Definition {
	def := []Definition{
		{
			Method: "GET",
			Path:   "/ping",
			Action: []func(c *gin.Context){
				getPing,
			},
		},
	}
	return def
}
