package src

import (
	"github.com/REXY4/golang-gogin-api/config"
	"github.com/gin-gonic/gin"
)

func Start() {
	cfg := config.ConfigApp()
	r := gin.Default()
	port := ":" + cfg.Port
	r.Run(port)
}
