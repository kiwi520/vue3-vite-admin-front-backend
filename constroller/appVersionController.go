package constroller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/helper"
	"golang_api/service"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type AppVersionController interface {
	Insert(ctx *gin.Context)
	SaveChunk(ctx *gin.Context)
	MergeChunk(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	SearchList(ctx *gin.Context)
	List(ctx *gin.Context)
	Delete(ctx *gin.Context)
	DeleteAppApk(ctx *gin.Context)
	DownloadAppVersionFile(ctx *gin.Context)
}

type appVersionController struct {
	jwtService service.JwtService
	appVersionService service.AppVersionService
}

func (a appVersionController) DeleteAppApk(ctx *gin.Context) {

	var appDeleteAppApk dto.DeleteAppApk
	errDTO := ctx.ShouldBind(&appDeleteAppApk)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err := a.appVersionService.DeleteAppApk(appDeleteAppApk)
	if err != nil {
		res := helper.BuildErrorResponse("删除失败", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(http.StatusOK, "删除成功", dto.DeleteAppApk{})
	ctx.JSON(http.StatusOK, res)
}

func (a appVersionController) MergeChunk(ctx *gin.Context) {
	var appVersionMergeChunk dto.AppVersionMergeChunk
	errDTO := ctx.ShouldBind(&appVersionMergeChunk)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	hasPath := fmt.Sprintf("./%s/%s/",os.Getenv("Article_App_Path"),appVersionMergeChunk.FileHash)
	hasFile := fmt.Sprintf("./%s/%s",os.Getenv("Article_App_Path"),appVersionMergeChunk.FileName)

	isExistPath, err := helper.PathExists(hasPath)

	if err != nil {
		fmt.Println("获取hash路径失败",err.Error())
	}

	if !isExistPath {
		res := helper.BuildErrorResponse("文件夹不存在", errors.New(""))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	isExistFile,err := helper.PathExists(hasFile)

	if isExistFile {
		type Res struct {
			FileUrl string `json:"file_url"`
		}

		res := helper.BuildResponse(http.StatusOK, "合并成功", Res{
			FileUrl: fmt.Sprintf("%s://%s:%s/%s/%s",helper.GetHttps(),helper.GetLocalIP()[0],os.Getenv("PORT"),os.Getenv("Article_App_Path"),appVersionMergeChunk.FileName),
			//FileUrl: fmt.Sprintf("http://127.0.0.1:8080/uploadFile/%s",appVersionMergeChunk.FileName),
		})
		ctx.JSON(http.StatusOK, res)
		return
	}

    files,err := ioutil.ReadDir(hasPath)

	if err != nil {
		res := helper.BuildErrorResponse("创建文件读取失败", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// 创建文件
	complateFile,err := os.Create(hasFile)
	if err != nil {
		res := helper.BuildErrorResponse("创建文件失败", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	defer func(complateFile *os.File) {
		err := complateFile.Close()
		if err != nil {
			res := helper.BuildErrorResponse("关闭complateFile失败", err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}(complateFile)

	for _,file := range files{
		if file.Name() == ".DS_Store" {
			continue
		}

		//读取分片数据

		fileBuffer, err := ioutil.ReadFile(hasPath + "/" + file.Name())

		if err != nil {
			res := helper.BuildErrorResponse("打开文件失败", err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		_, err = complateFile.Write(fileBuffer)
		if err != nil {
			res := helper.BuildErrorResponse("写入文件失败", err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		err = os.Remove(hasPath + "/" + file.Name())

		if err != nil{
			res := helper.BuildErrorResponse("删除hasPath目录失败", err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}

	err = os.RemoveAll(hasPath)

	if err != nil{
		res := helper.BuildErrorResponse("删除hasPath目录失败", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	type Res struct {
		FileUrl string `json:"file_url"`
		FileName string `json:"file_name"`
		Flag bool `json:"flag"`
	}

	res := helper.BuildResponse(http.StatusOK, "合并成功", Res{
		//FileUrl: fmt.Sprintf("http://127.0.0.1:8080/uploadFile/%s",appVersionMergeChunk.FileName),
		FileUrl: fmt.Sprintf("./%s/%s",os.Getenv("Article_App_Path"),appVersionMergeChunk.FileName),
		FileName: appVersionMergeChunk.FileName,
		Flag:false,

	})
	ctx.JSON(http.StatusOK, res)
	return

}

func (a appVersionController) SaveChunk(ctx *gin.Context) {
	file, _ := ctx.FormFile("chunk")
	fileHash := ctx.PostForm("fileHash")
	chunkHash := ctx.PostForm("chunkHash")
	chunkNumber := ctx.PostForm("chunkNumber")

	fmt.Println("fileHash")
	fmt.Println(fileHash)
	fmt.Println("fileHash")
	fmt.Println("chunkHash")
	fmt.Println(chunkHash)
	fmt.Println("chunkHash")
	fmt.Println("chunkNumber")
	fmt.Println(chunkNumber)
	fmt.Println("chunkNumber")

	hashPath := fmt.Sprintf("./%s/%s/",os.Getenv("Article_App_Path"),fileHash)

	isExistPath, err := helper.PathExists(hashPath)

	if err != nil {
		fmt.Println("获取hash路径失败",err.Error())
	}

	if !isExistPath {
		err := os.Mkdir(hashPath, os.ModePerm)
		if err != nil {
			fmt.Println("创建hash路径失败",err.Error())
		}
	}

	err = ctx.SaveUploadedFile(file, fmt.Sprintf("./%s/%s/%s",os.Getenv("Article_App_Path"), fileHash, chunkHash))
	if err != nil {
		res := helper.BuildErrorResponse("保存失败", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}else{
		type Res struct {
			ChunkNumber string `json:"chunk_number"`
		}

		res := helper.BuildResponse(http.StatusOK, "添加成功", Res{
			ChunkNumber: chunkNumber,
		})
		ctx.JSON(http.StatusOK, res)
	}
	//5765eafc8b2ca541a598b2d8ee0c799b
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

//TODO 文件下载
func (a appVersionController) DownloadAppVersionFile(c *gin.Context) {
	var app entity.AppVersion
	id, errDTO := strconv.ParseUint(c.DefaultQuery("id", "0"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("请求参数有误", errDTO)
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	app.ID = uint(id)

	println("app")
	println(app.ID)
	println("app")

	filePath :=a.appVersionService.DownloadAppVersionFile(app)

	if filePath == "" {
		res := helper.BuildErrorResponse("没有附件文件", errDTO)
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	//打开文件
	//_, errByOpenFile := os.Open(fileDir + "/" + fileName)
	//非空处理

	fileName, err := helper.GetFileName(filePath, "/uploadFile/")
	if err != nil {
		res := helper.BuildErrorResponse("获取文件名失败", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	fmt.Println("fileName")
	fmt.Println(fileName)
	fmt.Println(filePath)
	fmt.Println("fileName")

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	//浏览器下载或预览
	c.Header("Content-Disposition", "inline;filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.File(filePath)

	//c.FileAttachment("./uploadFile/Firefox_90.0.2.dmg",fileName)
	return

	//res := helper.BuildResponse(http.StatusOK, "获取成功", filePath)
	//c.JSON(http.StatusOK, res)
	//return
}

func NewAppVersionController(	jwtService service.JwtService,appVersionService service.AppVersionService) AppVersionController {
	return &appVersionController{
		jwtService: jwtService,
		appVersionService: appVersionService,
	}
}

