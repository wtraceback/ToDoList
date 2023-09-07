package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "ToDoList/models"
    "time"
    "html"
)


func IndexHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", nil)
}

func GetTodoList(c *gin.Context) {
    todolist, err := models.GetAllTodo()
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "message": "操作失败",
            "error": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, todolist)
    }
}

func CreateTodo(c *gin.Context) {
    // 获取表单数据
    notes := c.PostForm("notes")

    todo := models.Todo{
        Notes: html.EscapeString(notes),
    }

    // 存入数据库
    err := models.CreateTodo(&todo)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "message": "操作失败",
            "error": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{"message": "操作成功",})
    }
}

func DeleteTodo(c *gin.Context) {
    // 获取路由参数 id，类型为 string
    id, ok := c.Params.Get("id")
    if !ok {
        c.JSON(http.StatusOK, gin.H{"error": "无效的 id"})
        return
    }

    if err := models.DeleteTodo(id); err != nil {
        c.JSON(http.StatusOK, gin.H{"error": err.Error()})
    } else {
        c.JSON(http.StatusOK, gin.H{"message": "操作成功"})
    }
}

func UpdateTodo(c *gin.Context) {
    // 获取路由参数 id，类型为 string
    id, ok := c.Params.Get("id")
    if !ok {
        c.JSON(http.StatusOK, gin.H{"error": "无效的 id"})
        return
    }

    // 获取数据库中对应 todo 的值
    todo, err := models.GetTodoById(id)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"error": err.Error()})
        return
    }

    // 获取表单数据
    notes := c.PostForm("notes")
    todo.Notes = notes
    todo.UpdatedAt = time.Now()

    err = models.UpdateTodo(todo)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"error": err.Error()})
    } else {
        c.JSON(http.StatusOK, gin.H{"message": "操作成功",})
    }
}
