package service

import (
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/respository"
)

type RoleService interface {
	Insert(role dto.RoleInsertDTO) entity.Role
	Update(role dto.RoleUpdateDTO) entity.Role
	Delete(role entity.Role)
	List(search dto.RoleSearchParam) dto.RoleSearchList
	//GetDepartmentTreeList(pid uint) []dto.Role
	AllList() []entity.Role
	FindById(roleID uint) entity.Role
}

type roleService struct {
  respository.RoleRepository
}

func (r roleService) AllList() []entity.Role {
	return r.RoleRepository.RoleList()
}

func (r roleService) Insert(role dto.RoleInsertDTO) entity.Role {
	roleToCreate := entity.Role{}
	roleToCreate.RoleName =role.RoleName
	roleToCreate.Remark =role.Remark

	res := r.RoleRepository.InsertRole(roleToCreate)

	return res
}

func (r roleService) Update(role dto.RoleUpdateDTO) entity.Role {
	roleToUpdate := entity.Role{}

	roleToUpdate.ID = uint(role.ID)
	roleToUpdate.RoleName =role.RoleName
	roleToUpdate.Remark =role.Remark

	res := r.RoleRepository.UpdateRole(roleToUpdate)
	return res
}

func (r roleService) Delete(role entity.Role) {
	r.RoleRepository.DeleteRole(role)
}

func (r roleService) List(search dto.RoleSearchParam) dto.RoleSearchList {
	return  r.RoleRepository.RoleSearchList(search)
}

func (r roleService) FindById(roleID uint) entity.Role {
	panic("implement me")
}

func NewRoleService(rep respository.RoleRepository) RoleService {
	return &roleService{
		rep,
	}
}
