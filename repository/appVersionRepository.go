package repository

import (
	"fmt"
	"golang_api/dto"
	"golang_api/entity"
	"gorm.io/gorm"
	"os"
	"regexp"
)

type AppVersionRepository interface {
	InsertAppVersion(app entity.AppVersion) entity.AppVersion
	UpdateAppVersion(app entity.AppVersion) entity.AppVersion
	DeleteAppVersion(app entity.AppVersion)
	DeleteAppApk(app dto.DeleteAppApk) error
	AppVersionSearchList(search dto.AppVersionSearchParam) dto.AppVersionSearchList
	AppVersionList() []entity.AppVersion
}

type appVersionRepository struct {
	appConnect *gorm.DB
}

func (a appVersionRepository) DeleteAppApk(app dto.DeleteAppApk) (err error) {
	var appVersion entity.AppVersion

	if app.Flag {
		appVersion.ID = uint(app.ID)
		err = a.appConnect.Model(&appVersion).UpdateColumn("file_path", "").Error

		if err != nil {
			return err
		}
	}

	err = os.Remove(app.FilePath)


	if err != nil {
		return err
	}

	return err
}

func (a appVersionRepository) InsertAppVersion(app entity.AppVersion) entity.AppVersion {
	a.appConnect.Save(&app)
	a.appConnect.Find(&app)
	return app
}

func (a appVersionRepository) UpdateAppVersion(app entity.AppVersion) entity.AppVersion {
	println("app")
	println(app.ID)
	println(app.Name)
	println(app.Platform)
	println(app.State)
	println(app.Version)
	println(app.Remark)
	println("role")
	err := a.appConnect.Model(&app).Select("*").Omit("id","CreatedAt").Updates(map[string]interface{}{"name": app.Name,"file_path": app.FilePath, "version": app.Version, "platform": app.Platform, "state": app.State, "remark": app.Remark}).Error

	if err != nil{
		println("err.Error()")
		println(err.Error())
		println("err.Error()")
	}

	a.appConnect.Find(&app)
	return app
}

func (a appVersionRepository) DeleteAppVersion(app entity.AppVersion) {

	var appVersion entity.AppVersion
	a.appConnect.Model(&app).Find(&appVersion)

	if appVersion.FilePath != "" {
		re := regexp.MustCompile(":8080/")
		match := re.FindIndex([]byte(appVersion.FilePath))
		fmt.Println(match)
		if len(match) == 0 {
			fmt.Println("没有匹配的ptah，文件路径有问题")
		}
		content := appVersion.FilePath[match[1] : len(appVersion.FilePath)]
		fmt.Println(content)

		err := os.Remove(fmt.Sprintf("./%s",content))


		if err != nil {
			fmt.Println(err.Error())
		}
	}


	a.appConnect.Unscoped().Delete(&app)
}

func (a appVersionRepository) AppVersionSearchList(search dto.AppVersionSearchParam) (data dto.AppVersionSearchList) {
	appDb:= a.appConnect.Model(&entity.AppVersion{})

	if search.Name != "" {
		appDb.Where("name LIKE ?","%"+search.Name+"%")
	}

	if search.Platform > 0 {
		appDb.Where("platform = ? ",search.Platform)
	}

	if search.State > 0 {
		appDb.Where("state = ? ",search.State)
	}

	var count int64
	appDb.Count(&count)

	appList := []dto.AppVersion{}
	appDb.Offset(int(search.PageIndex - 1) * int(search.PageSize)).Limit(int(search.PageSize)).Find(&appList)

	data.Count = count
	data.List = appList
	return data
}

func (a appVersionRepository) AppVersionList() (data []entity.AppVersion) {
	a.appConnect.Find(&data)
	return data
}

func NewAppVersionRepository(db *gorm.DB)  AppVersionRepository{
	return &appVersionRepository{
		appConnect: db,
	}
}