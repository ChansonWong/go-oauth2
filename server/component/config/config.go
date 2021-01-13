package config

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
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
