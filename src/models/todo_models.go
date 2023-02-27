package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title string `gorm:"type:varchar(100)" json:"title"`
}
