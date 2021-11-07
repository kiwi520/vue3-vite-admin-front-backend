package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang_api/config"
	"golang_api/constroller"
	"golang_api/middleware"
	"golang_api/respository"
	"golang_api/service"
	"gorm.io/gorm"
)

var (
	db * gorm.DB = config.SetupDatabaseConnection()
	userRepository respository.UserRepository = respository.NewUserRepository(db)
	bookRepository respository.BookRepository = respository.NewBookRepository(db)
	departmetRepository respository.DepartmentRepository = respository.NewDepartmentRepository(db)
	roleRepository respository.RoleRepository = respository.NewRoleRepository(db)
	adminRepository respository.AdminRepository = respository.NewAdminRepository(db)

	jwtService service.JwtService = service.NewJwtService()
	bookService service.BookService = service.NewBookService(bookRepository)
	authService service.AuthService = service.NewAuthService(userRepository)
	userService service.UserService = service.NewUserService(userRepository)
	departmentService service.DepartmentService = service.NewDepartmentService(departmetRepository)
	roleService service.RoleService = service.NewRoleService(roleRepository)
	adminService service.AdminService = service.NewAdminService(adminRepository)

	authController constroller.AuthController = constroller.NewAuthController(authService,jwtService)
	userController constroller.UserController = constroller.NewUserController(userService,jwtService)
	bookController constroller.BookController = constroller.NewBookController(bookService,jwtService)
	departmentController constroller.DepartmentController = constroller.NewDepartmentController(departmentService,jwtService)
	roleController constroller.RoleController = constroller.NewRoleController(jwtService,roleService)
	adminController constroller.AdminController = constroller.NewAdminService(jwtService,adminService)
)

func main()  {
	defer config.CloseDatabaseConnection(db)

	r := gin.Default()
	router := gin.Default()

	router.Use(cors.Default())
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	//router.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"https://foo.com"},
	//	AllowMethods:     []string{"PUT", "PATCH"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return origin == "https://github.com"
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))
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
		roleRoutes.DELETE("/", departmentController.Delete)
	}

	adminRoutes := r.Group("api/admin",middleware.AuthorizeJwt(jwtService))
	{
		adminRoutes.POST("/list", adminController.List)
		adminRoutes.POST("/", adminController.Insert)
		//roleRoutes.GET("/:id", roleController.FindByID)
		adminRoutes.PUT("/", adminController.Update)
		adminRoutes.DELETE("/", adminController.Delete)
	}


	r.Run()
}
