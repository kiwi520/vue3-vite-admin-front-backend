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

type MenuController interface {
	Insert(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	List(ctx *gin.Context)
	GetMenuTreeList(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
type menuController struct {
	jwtService service.JwtService
	menuService service.MenuService
}

func (m menuController) GetMenuTreeList(ctx *gin.Context) {
	pid, errDTO := strconv.ParseUint(ctx.DefaultQuery("pid", "0"), 0, 0)

	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		var deptList []dto.MenuTree = m.menuService.GetMenuTreeList(uint(pid))

		res := helper.BuildResponse(http.StatusOK, "获取成功", deptList)
		ctx.JSON(http.StatusOK, res)
	}
}

func (m menuController) Insert(ctx *gin.Context) {
	var menuCreateDTo dto.MenuCreteDTO
	errDTO := ctx.ShouldBind(&menuCreateDTo)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		result := m.menuService.Insert(menuCreateDTo)
		res := helper.BuildResponse(http.StatusOK, "添加成功", result)
		ctx.JSON(http.StatusOK, res)
	}
}

func (m menuController) FindByID(ctx *gin.Context) {
	panic("implement me")
}

func (m menuController) Update(ctx *gin.Context) {
	var menuUpdateDTO dto.MenuUpdateDTO
	errDTO := ctx.ShouldBind(&menuUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result := m.menuService.Update(menuUpdateDTO)

	res := helper.BuildResponse(http.StatusOK, "修改成功", result)
	ctx.JSON(http.StatusOK, res)
	return
}

func (m menuController) List(ctx *gin.Context) {
	var menuSearchParam dto.MenuSearchParam
	errDTO := ctx.ShouldBind(&menuSearchParam)

	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		data := m.menuService.List(menuSearchParam)

		res := helper.BuildResponse(http.StatusOK, "获取成功", data)
		ctx.JSON(http.StatusOK, res)
	}
}

func (m menuController) Delete(ctx *gin.Context) {
	var Menu entity.Menu
	id, errDTO := strconv.ParseUint(ctx.DefaultQuery("id", "0"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	Menu.ID = uint(id)
	m.menuService.Delete(Menu)
	res := helper.BuildResponse(http.StatusOK, "删除成功", []string{})
	ctx.JSON(http.StatusOK, res)
	return
}

func NewMenuController(jwtService service.JwtService,menuService service.MenuService) MenuController {
	return &menuController{
		jwtService: jwtService,
		menuService: menuService,
	}
}