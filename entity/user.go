package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(255)" json:"name"`
	Email string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token string `gorm:"-" json:"token,omitempty"`
	Books []Book `json:"books,omitempty"`
}
