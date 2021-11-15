package repository

import (
	"golang_api/dto"
	"golang_api/entity"
	"gorm.io/gorm"
)

type AppVersionRepository interface {
	InsertAppVersion(app entity.AppVersion) entity.AppVersion
	UpdateAppVersion(app entity.AppVersion) entity.AppVersion
	DeleteAppVersion(app entity.AppVersion)
	AppVersionSearchList(search dto.AppVersionSearchParam) dto.AppVersionSearchList
	AppVersionList() []entity.AppVersion
}

type appVersionRepository struct {
	appConnect *gorm.DB
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