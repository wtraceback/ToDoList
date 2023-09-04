package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "math/rand"
    "strconv"
    "time"
)

// 定义结构体 NotesFormat
type NotesFormat struct {
    ID int
    Notes string
    Time string
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
    var todo_data []NotesFormat = make([]NotesFormat, 0)

    // 获取所有的待办事项
    r.GET("/todo", func(c *gin.Context) {
        c.JSON(http.StatusOK, todo_data)
    })

    // 添加待办事项
    r.POST("/todo", func(c *gin.Context) {
        // 生成一个[0, 100)之间的随机整数
        randomNum := rand.Intn(1000)

        // 获取表单数据
        notes := c.PostForm("notes")
        todo_data = append(todo_data, NotesFormat{
            ID: randomNum,
            Notes: notes,
            Time: time.Now().Format("2006-01-02 15:04:05"),
        })
        c.JSON(http.StatusOK, "操作成功")
    })

    r.DELETE("/todo/:id", func(c *gin.Context) {
        id := c.Param("id")   // 获取路由参数 id，类型为 string
        // 将 id 转换为 int 类型
        targetID, err := strconv.Atoi(id)
        if err != nil {
            // 处理转换错误
            c.JSON(400, gin.H{"error": "无效的参数"})
            return
        }

        // 遍历切片
        for i, item := range todo_data {
            // 如果找到目标 ID
            if item.ID == targetID {
                // 从切片中删除该元素
                todo_data = append(todo_data[:i], todo_data[i+1:]...)
                break
            }
        }

        c.JSON(http.StatusOK, targetID)
    })

    r.PUT("/todo/:id", func(c *gin.Context) {
        id := c.Param("id")   // 获取路由参数 id，类型为 string
        // 将 id 转换为 int 类型
        targetID, err := strconv.Atoi(id)
        if err != nil {
            // 处理转换错误
            c.JSON(400, gin.H{"error": "无效的参数"})
            return
        }

        // 获取表单数据
        notes := c.PostForm("notes")

        // 遍历切片
        for i, item := range todo_data {
            // 如果找到目标 ID
            if item.ID == targetID {
                // 从切片中删除该元素
                todo_data[i].Notes = notes
                break
            }
        }

        c.JSON(http.StatusOK, notes)
    })


    // 启动服务器，默认在 0.0.0.0:8080 启动服务
    r.Run()
}