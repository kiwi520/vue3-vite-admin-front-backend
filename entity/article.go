package entity

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title string `json:"title" gorm:"type:varchar(255);comment:'app名称'"`
	ImgPath string `json:"img_path" gorm:"type:varchar(555);comment:'图片路径'"`
	CategoryID int64 `json:"category_id" gorm:"type:int;not null default 0;comment:'所属分类'"`
	Recommend uint `json:"recommend" gorm:"type:smallint;default:0;comment:'1: 推荐 0: 不推荐'"`
	Content string `json:"content" gorm:"type:text;;comment:'内容'"`
}
