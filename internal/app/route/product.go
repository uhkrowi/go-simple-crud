package route

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/uhkrowi/go-simple-crud/internal/app/handler"
	"github.com/uhkrowi/go-simple-crud/internal/app/usecase"
	"gorm.io/gorm"
)

func ProductRoute(route fiber.Router, db *gorm.DB, validate *validator.Validate) {
	productUseCase := usecase.NewProductUseCase(db, validate)
	productHandler := handler.NewProductHandler(productUseCase)

	route.Get("", productHandler.GetList)
	route.Get("/:id", productHandler.GetSingle)
	route.Post("", productHandler.Create)
	route.Put("/:id", productHandler.Update)
	route.Delete("/:id", productHandler.Delete)
}
