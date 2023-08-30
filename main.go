package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // 创建一个默认的 Gin 路由引擎实例
    // 默认路由引擎实例已经预先配置了一些常用的中间件，例如日志记录和错误恢复中间件
    r := gin.Default()

    // 静态文件的处理
    r.Static("/static", "./static")

    // HTML 模板渲染
    r.LoadHTMLGlob("templates/*.html")

    // 注册首页的路由
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{})
    })

    // 启动服务器，默认在 0.0.0.0:8080 启动服务
    r.Run()
}