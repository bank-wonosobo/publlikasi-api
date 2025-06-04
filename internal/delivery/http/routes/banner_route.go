package routes

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/handlers"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func RegisterBannerRouter(router fiber.Router, service services.BannerService, validator validator.Validator) {
	handler := handlers.NewBanner(service, validator)

	// public routes
	banners := router.Group("/banners")
	banners.Get("/", handler.Index)

	// admin routes
	admin := router.Group("admin/banners")
	admin.Post("/", handler.Create)
	admin.Get("/", handler.Index)
}
