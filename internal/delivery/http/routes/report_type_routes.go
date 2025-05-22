package routes

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/handlers"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

// register report router
func RegisterReportTypeRouter(router fiber.Router, service services.ReportTypeService, validator validator.Validator) {
	handler := handlers.NewReportType(service, validator)

	// public routes
	reportTypes := router.Group("report-types")
	reportTypes.Get("/", handler.Index)

	// admin routes
	admin := router.Group("admin/report-types")
	admin.Get("/", handler.Index)
	admin.Post("/", handler.Create)
	admin.Put("/:id", handler.Update)
	admin.Delete("/:id", handler.Delete)
}
