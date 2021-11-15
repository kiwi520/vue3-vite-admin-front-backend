package service

import (
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/repository"
)

type AppVersionService interface {
	Insert(app dto.AppVersionCreateDTO) entity.AppVersion
	Update(app dto.AppVersionUpdateDTO) entity.AppVersion
	Delete(app entity.AppVersion)
	SearchList(search dto.AppVersionSearchParam) dto.AppVersionSearchList
	List() []entity.AppVersion
}

type appVersionService struct {
	repository repository.AppVersionRepository
}

func (a appVersionService) Insert(app dto.AppVersionCreateDTO) entity.AppVersion {
	appToCreate := entity.AppVersion{}
	appToCreate.Name =app.Name
	appToCreate.FilePath =app.FilePath
	appToCreate.Platform =app.Platform
	appToCreate.State =app.State
	appToCreate.Remark =app.Remark

	res := a.repository.InsertAppVersion(appToCreate)

	return res
}

func (a appVersionService) Update(app dto.AppVersionUpdateDTO) entity.AppVersion {
	appToUpdate := entity.AppVersion{}

	appToUpdate.ID = uint(app.ID)
	appToUpdate.Name =app.Name
	appToUpdate.FilePath =app.FilePath
	appToUpdate.Platform =app.Platform
	appToUpdate.State =app.State
	appToUpdate.Remark =app.Remark

	res := a.repository.UpdateAppVersion(appToUpdate)
	return res
}

func (a appVersionService) Delete(app entity.AppVersion) {
	a.repository.DeleteAppVersion(app)
}

func (a appVersionService) SearchList(search dto.AppVersionSearchParam) dto.AppVersionSearchList {
	return  a.repository.AppVersionSearchList(search)
}

func (a appVersionService) List() []entity.AppVersion {
	return a.repository.AppVersionList()
}

func NewAppVersionService(repository repository.AppVersionRepository) AppVersionService {
	return &appVersionService{
		repository:repository,
	}
}
