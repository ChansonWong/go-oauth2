package model

type AuthClientDetail struct {
	Id string `json:"id"`
	Appid string `json:"appid"`
	Secret string `json:"secret"`
	AppName string `json:"app_name"`
}
