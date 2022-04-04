package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vincentconace/concurrencia-product/cmd/server/handler"
	"github.com/vincentconace/concurrencia-product/internal"
	"github.com/vincentconace/concurrencia-product/pkg/db"
)

func main() {
	//Connection database
	db, err := db.DB()
	if err != nil {
		panic(err)
	}

	//Start server
	app := fiber.New()

	app.Use(logger.New())

	repositoryProduct := internal.NewRepository(db)
	serviceProduct := internal.NewService(repositoryProduct)
	handlerProduct := handler.NewHandler(serviceProduct)

	r := app.Group("/products")

	r.Get("/", handlerProduct.GetAll)
	r.Get("/:id", handlerProduct.GetById)
	r.Post("/", handlerProduct.Create)
	r.Put("/:id", handlerProduct.Update)
	r.Delete("/:id", handlerProduct.Delete)
	app.Listen(":8080")

}
