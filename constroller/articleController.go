package constroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/helper"
	"golang_api/service"
	"net/http"
	"os"
	"strconv"
)

type ArticleController interface {
	Insert(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	SaveImg(ctx *gin.Context)
	Update(ctx *gin.Context)
	SearchList(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type articleController struct {
	jwtService service.JwtService
	articleService service.ArticleService
}

func (a articleController) SaveImg(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")

	fmt.Println("file")
	fmt.Println(file.Filename)
	fmt.Println("file")


	isExistPath, err := helper.PathExists(os.Getenv("Article_Img_Path"))

	if err != nil {
		fmt.Println("获取hash路径失败",err.Error())
	}

	if !isExistPath {
		err := os.Mkdir("./uploadFile/images/", os.ModePerm)
		if err != nil {
			fmt.Println("创建hash路径失败",err.Error())
		}
	}

	err = ctx.SaveUploadedFile(file, fmt.Sprintf(os.Getenv("Article_Img_Path")+"/%s", file.Filename))
	if err != nil {
		res := helper.BuildErrorResponse("保存失败", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}else{
		type Res struct {
			ImgPath string `json:"img_path"`
		}

		imgPath := fmt.Sprintf("%s/%s/%s",os.Getenv("Article_Img_Path_Host"),os.Getenv("Article_Img_Path"),file.Filename)
		res := helper.BuildResponse(http.StatusOK, "添加成功", Res{
			ImgPath: imgPath,
		})
		ctx.JSON(http.StatusOK, res)
	}
}

func (a articleController) Insert(ctx *gin.Context) {
	var param dto.ArticleCreateDTO
	errDTO := ctx.ShouldBind(&param)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		result := a.articleService.Insert(param)
		res := helper.BuildResponse(http.StatusOK, "添加成功", result)
		ctx.JSON(http.StatusOK, res)
	}
}

func (a articleController) FindByID(ctx *gin.Context) {
	id, errDTO := strconv.ParseUint(ctx.Param("id"), 0, 0)

	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		var cate entity.Article = a.articleService.FindById(uint(id))

		res := helper.BuildResponse(http.StatusOK, "获取成功", cate)
		ctx.JSON(http.StatusOK, res)
	}
}

func (a articleController) Update(ctx *gin.Context) {
	var param dto.ArticleUpdateDTO
	errDTO := ctx.ShouldBind(&param)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result := a.articleService.Update(param)

	res := helper.BuildResponse(http.StatusOK, "修改成功", result)
	ctx.JSON(http.StatusOK, res)
	return
}

func (a articleController) SearchList(ctx *gin.Context) {
	var param dto.ArticleSearchParam
	errDTO := ctx.ShouldBind(&param)

	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		data := a.articleService.SearchList(param)

		res := helper.BuildResponse(http.StatusOK, "获取成功", data)
		ctx.JSON(http.StatusOK, res)
	}
}

func (a articleController) Delete(ctx *gin.Context) {
	var article entity.Article
	id, errDTO := strconv.ParseUint(ctx.DefaultQuery("id", "0"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	article.ID = uint(id)
	a.articleService.Delete(article)
	res := helper.BuildResponse(http.StatusOK, "删除成功", []string{})
	ctx.JSON(http.StatusOK, res)
	return
}

func NewArticleController(	jwtService service.JwtService,articleService service.ArticleService) ArticleController {
	return articleController{
		jwtService: jwtService,
		articleService: articleService,
	}
}
