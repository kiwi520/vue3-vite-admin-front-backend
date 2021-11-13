package entity

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName string `json:"role_name" gorm:"type:varchar(255);comment:'角色名'"`
	Remark string `json:"remark"  gorm:"type:varchar(1000);default:'';comment:'备注'"`
	Permission string `json:"permission" gorm:"type:varchar(2000);default:'';comment:'权限'"`
}