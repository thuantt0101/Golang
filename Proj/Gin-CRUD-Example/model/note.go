package model

import "github.com/jinzhu/gorm"

type (
	Note1 struct {
		gorm.Model
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	// transformedTodo represents a formatted note1
	TranformNote1 struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)
