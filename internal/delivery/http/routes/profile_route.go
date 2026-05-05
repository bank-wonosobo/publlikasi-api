package routes

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/handlers"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func RegisterProfileRouter(router fiber.Router, service services.ProfileService, validator validator.Validator) {
	handler := handlers.NewProfile(service, validator)

	profiles := router.Group("/profiles")
	profiles.Post("/", handler.Create)
	profiles.Get("/", handler.GetAll)
	profiles.Put("/:id", handler.Update)
	profiles.Delete("/:id", handler.Delete)
}
