package main

import (
	"github.com/Yefhem/rest-api-cleancode/configs"
	"github.com/Yefhem/rest-api-cleancode/controller"
	"github.com/Yefhem/rest-api-cleancode/middleware"
	"github.com/Yefhem/rest-api-cleancode/repository"
	"github.com/Yefhem/rest-api-cleancode/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	// db
	db *gorm.DB = configs.SetupDatabaseConnection()
	// repository layer
	userRepository    repository.UserRepository    = repository.NewUserRepository(db)
	productRepository repository.ProductRepository = repository.NewProductRepository(db)
	// services layer
	jwtService     service.JWTService     = service.NewJWTService()
	userService    service.UserService    = service.NewUserService(userRepository)
	productService service.ProductService = service.NewProductService(productRepository)
	authService    service.AuthService    = service.NewAuthService(userRepository)
	// controller layer
	authController    controller.AuthController    = controller.NewAuthController(authService, jwtService)
	userController    controller.UserController    = controller.NewUserController(userService, jwtService)
	productController controller.ProductController = controller.NewProductController(productService, jwtService)
)

func main() {
	defer configs.CloseDatabaseConnection(db)

	router := gin.Default()

	authRoutes := router.Group("api/v1/auth")
	{
		authRoutes.GET("/home", authController.Home)          // ok
		authRoutes.POST("/login", authController.Login)       // ok
		authRoutes.POST("/register", authController.Register) // ok
	}

	userRoutes := router.Group("api/v1/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile) // ok
		userRoutes.PUT("/profile", userController.Update)  // ok
	}

	productRoutes := router.Group("api/v1/products", middleware.AuthorizeJWT(jwtService))
	{
		productRoutes.GET("/", productController.All)          // ok
		productRoutes.POST("/", productController.Insert)      // ok
		productRoutes.GET("/:id", productController.FindByID)  // ok
		productRoutes.PUT("/:id", productController.Update)    // ok
		productRoutes.DELETE("/:id", productController.Delete) // ok
	}

	router.Run()

}
