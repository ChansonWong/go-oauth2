package model

import "server/dao"

type AuthClientDetail struct {
	dao.Model
	Appid   string `json:"appid"`
	Secret  string `json:"secret"`
	AppName string `json:"app_name"`
}

type AuthAccessToken struct {
	dao.Model
	AccessToken string `json:"access_token"`
	UserId      string `json:"user_id"`
	UserName    string `json:"user_name"`
	Appid       string `json:"appid"`
	ExpiresIn   int    `json:"expires_in"`
	GrantType   string `json:"grant_type"`
	Scope       string `json:"scope"`
}

type AuthRefreshToken struct {
	dao.Model
	AccessTokenId string `json:"access_token_id"` // AuthAccessToken表的主键
	RefreshToken  string `json:"refresh_token"`
	ExpiresIn     int    `json:"expires_in"`
}
