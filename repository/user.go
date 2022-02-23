package repository

import (
	"log"

	"github.com/Yefhem/rest-api-cleancode/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user models.User) models.User
	UpdateUser(user models.User) models.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) models.User
	ProfileUser(userID string) models.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

// Methodlar -------------

// User password hash işlemi...
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

// Parolayı doğrulama
// func VerifyPassword(hashedPassword, password string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// 	if err != nil {
// 		log.Println(err)
// 		fmt.Println("parola doğrulanamadı")
// 		return false
// 	}
// 	return true
// }

// User password yukarıdaki hashAndSalt fonk. ile hashlendikten sonra db'ye kaydedilir...
func (db *userConnection) InsertUser(user models.User) models.User {
	user.Password = HashPassword(user.Password)
	db.connection.Create(&user)
	return user
}

func (db *userConnection) UpdateUser(user models.User) models.User {
	if user.Password != "" {
		user.Password = HashPassword(user.Password)
	} else {
		var tempUser models.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}
	db.connection.Save(&user)
	return user
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user models.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) models.User {
	var user models.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) models.User {
	var user models.User
	db.connection.Preload("Products").Preload("Products.User").Find(&user, userID)
	return user
}

// Kimlik Doğrulama...
func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user models.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}
