package oauth2

import (
	"server/component/cache"
	"server/enum"
	"server/utils"
)

// @title CreateAuthorizationCode
// @description 根据clientId、scope以及当前时间戳生成AuthorizationCode（有效期为10分钟）
// @auth kwokchiu
// @param appId string 应用id
// @param scope string 权限范围
// @param user interface{} 用户信息
// @return encryptedStr AuthorizationCode
func CreateAuthorizationCode(appId, scope string, user interface{}) (encryptedStr string){
	// 1. 拼装待加密字符串（clientId + scope + 当前精确到毫秒的时间戳）
	str := appId + scope + utils.Int64ToString(utils.CurrentTimeMillis())

	// 2. SHA1加密
	encryptedStr = utils.Sha1(str)

	// 3.1 保存本次请求的授权范围
	cache.GetConn().SetEx(encryptedStr + ":scope", enum.AUTHORIZATION_CODE_EXPIRE, scope)
	// 3.2 保存本次请求所属的用户信息
	cache.GetConn().SetEx(encryptedStr + ":user", enum.AUTHORIZATION_CODE_EXPIRE, user)

	// 4. 返回Authorization Code
	return
}
