package services

import (
	"go-echo-crud/database"
	"go-echo-crud/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	CartNotFound    = "Cart not found"
	ProductNotFound = "Product not found"
)

func CreateCart(c echo.Context) error {
	var cart models.Cart
	if err := c.Bind(&cart); err != nil {
		return err
	}
	database.DB.Create(&cart)
	return c.JSON(http.StatusCreated, cart)
}

func AddProductToCart(c echo.Context) error {
	cartID, _ := strconv.Atoi(c.Param("cartId"))
	productID, _ := strconv.Atoi(c.Param("productId"))
	var cart models.Cart
	var product models.Product
	if err := database.DB.Preload("Products").First(&cart, cartID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, CartNotFound)
	}
	if err := database.DB.First(&product, productID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, ProductNotFound)
	}
	database.DB.Model(&cart).Association("Products").Append(&product)
	return c.JSON(http.StatusOK, cart)
}

func GetCart(c echo.Context) error {
	cartID, _ := strconv.Atoi(c.Param("cartId"))
	var cart models.Cart
	if err := database.DB.Preload("Products").First(&cart, cartID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, CartNotFound)
	}
	return c.JSON(http.StatusOK, cart)
}

func DeleteCart(c echo.Context) error {
	cartID, _ := strconv.Atoi(c.Param("cartId"))
	var cart models.Cart
	if err := database.DB.First(&cart, cartID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, CartNotFound)
	}
	database.DB.Model(&cart).Association("Products").Clear()
	database.DB.Delete(&cart)
	return c.NoContent(http.StatusNoContent)
}

func RemoveProductFromCart(c echo.Context) error {
	cartID, _ := strconv.Atoi(c.Param("cartId"))
	productID, _ := strconv.Atoi(c.Param("productId"))
	var cart models.Cart
	var product models.Product
	if err := database.DB.Preload("Products").First(&cart, cartID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, CartNotFound)
	}
	if err := database.DB.First(&product, productID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, ProductNotFound)
	}
	database.DB.Model(&cart).Association("Products").Delete(&product)
	return c.JSON(http.StatusOK, cart)
}
