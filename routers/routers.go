package routers

import (
    "github.com/gin-gonic/gin"
    "ToDoList/controller"
)

func SetupRouter() *gin.Engine {
    // 创建一个默认的 Gin 路由引擎实例
    // 默认路由引擎实例已经预先配置了一些常用的中间件，例如日志记录和错误恢复中间件
    route := gin.Default()

    // 静态文件的处理 (告诉 gin 框架模板文件引用的静态文件去哪里找)
    route.Static("/static", "./static")

    // HTML 模板渲染 (告诉 gin 框架去哪里找模板文件)
    route.LoadHTMLGlob("templates/*.html")

    // 注册首页的路由
    route.GET("/", controller.IndexHandler)

    // ToDoList 任务相关
    // 获取所有的待办事项
    route.GET("/todo", controller.GetTodoList)

    // 创建新的 todo 任务
    route.POST("/todo", controller.CreateTodo)

    // // 删除 todo 任务
    route.DELETE("/todo/:id", controller.DeleteTodo)

    // // 更新 todo 任务
    route.PUT("/todo/:id", controller.UpdateTodo)

    return route
}