package constroller

import (
	"github.com/gin-gonic/gin"
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/helper"
	"golang_api/service"
	"net/http"
	"strconv"
)

type AppVersionController interface {
	Insert(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	SearchList(ctx *gin.Context)
	List(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type appVersionController struct {
	jwtService service.JwtService
	appVersionService service.AppVersionService
}

func (a appVersionController) Insert(ctx *gin.Context) {
	var appCreateDTo dto.AppVersionCreateDTO
	errDTO := ctx.ShouldBind(&appCreateDTo)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		result := a.appVersionService.Insert(appCreateDTo)
		res := helper.BuildResponse(http.StatusOK, "添加成功", result)
		ctx.JSON(http.StatusOK, res)
	}
}

func (a appVersionController) FindByID(ctx *gin.Context) {
	panic("implement me")
}

func (a appVersionController) Update(ctx *gin.Context) {
	var appUpdateDTO dto.AppVersionUpdateDTO
	errDTO := ctx.ShouldBind(&appUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result := a.appVersionService.Update(appUpdateDTO)

	res := helper.BuildResponse(http.StatusOK, "修改成功", result)
	ctx.JSON(http.StatusOK, res)
	return
}

func (a appVersionController) SearchList(ctx *gin.Context) {
	var appSearchParam dto.AppVersionSearchParam
	errDTO := ctx.ShouldBind(&appSearchParam)

	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		data := a.appVersionService.SearchList(appSearchParam)

		res := helper.BuildResponse(http.StatusOK, "获取成功", data)
		ctx.JSON(http.StatusOK, res)
	}
}

func (a appVersionController) List(ctx *gin.Context) {
	var appList []entity.AppVersion = a.appVersionService.List()

	res := helper.BuildResponse(http.StatusOK, "获取成功", appList)
	ctx.JSON(http.StatusOK, res)
}

func (a appVersionController) Delete(ctx *gin.Context) {
	var app entity.AppVersion
	id, errDTO := strconv.ParseUint(ctx.DefaultQuery("id", "0"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	app.ID = uint(id)

	println("app")
	println(app.ID)
	println("app")
	a.appVersionService.Delete(app)
	res := helper.BuildResponse(http.StatusOK, "删除成功", []string{})
	ctx.JSON(http.StatusOK, res)
	return
}

func NewAppVersionController(	jwtService service.JwtService,appVersionService service.AppVersionService) AppVersionController {
	return &appVersionController{
		jwtService: jwtService,
		appVersionService: appVersionService,
	}
}

