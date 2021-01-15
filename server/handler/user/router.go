package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/common/constants"
	"server/response"
	"server/service/user"
)

func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func Check(ctx *gin.Context) {
	username := ctx.DefaultPostForm("username", "")
	password := ctx.DefaultPostForm("password", "")

	errMap := make(map[string]string)

	if username != "" && password != "" {
		// 1. 登录验证
		passwordVerify, userInfo := user.CheckLogin(username, password)

		// 登录验证通过
		if passwordVerify && userInfo.Id != "" {
			// 2. session中添加用户信息
			ctx.Set(constants.SESSION_USER, userInfo)
			// 登录成功之后的回调地址
			redirectUrl := ctx.GetString(constants.SESSION_LOGIN_REDIRECT_URL)
			ctx.Set(constants.SESSION_LOGIN_REDIRECT_URL, "")

			result := make(map[string]string)
			result["redirect_uri"] = redirectUrl
			ctx.JSON(http.StatusOK, response.GenSuccess(result))
			return
		} else {
			if !passwordVerify {
				errMap["msg"] = "用户名或密码错误！"
			}

			if userInfo.Id == "" {
				errMap["msg"] = "该用户还未开通！"
			}
		}
	}

	ctx.JSON(http.StatusOK, errMap)
	return
}
