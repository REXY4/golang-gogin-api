package src

import (
	"github.com/gin-gonic/gin"
)

func Server(_author string, _name string, _port string) {
	r := gin.New()
	port := ":" + _port
	r.Run(port)
}
