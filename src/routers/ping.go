package routers

import (
	"github.com/REXY4/golang-gogin-api/src/controllers"
	"github.com/gin-gonic/gin"
)

type defenition struct {
	Method string
	Path   string
	Action []func(c *gin.Context)
}

var (
	getPing = controllers.GetPing
)

func PingDefenition() []defenition {
	def := []defenition{
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
