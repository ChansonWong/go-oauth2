package client

import (
	clientDao "server/dao/client"
	"server/model"
)

func FindAppClientByAppid(appid string) (appClient *model.AuthClientDetail){
	appClient = clientDao.FindAppClientByAppid(appid)
	return
}
