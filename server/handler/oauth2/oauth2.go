package oauth2

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/common/constants"
	"server/component/cache"
	"server/enum"
	"server/model"
	"server/response"
	"server/service/client"
	"server/service/oauth2"
	"server/utils"
)

func Authorize(ctx *gin.Context) {
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
	ctx.Redirect(http.StatusMovedPermanently, redirectUri+params)
}

func Code2Session(ctx *gin.Context) {
	// 授权方式
	grantType := ctx.DefaultQuery("grant_type", "authorization_code")
	// 客户端ID
	appid := ctx.DefaultQuery("appid", "")
	// //接入的客户端的密钥
	secret := ctx.DefaultQuery("secret", "")
	// //前面获取的Authorization Code
	code := ctx.DefaultQuery("code", "")

	if grantType != enum.AUTHORIZATION_CODE {
		ctx.JSON(http.StatusOK, response.GenError(response.UNSUPPORTED_GRANT_TYPE, ""))
		return
	}

	appClient := client.FindAppClientByAppid(appid)
	if !(appClient != nil && appClient.Secret == secret) {
		ctx.JSON(http.StatusOK, response.GenError(response.INVALID_CLIENT, ""))
		return
	}

	scopeByte, _ := cache.GetConn().Get(code + ":scope")
	userByte, _ := cache.GetConn().Get(code + ":user")

	if utils.ByteIsNotEmpty(scopeByte) && utils.ByteIsNotEmpty(userByte) {
		var user *model.User
		json.Unmarshal(userByte, &user)
		// 过期时间
		expiresIn := utils.DayToSecond(enum.ACCESS_TOKEN_EXPIRE)
		// 生成Access Token
		accessToken := oauth2.CreateAccessToken(user, appClient, grantType, string(scopeByte), expiresIn)
		// 查询已经插入到数据库的Access Token
		authAccessToken, _ := client.FindByAccessToken(accessToken)
		// 生成Refresh Token
		refreshToken := oauth2.CreateRefreshToken(user, authAccessToken)

		resp := response.AccessTokenResp{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresIn:    expiresIn,
			Scope:        string(scopeByte),
		}

		ctx.JSON(http.StatusOK, response.GenSuccess(resp))
		return
	} else {
		ctx.JSON(http.StatusOK, response.GenError(response.INVALID_GRANT, ""))
		return
	}
}
