package dto

import (
	"gorm.io/gorm"
)

type DepartmentUpdateDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
	DepartmentName string `json:"department_name" form:"department_name" binding:"required"`
	DepartmentLeaderID uint `json:"department_leader_id" form:"department_leader_id" binding:"required"`
	ParentID uint `json:"parent_id" form:"parent_id"`
}

type DepartmentCreateDTO struct {
	DepartmentName string `json:"department_name" form:"department_name" binding:"required"`
	DepartmentLeaderID uint `json:"department_leader_id" form:"department_leader_id" binding:"required"`
	ParentID uint `json:"parent_id,default=0" form:"parent_id,default=0"`
}

type DepartmentTree struct {
	ID uint `json:"id"`
	ParentID uint `json:"parent_id"`
	DepartmentName string `json:"department_name"`
	DepartmentLeaderID uint `json:"department_leader_id"`
	Children []DepartmentTree `json:"children"`
}

type DepartmentSearchParam struct {
	DepartmentName string `json:"department_name"`
	PageIndex uint `json:"page_index" binding:"required"`
	PageSize uint `json:"page_size" binding:"required"`
}


//部门列表查询response结构
type DepartmentSearchList struct {
	Count int64 `json:"total"`
	List []Department `json:"list"`
}

type Department struct {
	gorm.Model
	DepartmentName string `json:"department_name" gorm:"type:varchar(255)"`
	DepartmentLeaderID uint `json:"department_leader_id" gorm:"type:uint"`
	ParentID uint `json:"parent_id" gorm:"type:uint;default:0;not null"`
	LeaderName string `json:"leader_name" gorm:"leader_name"`
	LeaderEmail string `json:"leader_email" gorm:"leader_email"`
}