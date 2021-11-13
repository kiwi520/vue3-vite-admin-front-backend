package repository

import (
	"golang_api/dto"
	"golang_api/entity"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	InsertDepartment(dept entity.Department) entity.Department
	UpdateDepartment(dept entity.Department) entity.Department
	DeleteDepartment(dept entity.Department)
	DepartmentSearchList(search dto.DepartmentSearchParam) dto.DepartmentSearchList
	DepartmentList() []entity.Department
	FindDepartmentByID(departID uint) entity.Department
}

type departmentRepository struct {
	departConnection *gorm.DB
}

func (d departmentRepository) DepartmentSearchList(search dto.DepartmentSearchParam) (data dto.DepartmentSearchList) {


	//departDb:= d.departConnection.Model(&entity.Department{}).Select("").Where(&entity.Department{DepartmentName: search.DepartmentName})
	departDb:= d.departConnection.Table("departments as d").Select("d.id,d.department_name,d.department_leader_id,d.parent_id,d.created_at,d.updated_at, ( select u.name  from users as u where u.id=d.department_leader_id) as leader_name, ( select u.email  from users as u where u.id=d.department_leader_id) as leader_email ").Omit("DeletedAt")

	if search.DepartmentName != "" {
		departDb.Where("d.department_name LIKE ? ", "%"+search.DepartmentName+"%")
	}

	var count int64
	departDb.Count(&count)

	DepartmentList := []dto.Department{}
	departDb.Offset(int(search.PageIndex - 1) * int(search.PageSize)).Limit(int(search.PageSize)).Find(&DepartmentList)

	data.Count = count
	data.List = DepartmentList
	return data
}

func (d departmentRepository) InsertDepartment(dept entity.Department) entity.Department {
	d.departConnection.Save(&dept)
	d.departConnection.Find(&dept)
	return dept
}

func (d departmentRepository) UpdateDepartment(dept entity.Department) entity.Department {
	println("dept")
	println(dept.ID)
	println(dept.DepartmentLeaderID)
	println(dept.DepartmentName)
	println(dept.ParentID)
	println("dept")
	err := d.departConnection.Model(&dept).Select("*").Omit("id", "CreatedAt").Updates(map[string]interface{}{"department_name": dept.DepartmentName, "department_leader_id": dept.DepartmentLeaderID, "parent_id": dept.ParentID}).Error

	if err != nil{
		println("err.Error()")
		println(err.Error())
		println("err.Error()")
	}


	d.departConnection.Find(&dept)
	return dept
}

func (d departmentRepository) DeleteDepartment(dept entity.Department) {
	d.departConnection.Unscoped().Delete(&dept)
}

func (d departmentRepository) DepartmentList() []entity.Department {
	var deptList = []entity.Department{}
	d.departConnection.Find(&deptList)
	return deptList
}

func (d departmentRepository) FindDepartmentByID(departID uint) entity.Department {
	var dept = entity.Department{}
	d.departConnection.Find(&dept,departID)
	return dept
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return  &departmentRepository{
		departConnection: db,
	}
}