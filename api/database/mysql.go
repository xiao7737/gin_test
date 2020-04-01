package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Eloquent *gorm.DB

func init() {
	var err error
	Eloquent, err = gorm.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/gin_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("database connect err: %v", err)
	}
	if Eloquent.Error != nil {
		fmt.Printf("database err: %v", Eloquent.Error)
	}
}
