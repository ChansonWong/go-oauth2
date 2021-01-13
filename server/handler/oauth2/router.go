package oauth2

import "github.com/gin-gonic/gin"

func LoadOAuth2(router *gin.Engine){
	oauth2Grp := router.Group("/oauth2")
	{
		// 获取Authorization Code
		oauth2Grp.GET("/authorize", Authorize)
		// 通过Authorization Code获取Access Token
		oauth2Grp.GET("/code2Session", Code2Session)
	}
}
