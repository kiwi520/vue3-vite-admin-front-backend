package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(255)" json:"name"`
	Email string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Mobile string `gorm:"type:varchar(255)" json:"mobile" `
	Gender uint `json:"gender" gorm:"type:smallint;default:0;comment:'性 别 0:男 1: 女'"`
	DepartmentID uint `json:"department_id"`
	RoleID uint `json:"role_id"`
	State uint `json:"state" gorm:"type:smallint;default:1;comment:'1: 在职 2: 离职 3: 试用期'"`
	Job string `json:"job" gorm:"type:varchar(255)" `
	Remark string `json:"remark" gorm:"type:varchar(255)" `
	Password string `gorm:"->;<-;not null" json:"-"`
	Token string `gorm:"-" json:"token,omitempty"`
	Books []Book `json:"books,omitempty"`
}
