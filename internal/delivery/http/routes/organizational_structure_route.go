package routes

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/handlers"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func RegisterOrganizationalStructureRouter(router fiber.Router, service services.OrganizationalStructureService, validator validator.Validator) {
	handler := handlers.NewOrganizationalStructure(service, validator)

	items := router.Group("/organizational-structures")
	items.Post("/", handler.Create)
	items.Get("/", handler.GetAll)
	items.Put("/:id", handler.Update)
	items.Delete("/:id", handler.Delete)
}
