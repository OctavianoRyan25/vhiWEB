package config

import (
	"github.com/OctavianoRyan25/VhiWEB/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/vhiweb?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func AutoMigrate() {
	err = DB.AutoMigrate(&model.User{}, &model.Vendor{}, &model.Catalog{})
	if err != nil {
		panic("Failed to auto migrate: " + err.Error())
	}
}
