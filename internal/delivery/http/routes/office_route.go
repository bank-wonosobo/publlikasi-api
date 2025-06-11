package routes

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/handlers"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func RegisterOfficeRouter(router fiber.Router, service services.OfficeService, validator validator.Validator) {
	handler := handlers.NewOffice(service, validator)

	// router
	admin := router.Group("/offices")
	admin.Post("/", handler.Create)
	admin.Get("/", handler.GetAll)
	admin.Put("/:id", handler.Update)
}
