package service

import (
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/helper"
	"golang_api/repository"
)

type MenuService interface {
	Insert(menu dto.MenuCreteDTO) entity.Menu
	Update(menu dto.MenuUpdateDTO) entity.Menu
	Delete(menu entity.Menu)
	List(search dto.MenuSearchParam) dto.MenuSearchList
	GetMenuTreeList(pid uint) []dto.MenuTree
	FindById(menuID uint) entity.Menu
}

type menuService struct {
	repository.MenuRepository
}

func (m menuService) GetMenuTreeList(pid uint) []dto.MenuTree {
	list:= m.MenuRepository.MenuList()

	return  helper.GetMenuTree(list,pid)
}

func (m menuService) Insert(menu dto.MenuCreteDTO) entity.Menu {
	menuToCreate := entity.Menu{}
	menuToCreate.Name =menu.Name
	menuToCreate.Type =menu.Type
	menuToCreate.State =menu.State
	menuToCreate.Code =menu.Code
	menuToCreate.Icon =menu.Icon
	menuToCreate.Path =menu.Path
	menuToCreate.ParentID =menu.ParentID
	menuToCreate.Component =menu.Component

	res := m.MenuRepository.InsertMenu(menuToCreate)

	return res
}

func (m menuService) Update(menu dto.MenuUpdateDTO) entity.Menu {
	menuToUpdate := entity.Menu{}

	menuToUpdate.ID = uint(menu.ID)
	menuToUpdate.Name =menu.Name
	menuToUpdate.Type =menu.Type
	menuToUpdate.State =menu.State
	menuToUpdate.Code =menu.Code
	menuToUpdate.Icon =menu.Icon
	menuToUpdate.Path =menu.Path
	menuToUpdate.ParentID =menu.ParentID
	menuToUpdate.Component =menu.Component

	res := m.MenuRepository.UpdateMenu(menuToUpdate)
	return res
}

func (m menuService) Delete(menu entity.Menu) {
	m.MenuRepository.DeleteMenu(menu)
}

func (m menuService) List(search dto.MenuSearchParam) dto.MenuSearchList {
	return  m.MenuRepository.MenuSearchList(search)
}

func (m menuService) FindById(menuID uint) entity.Menu {
	panic("implement me")
}

func NewMenuService(repository repository.MenuRepository) MenuService {
	return &menuService{
		repository,
	}
}
