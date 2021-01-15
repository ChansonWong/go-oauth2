package main

import (
	"github.com/gin-gonic/gin"
	"server/dao"
	"server/handler"
	"server/model"
)

var router *gin.Engine

func main() {
	dao.GetDB().AutoMigrate(model.User{}, model.AuthRefreshToken{}, model.AuthAccessToken{}, model.AuthClientDetail{})

	router = gin.Default()
	handler.ConfigRouter(router)
	router.Run(":9090")
}
