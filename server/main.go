package main

import (
	"github.com/gin-gonic/gin"
	"server/handler"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	handler.ConfigRouter(router)
	router.Run(":9090")
}
