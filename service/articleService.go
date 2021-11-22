package service

import (
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/repository"
)

type ArticleService interface {
	Insert(cate dto.ArticleCreateDTO) entity.Article
	Update(cate dto.ArticleUpdateDTO) entity.Article
	Delete(cate entity.Article)
	SearchList(search dto.ArticleSearchParam) dto.ArticleSearchList
	//GetTreeList(pid uint) []dto.CategoryTree
	FindById(cateID uint) entity.Article
}

type articleService struct {
	articleRepository repository.ArticleRepository
}

func (a articleService) Insert(article dto.ArticleCreateDTO) entity.Article {
	articleToCreate := entity.Article{}
	articleToCreate.Title =article.Title
	articleToCreate.ImgPath =article.ImgPath
	articleToCreate.CategoryID = int64(article.CategoryID)
	articleToCreate.Recommend = article.Recommend
	articleToCreate.Content = article.Content

	res := a.articleRepository.Insert(articleToCreate)

	return res
}

func (a articleService) Update(article dto.ArticleUpdateDTO) entity.Article {
	artileToUpdate := entity.Article{}

	artileToUpdate.ID = uint(article.ID)
	artileToUpdate.Title =article.Title
	artileToUpdate.ImgPath =article.ImgPath
	artileToUpdate.CategoryID = int64(article.CategoryID)
	artileToUpdate.Recommend = article.Recommend
	artileToUpdate.Content = article.Content

	res := a.articleRepository.Update(artileToUpdate)
	return res
}

func (a articleService) Delete(article entity.Article) {
	a.articleRepository.Delete(article)
}

func (a articleService) SearchList(search dto.ArticleSearchParam) dto.ArticleSearchList {
	return  a.articleRepository.SearchList(search)
}

func (a articleService) FindById(articleID uint) entity.Article {
	return a.articleRepository.FindByID(articleID)
}

func NewArticleService(r repository.ArticleRepository) ArticleService {
	return &articleService{
		articleRepository: r,
	}
}
