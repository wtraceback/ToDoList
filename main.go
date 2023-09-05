package main

import (
    "ToDoList/routers"
)

func main() {
    route := routers.SetupRouter()

    // 启动服务器，默认在 0.0.0.0:8080 启动服务
    route.Run()
}