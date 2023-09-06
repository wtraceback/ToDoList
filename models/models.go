package models

import (
    "time"
)

type Todo struct {
    ID    int    `gorm:"column:id;primaryKey;autoIncrement"`
    Notes string `gorm:"column:notes;not null"`
    CreatedAt time.Time `gorm:"column:created_at"`
    UpdatedAt time.Time `gorm:"column: update_at"`
}