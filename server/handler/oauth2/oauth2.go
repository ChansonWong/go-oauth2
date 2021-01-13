package oauth2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/common/constants"
	"server/component/cache"
	"server/enum"
	"server/response"
	"server/service/client"
	"server/service/oauth2"
)

func Authorize(ctx *gin.Context){
	// 在Session中获取用户信息
	user, _ := ctx.Get(constants.SESSION_USER)
	// 客户端ID
	appId := ctx.DefaultQuery("appid", "")
	// 权限范围
	scope := ctx.DefaultQuery("scope", "basic")
	// 回调URL (如果前后端分离不需要)
	redirectUri := ctx.DefaultQuery("redirect_uri", "")

	// 生成Authorization Code
	authorizationCode := oauth2.CreateAuthorizationCode(appId, scope, user)

	params := "?code=" + authorizationCode

	// 重定向页面
	ctx.Redirect(http.StatusMovedPermanently, redirectUri + params)
}

func Code2Session(ctx *gin.Context){
	// 授权方式
	grantType := ctx.DefaultQuery("grant_type", "authorization_code")
	// 客户端ID
	appid := ctx.DefaultQuery("appid", "")
	// //接入的客户端的密钥
	secret := ctx.DefaultQuery("secret", "")
	// //前面获取的Authorization Code
	code := ctx.DefaultQuery("code", "")
	// 回调URL (如果前后端分离不需要)
	redirectUri := ctx.DefaultQuery("redirect_uri", "")

	if grantType != enum.AUTHORIZATION_CODE {
		ctx.JSON(http.StatusOK, response.GenError(response.UNSUPPORTED_GRANT_TYPE, ""))
		return
	}

	appClient := client.FindAppClientByAppid(appid)
	if !(appClient != nil && appClient.Secret == secret) {
		ctx.JSON(http.StatusOK, response.GenError(response.INVALID_CLIENT, ""))
		return
	}

	scope, _ := cache.GetConn().Get(code + ":scope")
	user, _ := cache.GetConn().Get(code + ":user")

	fmt.Println(scope)
	fmt.Println(user)
	fmt.Println(redirectUri)
	return
}
