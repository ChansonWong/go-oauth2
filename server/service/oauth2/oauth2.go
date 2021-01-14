package oauth2

import (
	"github.com/jinzhu/gorm"
	"server/component/cache"
	"server/enum"
	"server/model"
	"server/service/client"
	"server/utils"
)

// @title CreateAuthorizationCode
// @description 根据clientId、scope以及当前时间戳生成AuthorizationCode（有效期为10分钟）
// @auth kwokchiu
// @param appId string 应用id
// @param scope string 权限范围
// @param user interface{} 用户信息
// @return encryptedStr AuthorizationCode
func CreateAuthorizationCode(appId, scope string, user interface{}) (encryptedStr string) {
	// 1. 拼装待加密字符串（clientId + scope + 当前精确到毫秒的时间戳）
	str := appId + scope + utils.Int64ToString(utils.CurrentTimeMillis())

	// 2. SHA1加密
	encryptedStr = utils.Sha1(str)

	// 3.1 保存本次请求的授权范围
	cache.GetConn().SetEx(encryptedStr+":scope", utils.MinuteDuration(enum.AUTHORIZATION_CODE_EXPIRE), scope)
	// 3.2 保存本次请求所属的用户信息
	cache.GetConn().SetEx(encryptedStr+":user", utils.MinuteDuration(enum.AUTHORIZATION_CODE_EXPIRE), user)

	// 4. 返回Authorization Code
	return
}

func CreateAccessToken(user *model.User, appClient *model.AuthClientDetail, grantType, scope string, expiresIn int) (accessTokenStr string) {
	expiresAt := utils.NextDaysSecond(enum.ACCESS_TOKEN_EXPIRE)

	// 1. 拼装待加密字符串（username + appId + 当前精确到毫秒的时间戳）
	str := user.Username + appClient.Appid + utils.Int64ToString(utils.CurrentTimeMillis())

	// 2. SHA1加密
	accessTokenStr = "1." + utils.Sha1(str) + "." + utils.IntToString(expiresIn) + "." + utils.IntToString(expiresAt)

	// 3. 保存Access Token
	authAccessToken, err := client.FindByUserIdAppidScope(user.Id, appClient.Appid, scope)
	// 如果存在userId + clientId + scope匹配的记录，则更新原记录，否则向数据库中插入新记录
	if err == gorm.ErrRecordNotFound {
		createAccessToken := model.AuthAccessToken{
			AccessToken: accessTokenStr,
			UserId:      user.Id,
			UserName:    user.Username,
			Appid:       appClient.Appid,
			ExpiresIn:   expiresIn,
			GrantType:   grantType,
			Scope:       scope,
		}
		client.CreateAccessToken(createAccessToken)
	} else {
		updateAccessToken := model.AuthAccessToken{
			AccessToken: accessTokenStr,
			ExpiresIn:   expiresIn,
		}
		client.UpdateAccessToken(authAccessToken.Id, updateAccessToken)
	}

	// 4. 返回Access Token
	return
}

func CreateRefreshToken(user *model.User, authAccessToken *model.AuthAccessToken) (refreshTokenStr string) {
	// 过期时间
	expiresIn := utils.DayToSecond(enum.REFRESH_TOKEN_EXPIRE)
	// 过期的时间戳
	expiresAt := utils.NextDaysSecond(enum.REFRESH_TOKEN_EXPIRE)

	// 1. 拼装待加密字符串（username + accessToken + 当前精确到毫秒的时间戳）
	str := user.Username + authAccessToken.AccessToken + utils.Int64ToString(utils.CurrentTimeMillis())

	// 2. SHA1加密
	refreshTokenStr = "2." + utils.Sha1(str) + "." + utils.IntToString(expiresIn) + "." + utils.IntToString(expiresAt)

	// 3. 保存Refresh Token
	authRefreshToken, err := client.FindByAccessTokenId(authAccessToken.Id)
	if err == gorm.ErrRecordNotFound {
		createRefreshToken := model.AuthRefreshToken{
			AccessTokenId: authAccessToken.Id,
			RefreshToken:  refreshTokenStr,
			ExpiresIn:     expiresIn,
		}
		client.CreateRefreshToken(createRefreshToken)
	} else {
		updateRefreshToken := model.AuthRefreshToken{
			RefreshToken: refreshTokenStr,
			ExpiresIn:    expiresIn,
		}
		client.UpdateRefreshToken(authRefreshToken.Id, updateRefreshToken)
	}

	// 4. 返回Refresh Token
	return
}
