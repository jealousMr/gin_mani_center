package dal

import (
	"fmt"
	"gin_mani_center/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbProxy *gorm.DB

func InitDB() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.DbUserName, conf.DbPassword, conf.ServerIp, conf.DbPort, conf.DbName)
	dbProxy, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}

func GetDBProxy() (*gorm.DB, error) {
	var err error
	if dbProxy == nil {
		err = InitDB()
	}
	return dbProxy, err
}
