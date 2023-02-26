package routers

import (
	"github.com/REXY4/golang-gogin-api/src/controllers"
)

var (
	getPing = controllers.GetPing
)

func PingDefenition() []Definition {
	def := []Definition{
		{
			Method: "GET",
			Path:   "/ping",
			Action: getPing,
		},
	}
	return def
}
