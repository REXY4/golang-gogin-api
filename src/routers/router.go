package routers

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
	R *gin.Engine
}

type Definition struct {
	Method string                 `json:"method"`
	Path   string                 `json:"path"`
	Action []func(c *gin.Context) `json:"action"`
}

func Router() func() []Definition {
	routes := []func() []Definition{
		PingDefenition,
	}
	return func() []Definition {
		result := []Definition{}
		for _, r := range routes {
			result = append(result, r()...)
		}
		return result
	}
}
