package respository

import (
	"golang_api/dto"
	"golang_api/entity"
	"gorm.io/gorm"
)

type RoleRepository interface {
	InsertRole(role entity.Role) entity.Role
	UpdateRole(role entity.Role) entity.Role
	DeleteRole(role entity.Role)
	RoleSearchList(search dto.RoleSearchParam) dto.RoleSearchList
	//DepartmentList() []entity.Role
	FindRoleByID(departID uint) entity.Role
}

type roleRepository struct {
	roleConnection *gorm.DB
}

func (r roleRepository) InsertRole(role entity.Role) entity.Role {
	r.roleConnection.Save(&role)
	r.roleConnection.Find(&role)
	return role
}

func (r roleRepository) UpdateRole(role entity.Role) entity.Role {
	println("role")
	println(role.ID)
	println(role.RoleName)
	println(role.Remark)
	println("role")
	err := r.roleConnection.Model(&role).Select("*").Omit("id", "CreatedAt").Updates(map[string]interface{}{"role_name": role.RoleName, "remark": role.Remark}).Error

	if err != nil{
		println("err.Error()")
		println(err.Error())
		println("err.Error()")
	}

	r.roleConnection.Find(&role)
	return role
}

func (r roleRepository) DeleteRole(role entity.Role) {
	r.roleConnection.Unscoped().Delete(&role)
}

func (r roleRepository) RoleSearchList(search dto.RoleSearchParam) (data dto.RoleSearchList) {
	//departDb:= r.roleConnection.Table("departments as d").Select("d.id,d.department_name,d.department_leader_id,d.parent_id,d.created_at,d.updated_at, ( select u.name  from users as u where u.id=d.department_leader_id) as leader_name, ( select u.email  from users as u where u.id=d.department_leader_id) as leader_email ").Omit("DeletedAt").Where(&entity.Department{DepartmentName: search.DepartmentName})

	departDb:= r.roleConnection.Model(&entity.Role{})

	if search.RoleName != "" {
		departDb.Where("role_name LIKE ?","%"+search.RoleName+"%")
	}
	var count int64
	departDb.Count(&count)

	RoleList := []dto.Role{}
	departDb.Offset(int(search.PageIndex - 1) * int(search.PageSize)).Limit(int(search.PageSize)).Find(&RoleList)

	data.Count = count
	data.List = RoleList
	return data
}

//func (r roleRepository) DepartmentList() []entity.Role {
//	panic("implement me")
//}

func (r roleRepository) FindRoleByID(departID uint) entity.Role {
	panic("implement me")
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		roleConnection: db,
	}
}