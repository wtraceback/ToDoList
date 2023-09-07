package main

import (
    "ToDoList/routers"
    "fmt"
    "ToDoList/dao"
    "ToDoList/models"
)

func main() {
    // 连接数据库
    err := dao.InitSqlite()
    if err != nil {
        fmt.Printf("初始化数据库失败，错误信息为：%v\n", err)
        return
    }
    // 当我们执行迁移操作（db.AutoMigrate(&models.Todo{})）时，
    // GORM 会根据结构体的定义自动创建相应的数据库表，并根据标签指定的列名进行映射。
    dao.DB.AutoMigrate(&models.Todo{})


    // 注册路由
    route := routers.SetupRouter()
    // 启动服务器，默认在 0.0.0.0:8080 启动服务
    err = route.Run()
    if err != nil {
        fmt.Printf("服务器启动失败，error: %v\n", err)
    }

    // https://gitee.com/gitwbbin/go-gin-demo
}