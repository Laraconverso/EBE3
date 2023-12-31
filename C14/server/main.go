package main

import (
	"github.com/Laraconverso/EBE3.git/cmd/server/handler"
	"github.com/Laraconverso/EBE3.git/internal/product"
	"github.com/Laraconverso/EBE3.git/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {

	storage := store.NewJsonStore("./products.json")

	repo := product.NewRepository(storage)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET(":id", productHandler.GetByID())
		products.POST("", productHandler.Post())
		products.DELETE(":id", productHandler.Delete())
		products.PATCH(":id", productHandler.Patch())
		products.PUT(":id", productHandler.Put())
	}

	r.Run(":8080")
}
