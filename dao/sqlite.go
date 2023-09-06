package dao

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "ToDoList/models"
)

var (
    DB *gorm.DB
)

func InitSqlite() (err error) {
    DB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        return err
    }

    // 当我们执行迁移操作（db.AutoMigrate(&models.Todo{})）时，
    // GORM 会根据结构体的定义自动创建相应的数据库表，并根据标签指定的列名进行映射。
    DB.AutoMigrate(&models.Todo{})

    return nil
}