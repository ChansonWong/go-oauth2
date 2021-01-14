package client

import (
	"server/dao"
	"server/model"
)

func FindAppClientByAppid(appid string) (appClient *model.AuthClientDetail) {
	dao.GetDB().Where("appid = ?", appid).First(&appClient)
	return
}

func FindByUserIdAppidScope(userId, appid, scope string) (authAccessToken *model.AuthAccessToken, err error) {
	dao.GetDB().First("user_id = ? and appid = ? and scope = ?", userId, appid, scope).First(&authAccessToken)
	return
}

func FindByAccessToken(accessToken string) (authAccessToken *model.AuthAccessToken, err error) {
	dao.GetDB().First("access_token = ?", accessToken).First(&authAccessToken)
	return
}

func FindByAccessTokenId(id string) (authRefreshToken *model.AuthRefreshToken, err error) {
	dao.GetDB().Where("access_token_id = ?", id).First(&authRefreshToken)
	return
}

func CreateAccessToken(model model.AuthAccessToken) (err error) {
	err = dao.GetDB().Create(&model).Error
	return
}

func UpdateAccessToken(id string, updateModel model.AuthAccessToken) (err error) {
	err = dao.GetDB().Where("id = ?", id).Updates(&updateModel).Error
	return
}

func CreateRefreshToken(model model.AuthRefreshToken) (err error) {
	err = dao.GetDB().Create(&model).Error
	return
}

func UpdateRefreshToken(id string, updateModel model.AuthRefreshToken) (err error) {
	err = dao.GetDB().Where("id = ?", id).Updates(&updateModel).Error
	return
}
