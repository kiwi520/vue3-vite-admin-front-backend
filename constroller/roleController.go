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

type RoleController interface {
	Insert(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	List(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type roleController struct {
	jwtService service.JwtService
	roleService service.RoleService
}

func (r roleController) Insert(ctx *gin.Context) {
	var roleCreateDTo dto.RoleInsertDTO
	errDTO := ctx.ShouldBind(&roleCreateDTo)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		result := r.roleService.Insert(roleCreateDTo)
		res := helper.BuildResponse(http.StatusOK, "添加成功", result)
		ctx.JSON(http.StatusOK, res)
	}
}

func (r roleController) FindByID(ctx *gin.Context) {
	panic("implement me")
}

func (r roleController) Update(ctx *gin.Context) {
	var roleUpdateDTO dto.RoleUpdateDTO
	errDTO := ctx.ShouldBind(&roleUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result := r.roleService.Update(roleUpdateDTO)

	res := helper.BuildResponse(http.StatusOK, "修改成功", result)
	ctx.JSON(http.StatusOK, res)
	return
}

func (r roleController) List(ctx *gin.Context) {
	var RoleSearchParam dto.RoleSearchParam
	errDTO := ctx.ShouldBind(&RoleSearchParam)

	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		data := r.roleService.List(RoleSearchParam)

		res := helper.BuildResponse(http.StatusOK, "获取成功", data)
		ctx.JSON(http.StatusOK, res)
	}
}

func (r roleController) Delete(ctx *gin.Context) {
	var Role entity.Role
	id, errDTO := strconv.ParseUint(ctx.DefaultQuery("id", "0"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	Role.ID = uint(id)
	r.roleService.Delete(Role)
	res := helper.BuildResponse(http.StatusOK, "删除成功", []string{})
	ctx.JSON(http.StatusOK, res)
	return
}

func NewRoleController(	jwtService service.JwtService,roleService service.RoleService) RoleController {
	return &roleController{
		jwtService: jwtService,
		roleService: roleService,
	}
}
