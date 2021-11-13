package service

import (
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/repository"
)

type AdminService interface {
	Insert(admin dto.AdminCreateDTO) entity.User
	Update(admin dto.AdminUpdateDTO) entity.User
	Delete(admin entity.User)
	List(search dto.AdminSearchParam) dto.AdminSearchList
}

type adminService struct {
	repository.AdminRepository
}

func (a adminService) Insert(admin dto.AdminCreateDTO) entity.User {
	adminToCreate := entity.User{}
	adminToCreate.Name =admin.Name
	adminToCreate.Email =admin.Email
	adminToCreate.Mobile =admin.Mobile
	adminToCreate.Gender =admin.Gender
	adminToCreate.DepartmentID =admin.DepartmentID
	adminToCreate.RoleID =admin.RoleID
	adminToCreate.State =admin.State
	adminToCreate.Job =admin.Job
	adminToCreate.Remark =admin.Remark
	adminToCreate.Password =admin.Password

	res := a.AdminRepository.InsertAdmin(adminToCreate)

	return res
}

func (a adminService) Update(admin dto.AdminUpdateDTO) entity.User {
	adminToUpdate := entity.User{}

	adminToUpdate.ID = uint(admin.ID)
	adminToUpdate.Name =admin.Name
	adminToUpdate.Email =admin.Email
	adminToUpdate.Mobile =admin.Mobile
	adminToUpdate.Gender =admin.Gender
	adminToUpdate.DepartmentID =admin.DepartmentID
	adminToUpdate.RoleID =admin.RoleID
	adminToUpdate.State =admin.State
	adminToUpdate.Job =admin.Job
	adminToUpdate.Remark =admin.Remark

	res := a.AdminRepository.UpdateAdmin(adminToUpdate)
	return res
}

func (a adminService) Delete(admin entity.User) {
	a.AdminRepository.DeleteAdmin(admin)
}

func (a adminService) List(search dto.AdminSearchParam) dto.AdminSearchList {
	return  a.AdminRepository.AdminSearchList(search)
}

func NewAdminService (repository repository.AdminRepository) AdminService  {
	return &adminService{
		repository,
	}
}
