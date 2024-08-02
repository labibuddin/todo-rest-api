package models

import (
	"gorm.io/gorm"
)

type Todo struct {
    gorm.Model
    Title       string `json:"title" gorm:"text;not null"`
    Description string `json:"description" gorm:"text;not null"`
    Completed   bool   `json:"completed" gorm:"bool;false"`
}
