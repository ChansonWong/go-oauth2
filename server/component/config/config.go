package config

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	DB struct {
		Ip       string
		Port     string
		Schema   string
		Name     string
		Password string
	}
	Cache struct {
		Ip   string
		Port string
		Db   string
	}
}{}

func init() {
	err := configor.Load(&Config, "conf/application.yml")
	if err != nil {
		panic(err)
	}
}
