package main

import (
	"fmt"
	"task4/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的Gin引擎
	r := gin.Default()
	// 注册路由分组
	routers.SetUserApi(r)
	routers.SetPostApi(r)
	routers.SetCommentApi(r)
	//加载前端模板
	// r.LoadHTMLGlob("templates/*") // 确保模板文件路径正确

	// r.Use(middlewares.JWTAuthMiddleware()) // 注册全局中间件，每个请求都会执行,用于登录鉴权及打印日志

	// r.Static("/static", "./static") // 静态文件目录映射

	r.Run(":8080") // 启动服务，监听8080端口
	fmt.Println("Server running on port 8080")
}
