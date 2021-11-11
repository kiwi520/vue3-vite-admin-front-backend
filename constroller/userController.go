package constroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/helper"
	"golang_api/service"
	"net/http"
	"strconv"
)

type UserController interface {
	Update(ctx *gin.Context)
	Profile(ctx *gin.Context)
	List(ctx *gin.Context)
	GetUserPermission(ctx *gin.Context)
	GetUserButtonList(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService service.JwtService
}

func (u userController) GetUserButtonList(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	token,errToken := u.jwtService.ValidateToken(authHeader)

	if errToken != nil{
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)

	id,err := strconv.ParseUint(fmt.Sprintf("%v",claims["user_id"]),10,64)

	if err != nil {
		res := helper.BuildErrorResponse("无效token",err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,res)
	}

	var menList = u.userService.GetUserButtonList(int64(id))

	res := helper.BuildResponse(http.StatusOK,"获取成功",menList)
	ctx.JSON(http.StatusOK,res)
}

func (u userController) GetUserPermission(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	token,errToken := u.jwtService.ValidateToken(authHeader)

	if errToken != nil{
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)

	id,err := strconv.ParseUint(fmt.Sprintf("%v",claims["user_id"]),10,64)

	if err != nil {
		res := helper.BuildErrorResponse("无效token",err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,res)
	}

	var menList = u.userService.GetUserPermission(int64(id))

	res := helper.BuildResponse(http.StatusOK,"获取成功",menList)
	ctx.JSON(http.StatusOK,res)
}

func (u userController) List(ctx *gin.Context) {
	var userList []entity.User = u.userService.List()

	res := helper.BuildResponse(http.StatusOK,"获取成功",userList)
	ctx.JSON(http.StatusOK,res)
}

func (u userController) Update(ctx *gin.Context) {
	var UserUpdateDTO dto.UserUpdateDTO
	errDTO := ctx.ShouldBind(&UserUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误",errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,res)

		return
	}

	authHeader := ctx.GetHeader("Authorization")

	token,errToken := u.jwtService.ValidateToken(authHeader)

	if errToken != nil{
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)

	id,err := strconv.ParseUint(fmt.Sprintf("%v",claims["user_id"]),10,64)

	if err != nil {
		res := helper.BuildErrorResponse("无效token",err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,res)
	}

	UserUpdateDTO.ID = id

	user := u.userService.Update(UserUpdateDTO)

	res := helper.BuildResponse(http.StatusOK,"更改成功",user)
	ctx.JSON(http.StatusOK,res)

}

func (u userController) Profile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	token,errToken := u.jwtService.ValidateToken(authHeader)

	if errToken != nil{
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)

	id,err := strconv.ParseUint(fmt.Sprintf("%v",claims["user_id"]),10,64)

	if err != nil {
		res := helper.BuildErrorResponse("无效token",err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,res)
	}

	user := u.userService.Profile(strconv.FormatUint(id,10))

	res := helper.BuildResponse(http.StatusOK,"获取成功",user)
	ctx.JSON(http.StatusOK,res)
}

func NewUserController(userService service.UserService,jwtService service.JwtService) UserController {
	return &userController{
		userService: userService,
		jwtService: jwtService,
	}
}