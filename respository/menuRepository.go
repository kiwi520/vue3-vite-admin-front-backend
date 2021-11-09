package respository

import (
	"golang_api/dto"
	"golang_api/entity"
	"gorm.io/gorm"
)

type MenuRepository interface {
	InsertMenu(menu entity.Menu) entity.Menu
	UpdateMenu(menu entity.Menu) entity.Menu
	DeleteMenu(menu entity.Menu)
	MenuSearchList(search dto.MenuSearchParam) dto.MenuSearchList
	MenuList() []entity.Menu
	FindMenuByID(menuID uint) entity.Menu
}

type menuRepository struct {
	menuConnect *gorm.DB
}

func (m menuRepository) MenuList() []entity.Menu {
	var menuList = []entity.Menu{}
	m.menuConnect.Find(&menuList)
	return menuList
}

func (m menuRepository) InsertMenu(menu entity.Menu) entity.Menu {
	m.menuConnect.Save(&menu)
	m.menuConnect.Find(&menu)
	return menu
}

func (m menuRepository) UpdateMenu(menu entity.Menu) entity.Menu {
	println("menu")
	println(menu.ID)
	println(menu.Type)
	println(menu.Name)
	println("role")
	err := m.menuConnect.Model(&menu).Select("*").Omit("id", "CreatedAt").Updates(map[string]interface{}{"type": menu.Type, "name": menu.Name, "code": menu.Code, "path": menu.Path, "component": menu.Component, "state": menu.State, "parent_id": menu.ParentID}).Error

	if err != nil{
		println("err.Error()")
		println(err.Error())
		println("err.Error()")
	}

	m.menuConnect.Find(&menu)
	return menu
}

func (m menuRepository) DeleteMenu(menu entity.Menu) {
	m.menuConnect.Unscoped().Delete(&menu)
}

func (m menuRepository) MenuSearchList(search dto.MenuSearchParam) (data dto.MenuSearchList) {
	departDb:= m.menuConnect.Model(&entity.Menu{})

	if search.Name != "" {
		departDb.Where("name LIKE ?","%"+search.Name+"%")
	}
	if search.State > 0 {
		departDb.Where("state = ? ",search.State)
	}
	var count int64
	departDb.Count(&count)

	MenuList := []dto.Menu{}
	departDb.Offset(int(search.PageIndex - 1) * int(search.PageSize)).Limit(int(search.PageSize)).Find(&MenuList)

	data.Count = count
	data.List = MenuList
	return data
}

func (m menuRepository) FindMenuByID(menuID uint) entity.Menu {
	panic("implement me")
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return  &menuRepository{
		menuConnect: db,
	}
}