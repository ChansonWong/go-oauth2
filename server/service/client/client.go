package client

import (
	clientDao "server/dao/client"
	"server/model"
)

func FindAppClientByAppid(appid string) (appClient *model.AuthClientDetail) {
	appClient = clientDao.FindAppClientByAppid(appid)
	return
}

func FindByUserIdAppidScope(userId, appid, scope string) (authAccessToken *model.AuthAccessToken, err error) {
	authAccessToken, err = clientDao.FindByUserIdAppidScope(userId, appid, scope)
	return
}

func FindByAccessToken(accessToken string) (authAccessToken *model.AuthAccessToken, err error) {
	authAccessToken, err = clientDao.FindByAccessToken(accessToken)
	return
}

func FindByAccessTokenId(id string) (authRefreshToken *model.AuthRefreshToken, err error) {
	authRefreshToken, err = clientDao.FindByAccessTokenId(id)
	return
}

func CreateAccessToken(model model.AuthAccessToken) (err error) {
	err = clientDao.CreateAccessToken(model)
	return
}

func UpdateAccessToken(id string, updateModel model.AuthAccessToken) (err error) {
	err = clientDao.UpdateAccessToken(id, updateModel)
	return
}

func CreateRefreshToken(model model.AuthRefreshToken) (err error) {
	err = clientDao.CreateRefreshToken(model)
	return
}

func UpdateRefreshToken(id string, updateModel model.AuthRefreshToken) (err error) {
	err = clientDao.UpdateRefreshToken(id, updateModel)
	return
}
