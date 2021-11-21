package dto

import (
"gorm.io/gorm"
)

type CategoryUpdateDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
	ParentID uint `json:"parent_id" form:"parent_id"`
	Remark string `json:"remark" binding:"min=0,max=255" `
}

type CategoryCreateDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
	ParentID uint `json:"parent_id,default=0" form:"parent_id,default=0"`
	Remark string `json:"remark" binding:"min=0,max=255" `
}

type CategoryTree struct {
	ID uint `json:"id"`
	ParentID uint `json:"parent_id"`
	Name string `json:"name"`
	Children []CategoryTree `json:"children"`
}

type CategorySearchParam struct {
	Name string `json:"name"`
	PageIndex uint `json:"page_index" binding:"required"`
	PageSize uint `json:"page_size" binding:"required"`
}


//部门列表查询response结构
type CategorySearchList struct {
	Count int64 `json:"total"`
	List []Category `json:"list"`
}

type Category struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(255)"`
	ParentID uint `json:"parent_id" gorm:"type:uint;default:0;not null"`
	Remark string `json:"remark" binding:"min=0,max=255" `
}
