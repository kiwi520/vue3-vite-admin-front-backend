package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(255)"`
	ParentID uint `json:"parent_id" gorm:"type:uint;default:0;not null"`
	Remark string `json:"remark" gorm:"type:varchar(255)" `
}
