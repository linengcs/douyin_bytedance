package services

import (
	"douyin/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDb() {
	var err error
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.Conf.DbConfig.USER,
		setting.Conf.DbConfig.PASSWORD,
		setting.Conf.DbConfig.HOST,
		setting.Conf.DbConfig.PORT,
		setting.Conf.DbConfig.NAME)
	Db, err = gorm.Open(mysql.Open(connectStr), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
}
