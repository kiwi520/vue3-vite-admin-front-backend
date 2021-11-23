package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang_api/config"
	"golang_api/constroller"
	"golang_api/middleware"
	"golang_api/repository"
	"golang_api/service"
	"gorm.io/gorm"
)

var (
	db                  * gorm.DB                       = config.SetupDatabaseConnection()
	userRepository      repository.UserRepository       = repository.NewUserRepository(db)
	bookRepository      repository.BookRepository       = repository.NewBookRepository(db)
	departmetRepository repository.DepartmentRepository = repository.NewDepartmentRepository(db)
	roleRepository      repository.RoleRepository       = repository.NewRoleRepository(db)
	adminRepository     repository.AdminRepository      = repository.NewAdminRepository(db)
	menuRepository      repository.MenuRepository       = repository.NewMenuRepository(db)
	appVersionRepository repository.AppVersionRepository = repository.NewAppVersionRepository(db)
	categoryRepository repository.CategoryRepository = repository.NewCategoryRepository(db)
	articleRepository repository.ArticleRepository = repository.NewArticleRepository(db)

	jwtService service.JwtService = service.NewJwtService()
	bookService service.BookService = service.NewBookService(bookRepository)
	authService service.AuthService = service.NewAuthService(userRepository)
	userService service.UserService = service.NewUserService(userRepository)
	departmentService service.DepartmentService = service.NewDepartmentService(departmetRepository)
	roleService service.RoleService = service.NewRoleService(roleRepository)
	adminService service.AdminService = service.NewAdminService(adminRepository)
	menuService service.MenuService = service.NewMenuService(menuRepository)
	appVersionService service.AppVersionService = service.NewAppVersionService(appVersionRepository)
	categoryService service.CategoryService = service.NewCategoryService(categoryRepository)
	articleService service.ArticleService = service.NewArticleService(articleRepository)

	authController constroller.AuthController = constroller.NewAuthController(authService,jwtService)
	userController constroller.UserController = constroller.NewUserController(userService,jwtService)
	bookController constroller.BookController = constroller.NewBookController(bookService,jwtService)
	departmentController constroller.DepartmentController = constroller.NewDepartmentController(departmentService,jwtService)
	roleController constroller.RoleController = constroller.NewRoleController(jwtService,roleService)
	adminController constroller.AdminController = constroller.NewAdminService(jwtService,adminService)
	menuController constroller.MenuController = constroller.NewMenuController(jwtService,menuService)
	appVersionController constroller.AppVersionController = constroller.NewAppVersionController(jwtService,appVersionService)
	categoryController constroller.CategoryController = constroller.NewCategoryController(jwtService,categoryService)
	articleController constroller.ArticleController = constroller.NewArticleController(jwtService,articleService)
)

func RouteMap() (*gin.Engine, *gorm.DB) {

	r := gin.Default()
	router := gin.Default()

	router.Use(cors.Default())


	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register",authController.Register)
	}

	userRoutes := r.Group("api/user",middleware.AuthorizeJwt(jwtService))
	{
		userRoutes.GET("/profile",userController.Profile)
		userRoutes.GET("/list",userController.List)
		userRoutes.POST("/update",userController.Update)
		userRoutes.GET("/permission",userController.GetUserPermission)
		userRoutes.GET("/button",userController.GetUserButtonList)
	}

	bookRoutes := r.Group("api/book",middleware.AuthorizeJwt(jwtService))
	{
		bookRoutes.GET("/list",bookController.List)
		bookRoutes.POST("/",bookController.Insert)
		bookRoutes.GET("/:id",bookController.FindByID)
		bookRoutes.PUT("/",bookController.Update)
		bookRoutes.DELETE("/:id",bookController.Delete)
	}

	departmentRoutes := r.Group("api/department",middleware.AuthorizeJwt(jwtService))
	{
		departmentRoutes.POST("/list", departmentController.List)
		departmentRoutes.GET("/tree", departmentController.GetDepartmentTreeList)
		departmentRoutes.POST("/", departmentController.Insert)
		departmentRoutes.GET("/:id", departmentController.FindByID)
		departmentRoutes.PUT("/", departmentController.Update)
		departmentRoutes.DELETE("/", departmentController.Delete)
	}

	roleRoutes := r.Group("api/role",middleware.AuthorizeJwt(jwtService))
	{
		roleRoutes.POST("/list", roleController.List)
		roleRoutes.POST("/", roleController.Insert)
		//roleRoutes.GET("/:id", roleController.FindByID)
		roleRoutes.PUT("/", roleController.Update)
		roleRoutes.PUT("/setPermission", roleController.SetPermission)
		roleRoutes.DELETE("/", roleController.Delete)
	}

	adminRoutes := r.Group("api/admin",middleware.AuthorizeJwt(jwtService))
	{
		adminRoutes.POST("/list", adminController.List)
		adminRoutes.POST("/", adminController.Insert)
		roleRoutes.GET("/all", roleController.AllList)
		//roleRoutes.GET("/:id", roleController.FindByID)
		adminRoutes.PUT("/", adminController.Update)
		adminRoutes.DELETE("/", adminController.Delete)
	}


	menuRoutes := r.Group("api/menu",middleware.AuthorizeJwt(jwtService))
	{
		menuRoutes.POST("/list", menuController.List)
		menuRoutes.GET("/tree", menuController.GetMenuTreeList)
		menuRoutes.POST("/", menuController.Insert)
		//roleRoutes.GET("/:id", menuController.FindByID)
		menuRoutes.PUT("/", menuController.Update)
		menuRoutes.DELETE("/", menuController.Delete)
	}


	appVersionRoutes := r.Group("api/appVersion",middleware.AuthorizeJwt(jwtService))
	{
		appVersionRoutes.POST("/list", appVersionController.SearchList)
		appVersionRoutes.POST("/uploadChunk", appVersionController.SaveChunk)
		appVersionRoutes.POST("/mergeChunk", appVersionController.MergeChunk)
		appVersionRoutes.GET("/downloadAppVersionFile", appVersionController.DownloadAppVersionFile)
		appVersionRoutes.POST("/", appVersionController.Insert)
		appVersionRoutes.PUT("/", appVersionController.Update)
		appVersionRoutes.DELETE("/", appVersionController.Delete)
		appVersionRoutes.DELETE("/deleteAppApk", appVersionController.DeleteAppApk)
	}


	categorytRoutes := r.Group("api/category",middleware.AuthorizeJwt(jwtService))
	{
		categorytRoutes.POST("/list", categoryController.SearchList)
		categorytRoutes.GET("/tree", categoryController.GetTreeList)
		categorytRoutes.POST("/", categoryController.Insert)
		categorytRoutes.GET("/:id", categoryController.FindByID)
		categorytRoutes.PUT("/", categoryController.Update)
		categorytRoutes.DELETE("/", categoryController.Delete)
	}

	articleRoutes := r.Group("api/article",middleware.AuthorizeJwt(jwtService))
	{
		articleRoutes.POST("/list", articleController.SearchList)
		articleRoutes.POST("/", articleController.Insert)
		articleRoutes.POST("/saveImg", articleController.SaveImg)
		articleRoutes.GET("/:id", articleController.FindByID)
		articleRoutes.PUT("/", articleController.Update)
		articleRoutes.DELETE("/", articleController.Delete)
		articleRoutes.DELETE("/deleteArticleImg", articleController.DeleteArticleImg)
	}

	return r,db
}