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

type AdminController interface {
	Insert(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	List(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type adminController struct {
  jwtService	service.JwtService
  adminService service.AdminService
}

func (a adminController) Insert(ctx *gin.Context) {
	var adminCreateDTo dto.AdminCreateDTO
	errDTO := ctx.ShouldBind(&adminCreateDTo)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		result := a.adminService.Insert(adminCreateDTo)
		res := helper.BuildResponse(http.StatusOK, "添加成功", result)
		ctx.JSON(http.StatusOK, res)
	}
}

func (a adminController) FindByID(ctx *gin.Context) {
	panic("implement me")
}

func (a adminController) Update(ctx *gin.Context) {
	var adminUpdateDTO dto.AdminUpdateDTO
	errDTO := ctx.ShouldBind(&adminUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result := a.adminService.Update(adminUpdateDTO)

	res := helper.BuildResponse(http.StatusOK, "修改成功", result)
	ctx.JSON(http.StatusOK, res)
	return
}

func (a adminController) List(ctx *gin.Context) {
	var AdminSearchParam dto.AdminSearchParam
	errDTO := ctx.ShouldBind(&AdminSearchParam)

	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		data := a.adminService.List(AdminSearchParam)

		res := helper.BuildResponse(http.StatusOK, "获取成功", data)
		ctx.JSON(http.StatusOK, res)
	}
}

func (a adminController) Delete(ctx *gin.Context) {
	var admin entity.User
	id, errDTO := strconv.ParseUint(ctx.DefaultQuery("id", "0"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	admin.ID = uint(id)
	a.adminService.Delete(admin)
	res := helper.BuildResponse(http.StatusOK, "删除成功", []string{})
	ctx.JSON(http.StatusOK, res)
	return
}

func NewAdminService(  jwtService	service.JwtService,adminService service.AdminService) AdminController {
	return  &adminController{
		jwtService: jwtService,
		adminService: adminService,
	}
}

