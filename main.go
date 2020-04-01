package main

import (
	_ "gin_test/api/database"
	orm "gin_test/api/database"
	"gin_test/api/router"
)

func main() {
	defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":9999")
}
