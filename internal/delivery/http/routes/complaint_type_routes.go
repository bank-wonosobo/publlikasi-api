package routes

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/handlers"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

// register complaint type router
func RegisterComplaintTypeRouter(router fiber.Router, service services.ComplaintTypeService, validator validator.Validator) {
	handler := handlers.NewComplaintType(service, validator)

	// admin routes
	admin := router.Group("admin/complaint-types")
	admin.Get("/", handler.Index)
	admin.Post("/", handler.Create)
	admin.Put("/:id", handler.Update)
	admin.Delete("/:id", handler.Delete)
}
