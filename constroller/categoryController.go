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

type CategoryController interface {
	Insert(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	SearchList(ctx *gin.Context)
	GetTreeList(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type categoryController struct {
	jwtService service.JwtService
	cateService service.CategoryService
}

func (c categoryController) Insert(ctx *gin.Context) {
	var categoryCreateDTO dto.CategoryCreateDTO
	errDTO := ctx.ShouldBind(&categoryCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		result := c.cateService.Insert(categoryCreateDTO)
		res := helper.BuildResponse(http.StatusOK, "添加成功", result)
		ctx.JSON(http.StatusOK, res)
	}
}

func (c categoryController) FindByID(ctx *gin.Context) {
	id, errDTO := strconv.ParseUint(ctx.Param("id"), 0, 0)

	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		var cate entity.Category = c.cateService.FindById(uint(id))

		res := helper.BuildResponse(http.StatusOK, "获取成功", cate)
		ctx.JSON(http.StatusOK, res)
	}
}

func (c categoryController) Update(ctx *gin.Context) {
	var categoryUpdateDTO dto.CategoryUpdateDTO
	errDTO := ctx.ShouldBind(&categoryUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result := c.cateService.Update(categoryUpdateDTO)

	res := helper.BuildResponse(http.StatusOK, "修改成功", result)
	ctx.JSON(http.StatusOK, res)
	return
}

func (c categoryController) SearchList(ctx *gin.Context) {
	var categorySearchParam dto.CategorySearchParam
	errDTO := ctx.ShouldBind(&categorySearchParam)

	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		data := c.cateService.SearchList(categorySearchParam)

		res := helper.BuildResponse(http.StatusOK, "获取成功", data)
		ctx.JSON(http.StatusOK, res)
	}
}

func (c categoryController) GetTreeList(ctx *gin.Context) {
	pid, errDTO := strconv.ParseUint(ctx.DefaultQuery("pid", "0"), 0, 0)

	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		var cateList []dto.CategoryTree = c.cateService.GetTreeList(uint(pid))

		res := helper.BuildResponse(http.StatusOK, "获取成功", cateList)
		ctx.JSON(http.StatusOK, res)
	}
}

func (c categoryController) Delete(ctx *gin.Context) {
	var cate entity.Category
	id, errDTO := strconv.ParseUint(ctx.DefaultQuery("id", "0"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	cate.ID = uint(id)
	c.cateService.Delete(cate)
	res := helper.BuildResponse(http.StatusOK, "删除成功", []string{})
	ctx.JSON(http.StatusOK, res)
	return
}

func NewCategoryController(	jwtService service.JwtService,cateService service.CategoryService) CategoryController {
	return &categoryController{
		jwtService: jwtService,
		cateService: cateService,
	}
}