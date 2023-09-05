package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "math/rand"
    "strconv"
    "time"
    "ToDoList/setting"
)

// 创建 NotesFormat 类型的切片
var todo_data []setting.NotesFormat = make([]setting.NotesFormat, 0)

func IndexHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", nil)
}

func GetTodoList(c *gin.Context) {
    c.JSON(http.StatusOK, todo_data)
}

func CreateTodo(c *gin.Context) {
    // 生成一个[0, 100)之间的随机整数
    randomNum := rand.Intn(1000)

    // 获取表单数据
    notes := c.PostForm("notes")
    todo_data = append(todo_data, setting.NotesFormat{
        ID: randomNum,
        Notes: notes,
        Time: time.Now().Format("2006-01-02 15:04:05"),
    })
    c.JSON(http.StatusOK, "操作成功")
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

    c.JSON(http.StatusOK, notes)
}
