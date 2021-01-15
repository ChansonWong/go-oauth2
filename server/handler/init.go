package handler

import (
	"github.com/gin-gonic/gin"
	"server/handler/oauth2"
	"server/handler/user"
)

func ConfigRouter(router *gin.Engine) {
	router.LoadHTMLGlob("view/pages/*.html")
	router.Static("/view", "./view")
	oauth2.LoadOAuth2(router)
	user.LoadUser(router)
}
