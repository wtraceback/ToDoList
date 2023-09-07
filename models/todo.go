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

// 获取单个的 todo 任务
func GetTodoById(id string) (todo *Todo, err error) {
    // 创建一个指向 Todo 类型的零值的指针。它会分配内存并将其初始化为类型的零值。
    todo = new(Todo)
    err = dao.DB.First(todo, id).Error
    if err != nil {
        return nil, err
    }

    return todo, nil
}

// 创建 todo
func CreateTodo(todo *Todo) (err error) {
    err = dao.DB.Create(todo).Error
    return err
}

// 删除 todo
func DeleteTodo(id string) (err error) {
    err = dao.DB.Where("id = ?", id).Delete(&Todo{}).Error
    return
}

// 更新 todo
func UpdateTodo(todo *Todo) (err error) {
    err = dao.DB.Save(todo).Error
    return
}