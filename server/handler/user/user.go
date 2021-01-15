package user

import "github.com/gin-gonic/gin"

func LoadUser(router *gin.Engine) {
	// 登录页面
	router.GET("/login", Login)
	// 校验用户名密码
	router.POST("/check", Check)
}
