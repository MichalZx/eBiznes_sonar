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

func getCartByID(c echo.Context) (*models.Cart, error) {
	cartID, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid cart ID")
	}
	var cart models.Cart
	if err := database.DB.Preload("Products").First(&cart, cartID).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, CartNotFound)
	}
	return &cart, nil
}

func getProductByID(c echo.Context) (*models.Product, error) {
	productID, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid product ID")
	}
	var product models.Product
	if err := database.DB.First(&product, productID).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, ProductNotFound)
	}
	return &product, nil
}

func CreateCart(c echo.Context) error {
	var cart models.Cart
	if err := c.Bind(&cart); err != nil {
		return err
	}
	database.DB.Create(&cart)
	return c.JSON(http.StatusCreated, cart)
}

func AddProductToCart(c echo.Context) error {
	cart, err := getCartByID(c)
	if err != nil {
		return err
	}
	product, err := getProductByID(c)
	if err != nil {
		return err
	}
	database.DB.Model(cart).Association("Products").Append(product)
	return c.JSON(http.StatusOK, cart)
}

func GetCart(c echo.Context) error {
	cart, err := getCartByID(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, cart)
}

func DeleteCart(c echo.Context) error {
	cartID, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid cart ID")
	}
	var cart models.Cart
	if err := database.DB.First(&cart, cartID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, CartNotFound)
	}
	database.DB.Model(&cart).Association("Products").Clear()
	database.DB.Delete(&cart)
	return c.NoContent(http.StatusNoContent)
}

func RemoveProductFromCart(c echo.Context) error {
	cart, err := getCartByID(c)
	if err != nil {
		return err
	}
	product, err := getProductByID(c)
	if err != nil {
		return err
	}
	database.DB.Model(cart).Association("Products").Delete(product)
	return c.JSON(http.StatusOK, cart)
}
