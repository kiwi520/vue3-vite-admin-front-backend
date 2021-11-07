package entity

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title string `json:"title" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:text"`
	UserId uint64 `json:"-" gorm:"not null"`
}
