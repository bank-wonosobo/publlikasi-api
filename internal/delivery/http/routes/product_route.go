package routes

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/handlers"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func RegisterProductRouter(router fiber.Router, service services.ProductService, validator validator.Validator) {
	handler := handlers.NewProduct(service, validator)

	// public routes
	products := router.Group("/products")
	products.Post("/", handler.Create)
	products.Get("/", handler.Index)
}
