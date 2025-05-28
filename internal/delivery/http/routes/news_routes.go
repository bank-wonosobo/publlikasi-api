package routes

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/handlers"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func RegisterNewsRouter(router fiber.Router, service services.NewsService, validator validator.Validator) {
	handler := handlers.NewNews(service, validator)

	// public routes
	news := router.Group("/news")
	news.Get("/", handler.GetAll)
	news.Get("/:id", handler.Detail)
	news.Get("/:slug/slug", handler.DetailWSlug)

	// admin router
	admin := router.Group("/admin/news")
	admin.Post("/", handler.Create)
	admin.Get("/", handler.GetAll)
	admin.Put("/:id", handler.Update)
	admin.Delete("/:id", handler.Delete)
	admin.Put("/:id/approve", handler.Approve)
	admin.Put("/:id/archive", handler.Archive)
	admin.Get("/:id", handler.Detail)
	admin.Get("/:slug/slug", handler.DetailWSlug)
}
