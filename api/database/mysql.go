package database

import (
	"fmt"
	"gin_test/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Eloquent *gorm.DB

func init() {
	var err error
	config := conf.LoadConfig()
	User := config.Db.User
	Password := config.Db.Password
	Address := config.Db.Address
	Driver := config.Db.Driver
	Database := config.Db.Database

	Eloquent, err = gorm.Open(
		Driver, User+":"+Password+"@tcp("+Address+")/"+Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("database connect err: %v", err)
	}
	if Eloquent.Error != nil {
		fmt.Printf("database err: %v", Eloquent.Error)
	}
}

//  "root:123456@tcp(127.0.0.1:3306)/gin_test?charset=utf8&parseTime=True&loc=Local")
