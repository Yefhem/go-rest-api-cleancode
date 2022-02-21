package main

import (
	"github.com/Yefhem/rest-api-cleancode/configs"
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
	// service layer
	userService service.UserService = service.NewUserService(userRepository)
)

func main() {

	r := gin.Default()

	r.Run()

}
