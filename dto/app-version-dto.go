package dto

import "gorm.io/gorm"

type AppVersionCreateDTO struct {
	Name string ` json:"name" binding:"required"`
	FilePath string `json:"file_path" binding:"required"`
	Version  float64 `json:"version" binding:"required"`
	Platform uint `json:"platform" binding:"required,oneof=1 2"`
	State uint `json:"state" binding:"required,oneof=1 2 3"`
	Remark string `json:"remark" binding:"min=0,max=255" `
}

type AppVersionUpdateDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
	Name string ` json:"name" binding:"required"`
	FilePath string `json:"file_path" binding:"required"`
	Version  float64 `json:"version" binding:"required"`
	Platform uint `json:"platform" binding:"required,oneof=1 2"`
	State uint `json:"state" binding:"required,oneof=1 2 3"`
	Remark string `json:"remark" binding:"min=0,max=255" `
}

type AppVersionMergeChunk struct {
	FileName string `json:"file_name"  binding:"required"`
	FileHash string `json:"file_hash"  binding:"required"`
}


type AppVersionSearchParam struct {
	Name string `json:"name"`
	FilePath string `json:"file_path"`
	State uint `json:"state"`
	Platform uint `json:"platform"`
	PageIndex uint `json:"page_index" binding:"required"`
	PageSize uint `json:"page_size" binding:"required"`
}


//角色列表查询response结构
type AppVersionSearchList struct {
	Count int64 `json:"total"`
	List []AppVersion `json:"list"`
}

type AppVersion struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(255);comment:'app名称'"`
	Version  float64 `json:"version" gorm:"type:double;not null default 0;comment:'版本号'"`
	Platform uint `json:"platform" gorm:"type:smallint;default:1;comment:'1: android 2: ios'"`
	State uint `json:"state" gorm:"type:smallint;default:1;comment:'1: 未发布 2: 已发布 3: 停用'"`
	Remark string `json:"remark" gorm:"type:varchar(255)" `
}