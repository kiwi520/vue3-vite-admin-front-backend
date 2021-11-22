package dto

import "gorm.io/gorm"

type ArticleUpdateDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
	Title string `json:"title" form:"title" binding:"required"`
	CategoryID uint `json:"category_id" form:"category_id" binding:"required"`
	Recommend uint `json:"recommend" form:"recommend" binding:"required,oneof=1 0 "`
	Content string `json:"content" binding:"min=10" `
	Remark string `json:"remark" binding:"min=0,max=255" `
	ImgPath string `json:"img_path" binding:"min=0,max=255" `
}

type ArticleCreateDTO struct {
	Title string `json:"title" form:"title" binding:"required"`
	CategoryID uint `json:"category_id" form:"category_id" binding:"required"`
	Recommend uint `json:"recommend" form:"recommend" binding:"oneof=1 0 "`
	Content string `json:"content" binding:"min=10" `
	Remark string `json:"remark" binding:"min=0,max=255" `
	ImgPath string `json:"img_path" binding:"min=0,max=255" `
}


type ArticleSearchParam struct {
	Title string `json:"title"`
	CategoryID uint `json:"category_id" form:"category_id"`
	PageIndex uint `json:"page_index" binding:"required"`
	PageSize uint `json:"page_size" binding:"required"`
}

//列表查询response结构
type ArticleSearchList struct {
	Count int64 `json:"total"`
	List []Article `json:"list"`
}

type Article struct {
	gorm.Model
	Title string `json:"title" gorm:"type:varchar(255);comment:'app名称'"`
	ImgPath string `json:"img_path" gorm:"type:varchar(555);comment:'图片路径'"`
	CategoryID int64 `json:"category_id" gorm:"type:int;not null default 0;comment:'所属分类'"`
	Recommend uint `json:"recommend" gorm:"type:smallint;default:0;comment:'1: 推荐 0: 不推荐'"`
	Content string `json:"content" gorm:"type:text;;comment:'内容'"`
}