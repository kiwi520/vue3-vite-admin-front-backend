package entity

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Type uint `json:"type" gorm:"type:smallint;default:1;comment:'性 别 1:菜单 2: 按钮'"`
	Name string `json:"name" gorm:"type:varchar(255);comment:'菜单名称'"`
	Code string `json:"code" gorm:"type:varchar(255);comment:'权限标识'"`
	Path string `json:"path" gorm:"type:varchar(255);comment:'路由地址'"`
	Icon string `json:"icon" gorm:"type:varchar(255);comment:'路由地址'"`
	Component string `json:"component" gorm:"type:varchar(255);comment:'组件地址'"`
	State uint `json:"state" gorm:"type:smallint;default:1;comment:'菜单状态1: 正常状态 2:停用'"`
	ParentID uint `json:"parent_id" gorm:"type:uint;default:0;not null;comment:'父级ID"`
}
