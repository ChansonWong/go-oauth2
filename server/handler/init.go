package handler

import (
	"github.com/gin-gonic/gin"
	"server/handler/oauth2"
)

func ConfigRouter(router *gin.Engine) {
	oauth2.LoadOAuth2(router)
}
