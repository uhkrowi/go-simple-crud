package main

import (
	"os"

	"github.com/go-playground/validator"
	"github.com/uhkrowi/go-simple-crud/config"
	"github.com/uhkrowi/go-simple-crud/internal/app/route"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	defer config.CloseDBConnection()

	app := fiber.New(fiber.Config{
		AppName: "CRUD",
	})

	app.Use(cors.New())

	setup(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}

func setup(app *fiber.App) {
	config.InitDB()
	db := config.DBConn
	validate := validator.New()

	apiV1 := app.Group("/api/v1")

	route.ProductRoute(apiV1.Group("/product"), db, validate)
}
