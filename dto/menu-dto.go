package dto

import "gorm.io/gorm"

type MenuCreteDTO struct {
	Type uint `json:"type" binding:"required"`
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"min=0,max=155"`
	Path string `json:"path" binding:"min=0,max=255"`
	Icon string `json:"icon" binding:"min=0,max=255"`
	Component string `json:"component" binding:"min=0,max=255"`
	State uint `json:"state" binding:"required,oneof=1 2 "`
	ParentID uint `json:"parent_id" binding:"gte=0"`
}

type MenuUpdateDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
	Type uint `json:"type" binding:"required"`
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"min=0,max=155"`
	Path string `json:"path" binding:"min=0,max=255"`
	Icon string `json:"icon" binding:"min=0,max=255"`
	Component string `json:"component" binding:"min=0,max=255"`
	State uint `json:"state" binding:"required,oneof=1 2 "`
	ParentID uint `json:"parent_id" binding:"gte=0"`
}

type MenuTree struct {
	ID uint `json:"id"`
	ParentID uint `json:"parent_id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Path string `json:"path"`
	Icon string `json:"icon"`
	Component string `json:"component"`
	Children []MenuTree `json:"children"`
}


type MenuSearchParam struct {
	Name string `json:"name"`
	State uint `json:"state"`
	PageIndex uint `json:"page_index" binding:"required"`
	PageSize uint `json:"page_size" binding:"required"`
}


//角色列表查询response结构
type MenuSearchList struct {
	Count int64 `json:"total"`
	List []Menu `json:"list"`
}

type Menu struct {
	gorm.Model
	Type uint `json:"type" gorm:"type:smallint;default:1;comment:'性 别 1:菜单 2: 按钮'"`
	Name string `json:"name" gorm:"type:varchar(255);comment:'菜单名称'"`
	Code string `json:"code" gorm:"type:varchar(255);comment:'权限标识'"`
	Icon string `json:"icon" gorm:"type:varchar(255);comment:'路由地址'"`
	Path string `json:"path" gorm:"type:varchar(255);comment:'路由地址'"`
	Component string `json:"component" gorm:"type:varchar(255);comment:'组件地址'"`
	State uint `json:"state" gorm:"type:smallint;default:1;comment:'菜单状态1: 正常状态 2:停用'"`
	ParentID uint `json:"parent_id" gorm:"type:uint;default:0;not null;comment:'父级ID"`
}