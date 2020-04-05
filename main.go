package main

import (
	_ "gin_test/api/database"
	orm "gin_test/api/database"
	"gin_test/api/router"
	"net/http"
	_ "net/http/pprof" // 开启监控
)

func main() {
	go func() {
		_ = http.ListenAndServe(":8080", nil) //查看  http://localhost:8080/debug/pprof/
	}()
	defer orm.Eloquent.Close()
	Router := router.InitRouter()
	_ = Router.Run(":9999")
}

//  go tool pprof -alloc_space http://127.0.0.1:8080/debug/pprof/heap	 内存的临时分配情况，可以提高程序的运行速度
//  go tool pprof -inuse_space http://127.0.0.1:8080/debug/pprof/heap    程序常驻内存的占用情况

//火炬图
// 在项目目录生成touch.svg，内存临时分配情况
//	go-torch -alloc_space http://127.0.0.1:8080/debug/pprof/heap --colors=mem
