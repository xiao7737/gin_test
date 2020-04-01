package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var Eloquent *gorm.DB

func init() {
	var err error
	Eloquent, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gin_test?charset=utf8&parseTime=True&loc=Local&timeout=10ms")
	if err != nil {
		fmt.Printf("database connect err: %v", err)
	}
	if Eloquent.Error != nil {
		fmt.Printf("database err: %v", Eloquent.Error)
	}
}
