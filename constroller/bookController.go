package constroller

import (
	"errors"
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

type BookController interface {
	Insert(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	List(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type bookController struct {
	bookService service.BookService
	jwtService service.JwtService
}

func (b bookController) Insert(ctx *gin.Context) {
	var bookCreateDTo dto.BookCreateDTO
	errDTO := ctx.ShouldBind(&bookCreateDTo)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误",errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,res)
		return
	}else {
		authHeader := ctx.GetHeader("Authorization")
		userId := b.getUserIdByToken(authHeader)
		var userIdToUint, err = strconv.ParseUint(userId, 10, 64)

		if err != nil {
			panic(err.Error())
		}

		bookCreateDTo.UserID = userIdToUint

		result := b.bookService.Insert(bookCreateDTo)

		res := helper.BuildResponse(http.StatusOK,"添加成功",result)
		ctx.JSON(http.StatusOK,res)
	}

}

func (b bookController) FindByID(ctx *gin.Context) {
	id, errDTO := strconv.ParseUint(ctx.Param("id"),0,0)

	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误",errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,res)
		return
	}else {
		var book entity.Book = b.bookService.FindById(uint(id))

		res := helper.BuildResponse(http.StatusOK,"获取成功",book)
		ctx.JSON(http.StatusOK,res)
	}

}

func (b bookController) Update(ctx *gin.Context) {
	var bookUpdateDTO dto.BookUpdateDTO
	errDTO := ctx.ShouldBind(&bookUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误",errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token,errToken := b.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		res := helper.BuildErrorResponse("token不合法",errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,res)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%v",claims["user_id"])
	var userIdToUint, err = strconv.ParseUint(userId, 10, 64)

	if err != nil {
		println("err.Error()")
		panic(err)
	}
	if b.bookService.IsAllowedToEdit(uint(userIdToUint),uint(bookUpdateDTO.ID)) {
		bookUpdateDTO.UserID = userIdToUint

		result := b.bookService.Update(bookUpdateDTO)

		res := helper.BuildResponse(http.StatusOK,"修改成功",result)
		ctx.JSON(http.StatusOK,res)
		return
	}else{
		res := helper.BuildErrorResponse("无权编辑",errors.New(""))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,res)
		return
	}
}

func (b bookController) List(ctx *gin.Context) {
	var bookList []entity.Book = b.bookService.List()

	res := helper.BuildResponse(http.StatusOK,"获取成功",bookList)
	ctx.JSON(http.StatusOK,res)
}

func (b bookController) Delete(ctx *gin.Context) {
	var book entity.Book
	id,errDTO := strconv.ParseUint(ctx.Param("id"),0,0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误",errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,res)
		return
	}

	book.ID = uint(id)

	authHeader := ctx.GetHeader("Authorization")
	userId := b.getUserIdByToken(authHeader)
	var userIdToUint, err = strconv.ParseUint(userId, 10, 64)

	if err != nil {
		panic(err.Error())
	}

	println("userIdToUint")
	println(userIdToUint)
	println("userIdToUint")
	println("book.ID")
	println(book.ID)
	println("book.ID")

	if b.bookService.IsAllowedToEdit(uint(userIdToUint),book.ID) {
		b.bookService.Delete(book)
		res := helper.BuildResponse(http.StatusOK,"删除成功",[]string{})
		ctx.JSON(http.StatusOK,res)
		return
	}else {
		res := helper.BuildErrorResponse("无权删除",errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,res)
		return
	}

}

func NewBookController(bookService service.BookService,jwtService service.JwtService) BookController {
	return &bookController{
		bookService: bookService,
		jwtService: jwtService,
	}
}


//通过token获取用户ID
func (b bookController) getUserIdByToken(token string) string  {
	aToken, err := b.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}

	claims := aToken.Claims.(jwt.MapClaims)
	 id := fmt.Sprintf("%v",claims["user_id"])
	return id
}