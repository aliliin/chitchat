package main

import (
	"github.com/aliliin/chitchat/routes"
	"log"
	"net/http"
)

func main() {
	startWebServer("9501")
}

func startWebServer(port string) {
	r := routes.NewRouter() // 通过 router.go 中定义的路由器来分发请求

	// 处理静态资源文件
	assets := http.FileServer(http.Dir("public"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r) // 通过 router.go 中定义的路由器来分发请求

	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil) // 启动监听程序
	if err != nil {
		log.Println("An error occured starting HTTP listener at port" + port)
		log.Println("Error: " + err.Error())
	}
}
