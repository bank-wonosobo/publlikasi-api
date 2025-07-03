package routes

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/handlers"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func RegisterComplaintRouter(router fiber.Router, service services.ComplaintService, validator validator.Validator) {
	handler := handlers.NewComplaint(service, validator)

	// public routes
	complaint := router.Group("/complaints")
	complaint.Post("/", handler.Create)

	// admin routes
	admin := router.Group("admin/complaints")
	admin.Get("/", handler.Index)
	admin.Put("/:id/process", handler.Process)
	admin.Put("/:id/done", handler.Done)
}
