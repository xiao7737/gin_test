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

// 利用go test作性能分析
//	 go test -bench . -benchmem -cpuprofile cpu.prof 获取CPU性能数据 或者 go test -bench=. -cpuprofile=cpu.prof
//        go tool pprof cpu.prof
//	 go test -bench . -benchmem -memprofile mem.prof    获取内存性能数据
// 		  go tool pprof men.prof

//  pprof 命令下
// 输入:list 函数名      查看函数具体耗时
// 输入：web  查看展示图
// 输入：svg  生成svg图像
