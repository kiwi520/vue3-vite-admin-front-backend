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

type DepartmentController interface {
	Insert(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	List(ctx *gin.Context)
	GetDepartmentTreeList(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type departmentController struct {
	departmentService service.DepartmentService
	jwtService        service.JwtService
}

func (d departmentController) Insert(ctx *gin.Context) {
	var departmentCreateDTo dto.DepartmentCreateDTO
	errDTO := ctx.ShouldBind(&departmentCreateDTo)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		result := d.departmentService.Insert(departmentCreateDTo)
		res := helper.BuildResponse(http.StatusOK, "添加成功", result)
		ctx.JSON(http.StatusOK, res)
	}
}

func (d departmentController) FindByID(ctx *gin.Context) {
	id, errDTO := strconv.ParseUint(ctx.Param("id"), 0, 0)

	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		var dept entity.Department = d.departmentService.FindById(uint(id))

		res := helper.BuildResponse(http.StatusOK, "获取成功", dept)
		ctx.JSON(http.StatusOK, res)
	}
}

func (d departmentController) Update(ctx *gin.Context) {
	var departmentUpdateDTO dto.DepartmentUpdateDTO
	errDTO := ctx.ShouldBind(&departmentUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result := d.departmentService.Update(departmentUpdateDTO)

	res := helper.BuildResponse(http.StatusOK, "修改成功", result)
	ctx.JSON(http.StatusOK, res)
	return
}

func (d departmentController) List(ctx *gin.Context) {
	var DepartmentSearchParam dto.DepartmentSearchParam
	errDTO := ctx.ShouldBind(&DepartmentSearchParam)

	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		data := d.departmentService.List(DepartmentSearchParam)

		res := helper.BuildResponse(http.StatusOK, "获取成功", data)
		ctx.JSON(http.StatusOK, res)
	}
}

func (d departmentController) GetDepartmentTreeList(ctx *gin.Context) {

	pid, errDTO := strconv.ParseUint(ctx.DefaultQuery("pid", "0"), 0, 0)

	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		var deptList []dto.DepartmentTree = d.departmentService.GetDepartmentTreeList(uint(pid))

		res := helper.BuildResponse(http.StatusOK, "获取成功", deptList)
		ctx.JSON(http.StatusOK, res)
	}
}

func (d departmentController) Delete(ctx *gin.Context) {
	var dept entity.Department
	//id, errDTO := strconv.ParseUint(ctx.Param("id"), 0, 0)
	id, errDTO := strconv.ParseUint(ctx.DefaultQuery("id", "0"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	dept.ID = uint(id)
	d.departmentService.Delete(dept)
	res := helper.BuildResponse(http.StatusOK, "删除成功", []string{})
	ctx.JSON(http.StatusOK, res)
	return

	//authHeader := ctx.GetHeader("Authorization")
	//userId := b.getUserIdByToken(authHeader)
	//var userIdToUint, err = strconv.ParseUint(userId, 10, 64)
	//
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//println("userIdToUint")
	//println(userIdToUint)
	//println("userIdToUint")
	//println("book.ID")
	//println(book.ID)
	//println("book.ID")
	//
	//if b.bookService.IsAllowedToEdit(uint(userIdToUint),book.ID) {
	//	b.bookService.Delete(book)
	//	res := helper.BuildResponse(http.StatusOK,"删除成功",[]string{})
	//	ctx.JSON(http.StatusOK,res)
	//	return
	//}else {
	//	res := helper.BuildErrorResponse("无权删除",errDTO)
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest,res)
	//	return
	//}
}

func NewDepartmentController(departmentService service.DepartmentService, jwtService service.JwtService) DepartmentController {
	return &departmentController{
		departmentService: departmentService,
		jwtService:        jwtService,
	}
}
