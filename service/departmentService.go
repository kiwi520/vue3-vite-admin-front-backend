package service

import (
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/respository"
)

type DepartmentService interface {
	Insert(dept dto.DepartmentCreateDTO) entity.Department
	Update(dept dto.DepartmentUpdateDTO) entity.Department
	Delete(dept entity.Department)
	List(search dto.DepartmentSearchParam) dto.DepartmentSearchList
	GetDepartmentTreeList(pid uint) []dto.DepartmentTree
	FindById(deptID uint) entity.Department
}

type departmentService struct {
	departmentRepository respository.DepartmentRepository
}

func (d departmentService) Insert(dept dto.DepartmentCreateDTO) entity.Department {
	deptToCreate := entity.Department{}
	deptToCreate.DepartmentName =dept.DepartmentName
	deptToCreate.DepartmentLeaderID =dept.DepartmentLeaderID
	deptToCreate.ParentID =dept.ParentID

	res := d.departmentRepository.InsertDepartment(deptToCreate)

	return res
}

func (d departmentService) Update(dept dto.DepartmentUpdateDTO) entity.Department {
	deptToUpdate := entity.Department{}

	deptToUpdate.ID = uint(dept.ID)
	deptToUpdate.DepartmentName =dept.DepartmentName
	deptToUpdate.DepartmentLeaderID =dept.DepartmentLeaderID
	deptToUpdate.ParentID =dept.ParentID

	res := d.departmentRepository.UpdateDepartment(deptToUpdate)
	return res
}

func (d departmentService) Delete(dept entity.Department) {
	d.departmentRepository.DeleteDepartment(dept)
}

func (d departmentService) List(search dto.DepartmentSearchParam) dto.DepartmentSearchList {

	return  d.departmentRepository.DepartmentSearchList(search)
}
func (d departmentService) GetDepartmentTreeList(pid uint) []dto.DepartmentTree {
	list:= d.departmentRepository.DepartmentList()

	return  getDepartmentTree(list,pid)
}

func (d departmentService) FindById(deptID uint) entity.Department {
	return d.departmentRepository.FindDepartmentByID(deptID)
}

func NewDepartmentService(repository respository.DepartmentRepository) DepartmentService  {
	return &departmentService{
		departmentRepository: repository,
	}
}

func getDepartmentTree(list []entity.Department, pid uint) []dto.DepartmentTree {
	var DepartmentTree []dto.DepartmentTree
	for _,val := range list {
		if val.ParentID == pid {
			child := getDepartmentTree(list,val.ID)
			node := dto.DepartmentTree {
				ID: val.ID,
				ParentID: val.ParentID,
				DepartmentName: val.DepartmentName,
				DepartmentLeaderID: val.DepartmentLeaderID,
			}
			node.Children = child
			DepartmentTree = append(DepartmentTree,node)
		}
	}

	return  DepartmentTree
}