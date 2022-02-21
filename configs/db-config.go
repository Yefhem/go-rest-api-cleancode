package configs

import (
	"fmt"
	"log"

	"github.com/Yefhem/rest-api-cleancode/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	var (
		DBURL string
		DB    *gorm.DB
		err   error
	)

	// if err := godotenv.Load(); err != nil {
	// 	log.Fatalf("Error getting env, not comming through %v", err)
	// } else {
	// 	fmt.Println("We are getting the env values")
	// }

	dbDriver := "postgres"                 //os.Getenv("DB_DRIVER")
	dbHost := "127.0.0.1"                  //os.Getenv("DB_HOST")
	dbPort := "5432"                       //os.Getenv("DB_PORT")
	dbUsername := "postgres"               //os.Getenv("DB_USERNAME")
	dbPassword := "admin"                  //os.Getenv("DB_PASSWORD")
	dbName := "my_first_postgres_database" //os.Getenv("DB_NAME")

	if dbDriver == "postgres" {
		DBURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUsername, dbPassword, dbName)

		DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", dbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", dbDriver)
		}
	}

	// if dbDriver == "mysql" {
	// 	DBURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)
	// 	DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	// 	if err != nil {
	// 		fmt.Printf("Cannot connect to %s database", dbDriver)
	// 		log.Fatal("This is the error:", err)
	// 	} else {
	// 		fmt.Printf("We are connected to the %s database", dbDriver)
	// 	}
	// }

	DB.AutoMigrate(&models.User{}, &models.Product{})
	return DB
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
