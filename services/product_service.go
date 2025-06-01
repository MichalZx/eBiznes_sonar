package services

import (
	"go-echo-crud/database"
	"go-echo-crud/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	errCategoryNotFound = "Category not found"
	errProductNotFound  = "Product not found"
)

func CreateProduct(c echo.Context) error {
	product := &models.Product{}
	if err := c.Bind(product); err != nil {
		return err
	}

	var category models.Category
	if err := database.DB.First(&category, product.CategoryId).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, errCategoryNotFound)
	}

	product.Category = category
	database.DB.Create(product)
	return c.JSON(http.StatusCreated, product)
}

func GetProducts(c echo.Context) error {
	var products []models.Product
	database.DB.Preload("Category").Find(&products)
	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid product ID")
	}
	var product models.Product
	if err := database.DB.Preload("Category").First(&product, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, errProductNotFound)
	}
	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	updated := &models.Product{}
	if err := c.Bind(updated); err != nil {
		return err
	}

	var existing models.Product
	if err := database.DB.First(&existing, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, errProductNotFound)
	}

	var category models.Category
	if err := database.DB.First(&category, updated.CategoryId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errCategoryNotFound)
	}

	existing.Name = updated.Name
	existing.Price = updated.Price
	existing.CategoryId = updated.CategoryId
	existing.Category = category

	database.DB.Save(&existing)
	return c.JSON(http.StatusOK, existing)
}

func DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, errProductNotFound)
	}
	database.DB.Delete(&product)
	return c.NoContent(http.StatusNoContent)
}
