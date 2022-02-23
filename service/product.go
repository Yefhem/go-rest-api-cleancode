package service

import (
	"fmt"
	"log"

	"github.com/Yefhem/rest-api-cleancode/dto"
	"github.com/Yefhem/rest-api-cleancode/models"
	"github.com/Yefhem/rest-api-cleancode/repository"
	"github.com/mashingan/smapping"
)

type ProductService interface {
	Insert(p dto.ProductCreateDTO) models.Product
	Update(p dto.ProductCreateDTO) models.Product
	Delete(p models.Product)
	All() []models.Product
	FindByID(productID uint64) models.Product
	IsAllowedToEdit(userID string, productID uint64) bool
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepo,
	}
}

func (service *productService) Insert(p dto.ProductCreateDTO) models.Product {
	product := models.Product{}
	if err := smapping.FillStruct(&product, smapping.MapFields(&p)); err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.productRepository.InsertProduct(product)
	return res
}

func (service *productService) Update(p dto.ProductCreateDTO) models.Product {
	product := models.Product{}
	if err := smapping.FillStruct(&product, smapping.MapFields(&p)); err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.productRepository.UpdateProduct(product)
	return res
}

func (service *productService) Delete(p models.Product) {
	service.productRepository.DeleteProduct(p)
}

func (service *productService) All() []models.Product {
	return service.productRepository.AllProduct()
}

func (service *productService) FindByID(productID uint64) models.Product {
	return service.productRepository.FindProductByID(productID)
}

// Düzenlemeye izin verir ürünü ekleyen kullanici ile işlem yapmak isteyen kullanicinin id si karşılastırır bool bir değer döner...
func (service *productService) IsAllowedToEdit(userID string, productID uint64) bool {
	p := service.productRepository.FindProductByID(productID)
	id := fmt.Sprintf("%v", p.UserID)
	return userID == id
}
