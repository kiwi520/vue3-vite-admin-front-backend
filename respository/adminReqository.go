package respository

import (
	"golang_api/dto"
	"golang_api/entity"
	"gorm.io/gorm"
)

type AdminRepository interface {
	InsertAdmin(admin entity.User) entity.User
	UpdateAdmin(admin entity.User) entity.User
	DeleteAdmin(admin entity.User)
	AdminSearchList(search dto.AdminSearchParam) dto.AdminSearchList
}

type adminRepository struct {
	adminConnection *gorm.DB
}

func (a adminRepository) InsertAdmin(admin entity.User) entity.User {
	a.adminConnection.Save(&admin)
	a.adminConnection.Find(&admin)
	return admin
}

func (a adminRepository) UpdateAdmin(admin entity.User) entity.User {
	println("admin")
	println(admin.ID)
	println(admin.Name)
	println(admin.Remark)
	println("role")
	err := a.adminConnection.Model(&admin).Select("*").Omit("id", "CreatedAt").Updates(map[string]interface{}{"name": admin.Name,"email": admin.Email,"mobile": admin.Mobile,"gender": admin.Gender,"department_id": admin.DepartmentID,"role_id": admin.RoleID,"state": admin.State,"job": admin.Job,"password": admin.Password, "remark": admin.Remark}).Error

	if err != nil{
		println("err.Error()")
		println(err.Error())
		println("err.Error()")
	}

	a.adminConnection.Find(&admin)
	return admin
}

func (a adminRepository) DeleteAdmin(admin entity.User) {
	a.adminConnection.Unscoped().Delete(&admin)
}

func (a adminRepository) AdminSearchList(search dto.AdminSearchParam) (data dto.AdminSearchList) {
	departDb:= a.adminConnection.Model(&entity.User{})
	if search.Name != "" {
		departDb.Where("name LIKE ? ", "%"+search.Name+"%")
	}

	if search.ID > 0 {
		departDb.Where("id = ? ",search.ID)
	}

	if search.State > 0  {
		departDb.Where("state = ? ",search.State)
	}

	var count int64
	departDb.Count(&count)

	AdminList := []dto.Admin{}
	departDb.Offset(int(search.PageIndex - 1) * int(search.PageSize)).Limit(int(search.PageSize)).Find(&AdminList)

	data.Count = count
	data.List = AdminList
	return data
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return  &adminRepository{
		adminConnection: db,
	}
}