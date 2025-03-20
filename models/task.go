package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title  string `gorm:"not null"`
	Status string `gorm:"default:'pending'"`
	UserID uint   `gorm:"not null"`
}
