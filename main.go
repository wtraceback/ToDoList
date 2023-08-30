package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// 定义结构体 NotesFormat
type NotesFormat struct {
    ID int
    Notes string
}

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

    // 创建 NotesFormat 类型的切片
    var todolist_data []NotesFormat
    // 添加元素到切片
    todolist_data = append(todolist_data, NotesFormat{ID: 1, Notes: "Hello"})
    todolist_data = append(todolist_data, NotesFormat{ID: 2, Notes: "World"})

    r.GET("/todolist", func(c *gin.Context) {
        c.JSON(http.StatusOK, todolist_data)
    })

    // 启动服务器，默认在 0.0.0.0:8080 启动服务
    r.Run()
}