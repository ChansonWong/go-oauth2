package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
	"server/component/config"
	"strings"
)

var db *gorm.DB

type Model struct {
	Id string `json:"id" gorm:"primary_key"`
}

func init() {
	dbStr := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?timeout=20s&readTimeout=20s",
		config.Config.DB.Name, config.Config.DB.Password, config.Config.DB.Ip, config.Config.DB.Port,
		config.Config.DB.Schema)
	log.Info().Msgf("db url %s.", dbStr)
	var err error
	db, err = gorm.Open("mysql", dbStr)
	if err != nil {
		log.Error().Msgf("connection error %v", err)
		panic(err)
	} else {
		log.Info().Msgf("database connected")
	}

	db.Callback().Create().Before("gorm:create").Register("before_create", createCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func createCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		if idField, ok := scope.FieldByName("Id"); ok {
			if idField.IsBlank {
				id := strings.Replace(uuid.NewV1().String(), "-", "", -1)
				idField.Set(id)
			}
		}
	}
}

func GetDB() *gorm.DB {
	return db
}
