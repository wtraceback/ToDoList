package main

import (
    "ToDoList/routers"
    "fmt"
    "ToDoList/dao"
)

func main() {
    // 连接数据库
    err := dao.InitSqlite()
    if err != nil {
        fmt.Printf("初始化数据库失败，错误信息为：%v\n", err)
        return
    }


    // 注册路由
    route := routers.SetupRouter()

    // 启动服务器，默认在 0.0.0.0:8080 启动服务
    route.Run()
}