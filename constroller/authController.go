package constroller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/helper"
	"golang_api/service"
	"net/http"
	"strconv"
)

type AuthController interface {
	Login(ctx * gin.Context)

	Register(ctx * gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService service.JwtService
}

func NewAuthController(authService service.AuthService,jwtService service.JwtService) AuthController {
	return &authController{
		authService: authService,
		jwtService: jwtService,
	}
}

func (c * authController) Login(ctx * gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)

	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误",errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,res)

		return
	}

	authResult := c.authService.VerifyCredential(loginDTO.Email,loginDTO.Password)

	println("authResult")
	println(authResult)
	println("authResult")

	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(uint64(v.ID), 10))
		v.Token = generatedToken
		res := helper.BuildResponse(http.StatusOK,"登录成功",v)
		ctx.JSON(http.StatusOK,res)
	}else{
		res := helper.BuildErrorResponse("认证失败",errors.New(""))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,res)
	}
}

func (c * authController) Register(ctx * gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)

	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误",errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,res)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		res := helper.BuildErrorResponse("邮箱已被注册",errors.New(""))
		ctx.AbortWithStatusJSON(http.StatusConflict,res)
	}else {
		createrUser :=c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(uint64(createrUser.ID),10))
		createrUser.Token = token

		res := helper.BuildResponse(http.StatusOK,"注册成功", createrUser)
		ctx.JSON(http.StatusOK,res)
	}


}