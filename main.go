package main

import (
	"go-echo-crud/database"
	"go-echo-crud/services"

	"github.com/labstack/echo/v4"
)

const (
	products         = "/products"
	productsId       = "/products/:id"
	categories       = "/categories"
	categoriesId     = "/categories/:id"
	cartId           = "/carts/:cartId"
	cartIdProductsId = "/carts/:cartId/products/:productId"
)

func main() {
	e := echo.New()
	database.ConnectDB()

	e.POST(products, services.CreateProduct)
	e.GET(products, services.GetProducts)
	e.GET(productsId, services.GetProduct)
	e.PUT(productsId, services.UpdateProduct)
	e.DELETE(productsId, services.DeleteProduct)

	e.POST(categories, services.CreateCategory)
	e.GET(categories, services.GetCategories)
	e.GET(categoriesId, services.GetCategory)
	e.DELETE(categoriesId, services.DeleteCategory)

	e.POST("/carts", services.CreateCart)
	e.GET(cartId, services.GetCart)
	e.DELETE(cartId, services.DeleteCart)
	e.POST(cartIdProductsId, services.AddProductToCart)
	e.DELETE(cartIdProductsId, services.RemoveProductFromCart)

	e.Logger.Fatal(e.Start(":8080"))
}
