package models

import (
    "time"
    "ToDoList/dao"
)

// Todo Model
type Todo struct {
    ID    int    `gorm:"column:id;primaryKey;autoIncrement"`
    Notes string `gorm:"column:notes;not null"`
    CreatedAt time.Time `gorm:"column:created_at"`
    UpdatedAt time.Time `gorm:"column: update_at"`
}


/*
* Todo 这个 Model 的增删查改操作都放在这里
*/
// 获取所有的 todo 任务
func GetAllTodo() (todolist []*Todo, err error) {
    err = dao.DB.Find(&todolist).Error
    if err != nil {
        return nil, err
    }

    return todolist, err
}

// 创建 todo
func CreateTodo(todo *Todo) (err error) {
    err = dao.DB.Create(&todo).Error
    return err
}