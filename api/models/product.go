package models

import (
	orm "gin_test/api/database"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Code  string `json:"code"`
	Price uint   `json:"price"`
}

func init() {
	orm.Eloquent.AutoMigrate(&Product{})
}
