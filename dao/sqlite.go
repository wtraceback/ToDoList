package dao

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "fmt"
)

var (
    DB *gorm.DB
)

func InitSqlite() (err error) {
    DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        return err
    }
    fmt.Println("已连接到数据库：", DB)

    return nil
}