package controllers

import (
	"net/http"

	"github.com/REXY4/golang-gogin-api/src/helpers"
	"github.com/gin-gonic/gin"
)

type PingController interface {
	GetPing(c *gin.Context)
}

func GetPing(c *gin.Context) {
	res := helpers.BuildResponse("Ping ok")
	c.JSON(http.StatusOK, res)
}
