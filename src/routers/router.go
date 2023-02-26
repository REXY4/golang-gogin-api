package routers

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
	R *gin.Engine
}

func Router() {
	PingDefenition()
}
