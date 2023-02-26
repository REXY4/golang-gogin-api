package routers

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
	R *gin.Engine
}

type route struct {
	Ping func() []defenition
}

func Router() []route {
	routes := []route{
		{
			Ping: PingDefenition,
		},
	}
	return routes
}
