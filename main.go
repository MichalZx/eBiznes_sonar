package main

import (
	"go-echo-crud/database"
	"go-echo-crud/services"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	database.ConnectDB()

	e.POST("/products", services.CreateProduct)
	e.GET("/products", services.GetProducts)
	e.GET("/products/:id", services.GetProduct)
	e.PUT("/products/:id", services.UpdateProduct)
	e.DELETE("/products/:id", services.DeleteProduct)

	e.POST("/categories", services.CreateCategory)
	e.GET("/categories", services.GetCategories)
	e.GET("/categories/:id", services.GetCategory)
	e.DELETE("/categories/:id", services.DeleteCategory)

	e.POST("/carts", services.CreateCart)
	e.GET("/carts/:cartId", services.GetCart)
	e.DELETE("/carts/:cartId", services.DeleteCart)
	e.POST("/carts/:cartId/products/:productId", services.AddProductToCart)
	e.DELETE("/carts/:cartId/products/:productId", services.RemoveProductFromCart)

	e.Logger.Fatal(e.Start(":8080"))
}
