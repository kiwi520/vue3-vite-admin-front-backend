package service

import (
	"fmt"
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/repository"
)

type AppVersionService interface {
	Insert(app dto.AppVersionCreateDTO) entity.AppVersion
	Update(app dto.AppVersionUpdateDTO) entity.AppVersion
	Delete(app entity.AppVersion)
	DownloadAppVersionFile(app entity.AppVersion) string
	DeleteAppApk(app dto.DeleteAppApk) error
	SearchList(search dto.AppVersionSearchParam) dto.AppVersionSearchList
	List() []entity.AppVersion
}

type appVersionService struct {
	repository repository.AppVersionRepository
}

func (a appVersionService) DownloadAppVersionFile(app entity.AppVersion) string {
	return  a.repository.DownloadAppVersionFile(app)
}

func (a appVersionService) DeleteAppApk(app dto.DeleteAppApk) error {
	fmt.Println("app")
	fmt.Println(app)
	fmt.Println("app")

	//re := regexp.MustCompile(":8080/")
	//match := re.FindIndex([]byte(app.FilePath))
	//fmt.Println(match)
	//if len(match) == 0 {
	//	fmt.Println("没有匹配的ptah，文件路径有问题")
	//	return errors.New("没有匹配的ptah，文件路径有问题")
	//}
	//content := app.FilePath[match[1] : len(app.FilePath)]
	//fmt.Println(content)
	//
	////app.FilePath = "./"+content
	//app.FilePath = fmt.Sprintf("./%s",content)

	//name, err := helper.GetFileName(app.FilePath, ":8080/")
	//if err != nil {
	//	return err
	//}
	//app.FilePath = fmt.Sprintf("./%s",name)
	err := a.repository.DeleteAppApk(app)

	return err
}

func (a appVersionService) Insert(app dto.AppVersionCreateDTO) entity.AppVersion {
	appToCreate := entity.AppVersion{}
	appToCreate.Name =app.Name
	appToCreate.FilePath =app.FilePath
	appToCreate.Platform =app.Platform
	appToCreate.Version =app.Version
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
	appToUpdate.Version =app.Version
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
