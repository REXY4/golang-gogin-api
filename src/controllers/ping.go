package controllers

import (
	"net/http"

	"github.com/REXY4/golang-gogin-api/src/helpers"
	"github.com/gin-gonic/gin"
)

type PingController interface {
	GetPing(c *gin.Context)
}

type response struct {
	Title   string
	Message string
}

func GetPing(c *gin.Context) {
	data := &response{
		Title:   "check ping",
		Message: "Ok",
	}
	res := helpers.BuildResponse("success get ping Apis", data)
	c.JSON(http.StatusOK, res)
}
