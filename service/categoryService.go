package service

import (
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/repository"
)

type CategoryService interface {
	Insert(cate dto.CategoryCreateDTO) entity.Category
	Update(cate dto.CategoryUpdateDTO) entity.Category
	Delete(cate entity.Category)
	SearchList(search dto.CategorySearchParam) dto.CategorySearchList
	GetTreeList(pid uint) []dto.CategoryTree
	FindById(cateID uint) entity.Category
}

type categoryService struct {
	repository.CategoryRepository
}

func (c categoryService) Insert(cate dto.CategoryCreateDTO) entity.Category {
	cateToCreate := entity.Category{}
	cateToCreate.Name =cate.Name
	cateToCreate.Remark =cate.Remark
	cateToCreate.ParentID =cate.ParentID

	res := c.CategoryRepository.Insert(cateToCreate)

	return res
}

func (c categoryService) Update(cate dto.CategoryUpdateDTO) entity.Category {
	cateToUpdate := entity.Category{}

	cateToUpdate.ID = uint(cate.ID)
	cateToUpdate.Name =cate.Name
	cateToUpdate.Remark =cate.Remark
	cateToUpdate.ParentID =cate.ParentID

	res := c.CategoryRepository.Update(cateToUpdate)
	return res
}

func (c categoryService) Delete(cate entity.Category) {
	 c.CategoryRepository.Delete(cate)
}

func (c categoryService) SearchList(search dto.CategorySearchParam) dto.CategorySearchList {
	return  c.CategoryRepository.SearchList(search)
}

func (c categoryService) GetTreeList(pid uint) []dto.CategoryTree {
	list:= c.CategoryRepository.List()

	return  getCateTree(list,pid)
}

func (c categoryService) FindById(cateID uint) entity.Category {
	return c.CategoryRepository.FindByID(cateID)
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository,
	}
}

func getCateTree(list []entity.Category, pid uint) []dto.CategoryTree {
	var CategoryTree []dto.CategoryTree
	for _,val := range list {
		if val.ParentID == pid {
			child := getCateTree(list,val.ID)
			node := dto.CategoryTree {
				ID: val.ID,
				ParentID: val.ParentID,
				Name: val.Name,
			}
			node.Children = child
			CategoryTree = append(CategoryTree,node)
		}
	}

	return  CategoryTree
}