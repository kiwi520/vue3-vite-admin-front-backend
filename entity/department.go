package entity

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	DepartmentName string `json:"department_name" gorm:"type:varchar(255)"`
	DepartmentLeaderID uint `json:"department_leader_id" gorm:"type:uint"`
	ParentID uint `json:"parent_id" gorm:"type:uint;default:0;not null"`
}
