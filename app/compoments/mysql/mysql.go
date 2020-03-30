package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetDb() *gorm.DB {

	// 配置
	MyUser := "root"
	Password := "123456"
	Host := "127.0.0.1"
	Port := 3306
	Db := "yii2"

	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MyUser, Password, Host, Port, Db)
	fmt.Println(connArgs)
	db, err := gorm.Open("mysql", connArgs)

	if err != nil {
		panic(err)
	}

	return db
}
