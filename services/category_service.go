package services

import (
	"go-echo-crud/database"
	"go-echo-crud/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	CategoryNotFound = "Category not found"
)

func CreateCategory(c echo.Context) error {
	var cat models.Category
	if err := c.Bind(&cat); err != nil {
		return err
	}
	database.DB.Create(&cat)
	return c.JSON(http.StatusCreated, cat)
}

func GetCategories(c echo.Context) error {
	var cats []models.Category
	database.DB.Preload("Products").Find(&cats)
	return c.JSON(http.StatusOK, cats)
}

func GetCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var cat models.Category
	if err := database.DB.Preload("Products").First(&cat, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, CategoryNotFound)
	}
	return c.JSON(http.StatusOK, cat)
}

func DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var cat models.Category
	var count int64
	if err := database.DB.First(&cat, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, CategoryNotFound)
	}
	database.DB.Model(&models.Product{}).Where("category_id = ?", id).Count(&count)
	if count > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "There is a product assigned to the category")
	}
	database.DB.Delete(&cat)
	return c.NoContent(http.StatusNoContent)
}
