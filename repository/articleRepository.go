package repository

import (
	"golang_api/dto"
	"golang_api/entity"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Insert(article entity.Article) entity.Article
	Update(article entity.Article) entity.Article
	Delete(article entity.Article)
	SearchList(search dto.ArticleSearchParam) dto.ArticleSearchList
	List() []entity.Article
	FindByID(articleID uint) entity.Article
}

type articleRepository struct {
	articleConnect *gorm.DB
}

func (a articleRepository) Insert(article entity.Article) entity.Article {
	a.articleConnect.Save(&article)
	a.articleConnect.Find(&article)
	return article
}

func (a articleRepository) Update(article entity.Article) entity.Article {
	err := a.articleConnect.Model(&article).Select("*").Omit("id", "CreatedAt").Updates(map[string]interface{}{"title": article.Title, "category_id": article.CategoryID,"img_path": article.ImgPath, "recommend": article.Recommend, "content": article.Content}).Error

	if err != nil{
		println("err.Error()")
		println(err.Error())
		println("err.Error()")
	}


	a.articleConnect.Find(&article)
	return article
}

func (a articleRepository) Delete(article entity.Article) {
	a.articleConnect.Unscoped().Delete(&article)
}

func (a articleRepository) SearchList(search dto.ArticleSearchParam) (data dto.ArticleSearchList) {
	departDb:= a.articleConnect.Model(&entity.Article{})
	if search.Title != "" {
		departDb.Where("title LIKE ? ", "%"+search.Title+"%")
	}

	if search.CategoryID > 0 {
		departDb.Where("category_id = ? ", search.CategoryID)
	}

	var count int64
	departDb.Count(&count)

	articleList := []dto.Article{}
	departDb.Offset(int(search.PageIndex - 1) * int(search.PageSize)).Limit(int(search.PageSize)).Find(&articleList)

	data.Count = count
	data.List = articleList
	return data
}

func (a articleRepository) List() (data []entity.Article) {
	a.articleConnect.Find(&data)
	return data
}

func (a articleRepository) FindByID(articleID uint) (data entity.Article) {
	data.ID =articleID
	a.articleConnect.Find(&data)
	return data
}

func NewArticleRepository (db *gorm.DB) ArticleRepository {
	return &articleRepository{
		articleConnect: db,
	}
}
