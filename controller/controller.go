package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "ToDoList/setting"
    "ToDoList/models"
)

// 创建 NotesFormat 类型的切片
var todo_data []setting.NotesFormat = make([]setting.NotesFormat, 0)

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
        Notes: notes,
    }

    // 存入数据库
    err := models.CreateTodo(&todo)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "message": "操作失败",
            "error": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "message": "操作成功",
        })
    }
}

func DeleteTodo(c *gin.Context) {
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
}

func UpdateTodo(c *gin.Context) {
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

    c.JSON(http.StatusOK, gin.H{
        "message": "操作成功",
    })
}
