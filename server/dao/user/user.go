package user

import (
	"server/dao"
	"server/model"
)

func SelectByUsername(username string) (user model.User) {
	dao.GetDB().Where("username = ?", username).First(&user)
	return
}
