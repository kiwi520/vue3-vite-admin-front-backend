package repository

import (
	"golang_api/dto"
	"golang_api/entity"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Insert(cate entity.Category) entity.Category
	Update(cate entity.Category) entity.Category
	Delete(cate entity.Category)
	SearchList(search dto.CategorySearchParam) dto.CategorySearchList
	List() []entity.Category
	FindByID(cateID uint) entity.Category
}

type categoryRepository struct {
	cateConnect *gorm.DB
}

func (c categoryRepository) Insert(cate entity.Category) entity.Category {
	c.cateConnect.Save(&cate)
	c.cateConnect.Find(&cate)
	return cate
}

func (c categoryRepository) Update(cate entity.Category) entity.Category {
	err := c.cateConnect.Model(&cate).Select("*").Omit("id", "CreatedAt").Updates(map[string]interface{}{"name": cate.Name, "parent_id": cate.ParentID, "remark": cate.Remark}).Error

	if err != nil{
		println("err.Error()")
		println(err.Error())
		println("err.Error()")
	}


	c.cateConnect.Find(&cate)
	return cate
}

func (c categoryRepository) Delete(cate entity.Category) {
	c.cateConnect.Unscoped().Delete(&cate)
}

func (c categoryRepository) SearchList(search dto.CategorySearchParam) (data dto.CategorySearchList) {
	departDb:= c.cateConnect.Model(&entity.Category{})
	if search.Name != "" {
		departDb.Where("name LIKE ? ", "%"+search.Name+"%")
	}

	var count int64
	departDb.Count(&count)

	cateList := []dto.Category{}
	departDb.Offset(int(search.PageIndex - 1) * int(search.PageSize)).Limit(int(search.PageSize)).Find(&cateList)

	data.Count = count
	data.List = cateList
	return data
}

func (c categoryRepository) List() (data []entity.Category) {
	c.cateConnect.Find(&data)
	return data
}

func (c categoryRepository) FindByID(cateID uint) entity.Category {
	var cat entity.Category
	cat.ID =cateID
	c.cateConnect.Find(&cat)
	return cat
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		cateConnect: db,
	}
}
