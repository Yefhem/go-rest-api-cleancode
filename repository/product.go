package repository

import (
	"github.com/Yefhem/rest-api-cleancode/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertProduct(product models.Product) models.Product
	UpdateProduct(product models.Product) models.Product
	DeleteProduct(product models.Product)
	FindProductByID(pID uint64) models.Product
	AllProduct() []models.Product
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productConnection{
		connection: db,
	}
}

// ----------------------------

// -Yeni bir ürün oluşturma
func (db *productConnection) InsertProduct(product models.Product) models.Product {
	db.connection.Create(&product)
	db.connection.Preload("User").Find(&product)
	return product
}

// UpdateProduct
func (db *productConnection) UpdateProduct(product models.Product) models.Product {
	db.connection.Save(&product)
	db.connection.Preload("User").Find(&product)
	return product
}

// DeleteProduct
func (db *productConnection) DeleteProduct(product models.Product) {
	db.connection.Delete(&product)
}

// FindProductByID
func (db *productConnection) FindProductByID(pID uint64) models.Product {
	var product models.Product
	db.connection.Preload("User").Find(&product, pID)
	return product
}

// AllProduct
func (db *productConnection) AllProduct() []models.Product {
	var products []models.Product
	db.connection.Preload("User").Find(&products)
	return products
}
