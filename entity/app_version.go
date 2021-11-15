package entity

import "gorm.io/gorm"

type AppVersion struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(255);comment:'app名称'"`
	FilePath string `json:"file_path" gorm:"type:varchar(555);comment:'存储路径'"`
	Version  float64 `json:"version" gorm:"type:double;not null default 0;comment:'版本号'"`
	Platform uint `json:"platform" gorm:"type:smallint;default:1;comment:'1: android 2: ios'"`
	State uint `json:"state" gorm:"type:smallint;default:1;comment:'1: 未发布 2: 已发布 3: 停用'"`
	Remark string `json:"remark" gorm:"type:varchar(255)" `
}
