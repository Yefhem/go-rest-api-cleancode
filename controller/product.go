package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Yefhem/rest-api-cleancode/dto"
	"github.com/Yefhem/rest-api-cleancode/helper"
	"github.com/Yefhem/rest-api-cleancode/models"
	"github.com/Yefhem/rest-api-cleancode/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	All(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type productController struct {
	productService service.ProductService
	jwtService     service.JWTService
}

func NewProductController(productService service.ProductService, jwtService service.JWTService) ProductController {
	return &productController{
		productService: productService,
		jwtService:     jwtService,
	}
}

func (cont *productController) All(ctx *gin.Context) {
	var products []models.Product = cont.productService.All()
	res := helper.BuildResponse(true, "OK", products)
	ctx.JSON(200, res)
}

func (cont *productController) FindByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var product models.Product = cont.productService.FindByID(id)
	if (product == models.Product{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", product)
		ctx.JSON(http.StatusOK, res)
	}
}

func (cont *productController) Insert(ctx *gin.Context) {
	var productCreateDTO dto.ProductCreateDTO
	if errDTO := ctx.ShouldBind(&productCreateDTO); errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(400, res)
	} else {
		authHeader := ctx.GetHeader("Authorization")
		userID := cont.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			productCreateDTO.UserID = convertedUserID
		}
		result := cont.productService.Insert(productCreateDTO)
		res2 := helper.BuildResponse(true, "OK", result)
		ctx.JSON(201, res2)
	}
	/*
		{
			"title":"Gezer Erkek Terlik",
			"description":"35-46 numara arası konforlu, plastik, dikişli erkek terliği.",
			"price":50,
			"stock":10,
			"stock_status":true
		}
	*/
}

func (cont *productController) Update(ctx *gin.Context) {
	var productUpdateDTO dto.ProductUpdateDTO
	errDTO := ctx.ShouldBind(&productUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(400, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := cont.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if cont.productService.IsAllowedToEdit(userID, productUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			productUpdateDTO.UserID = id
		}
		result := cont.productService.Update(productUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
	}

}

func (cont *productController) Delete(ctx *gin.Context) {
	var product models.Product
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id wew found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}
	product.ID = id
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := cont.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if cont.productService.IsAllowedToEdit(userID, product.ID) {
		cont.productService.Delete(product)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
	}

}

func (cont *productController) getUserIDByToken(token string) string {
	aToken, err := cont.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
