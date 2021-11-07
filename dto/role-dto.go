package dto

import "gorm.io/gorm"

type RoleInsertDTO struct {
	RoleName string `json:"role_name" form:"role_name" binding:"required"`
	Remark   string `json:"remark"  form:"remark"  binding:"required"`
}

type RoleUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	RoleName string `json:"role_name" form:"role_name" binding:"required"`
	Remark   string `json:"remark" form:"remark" binding:"required"`
}

type RoleSearchParam struct {
	RoleName string `json:"role_name"`
	PageIndex uint `json:"page_index" binding:"required"`
	PageSize uint `json:"page_size" binding:"required"`
}


//角色列表查询response结构
type RoleSearchList struct {
	Count int64 `json:"total"`
	List []Role `json:"list"`
}

type Role struct {
	gorm.Model
	RoleName string `json:"role_name" gorm:"type:varchar(255);comment:'角色名'"`
	Remark string `json:"remark"  gorm:"type:varchar(1000);comment:'备注'"`
}