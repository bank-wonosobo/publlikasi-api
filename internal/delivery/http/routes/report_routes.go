package routes

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/handlers"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func RegisterReportRouter(router fiber.Router, service services.ReportService, validator validator.Validator) {
	handler := handlers.NewReport(service, validator)

	// public routes
	report := router.Group("reports")
	report.Get("/test", handler.Create)
	// admin routes
	admin := router.Group("admin/reports")
	admin.Post("/", handler.Create)
}
