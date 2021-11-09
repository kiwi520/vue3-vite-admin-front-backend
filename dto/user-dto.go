package dto

import "gorm.io/gorm"

type UserUpdateDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"  validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required" `
}

//type UserCreateDTO struct {
//	Name string `json:"name" form:"name" binding:"required"`
//	Email string `json:"email" form:"email" binding:"required" validate:"email"`
//	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
//}

type AdminCreateDTO struct {
	Name string ` json:"name" binding:"required"`
	Email string ` json:"email" binding:"required"`
	Mobile string ` json:"mobile" binding:"required"`
	Gender uint `json:"gender" binding:"required,oneof=1 0"`
	DepartmentID uint `json:"department_id" binding:"required"`
	RoleID uint `json:"role_id" binding:"required"`
	State uint `json:"state" binding:"required,oneof=1 2 3"`
	Job string `json:"job"  binding:"required"`
	Remark string `json:"remark" binding:"min=0,max=255" `
	Password string ` json:"password" binding:"required"`
}

type AdminUpdateDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
	Name string ` json:"name" binding:"required"`
	Email string ` json:"email" binding:"required"`
	Mobile string ` json:"mobile" binding:"required"`
	Gender uint `json:"gender" binding:"required,oneof=1 0"`
	DepartmentID uint `json:"department_id" binding:"required"`
	RoleID uint `json:"role_id" binding:"required"`
	State uint `json:"state" binding:"required"`
	Job string `json:"job"  binding:"required"`
	Remark string `json:"remark" binding:"min=0,max=255" `
}



type AdminSearchParam struct {
	Name string `json:"name"`
	State uint `json:"state"`
	PageIndex uint `json:"page_index" binding:"required"`
	PageSize uint `json:"page_size" binding:"required"`
}


//角色列表查询response结构
type AdminSearchList struct {
	Count int64 `json:"total"`
	List []Admin `json:"list"`
}

type Admin struct {
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
}