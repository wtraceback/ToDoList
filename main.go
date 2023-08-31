package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "math/rand"
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
    var todolist_data []NotesFormat = make([]NotesFormat, 0)

    r.GET("/todolist", func(c *gin.Context) {
        c.JSON(http.StatusOK, todolist_data)
    })

    r.POST("/todolist", func(c *gin.Context) {
        // 生成一个[0, 100)之间的随机整数
        randomNum := rand.Intn(1000)

        // 获取表单数据
        notes := c.PostForm("notes")
        todolist_data = append(todolist_data, NotesFormat{ID: randomNum, Notes: notes})
        c.JSON(http.StatusOK, "操作成功")
    })

    // 启动服务器，默认在 0.0.0.0:8080 启动服务
    r.Run()
}