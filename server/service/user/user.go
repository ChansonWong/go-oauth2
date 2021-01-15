package user

import (
	"server/dao/user"
	"server/model"
	"server/utils"
)

func CheckLogin(username, password string) (result bool, model model.User) {
	user := user.SelectByUsername(username)
	result, _ = utils.PasswordVerify(password, user.Password)
	model = user
	return
}
