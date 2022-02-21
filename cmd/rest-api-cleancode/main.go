package main

import (
	"github.com/Yefhem/rest-api-cleancode/configs"
	"github.com/Yefhem/rest-api-cleancode/controller"
	"github.com/Yefhem/rest-api-cleancode/repository"
	"github.com/Yefhem/rest-api-cleancode/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	// db
	db *gorm.DB = configs.SetupDatabaseConnection()
	// repository layer
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	// services layer
	userService service.UserService = service.NewUserService(userRepository)
	jwtService  service.JWTService  = service.NewJWTService()
	// controller layer
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {
	defer configs.CloseDatabaseConnection(db)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	userRoutes := r.Group("api/v1/user")
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	r.Run()

}
