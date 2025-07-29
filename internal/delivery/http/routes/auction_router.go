package routes

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/handlers"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func RegisterAuctionRouter(router fiber.Router, service services.AuctionService, validator validator.Validator) {
	handler := handlers.NewAuction(service, validator)

	// public routes
	// auction := router.Group("/auctions")
	// auction.Get("/", handler.Create)

	// // admin routes
	admin := router.Group("admin/auctions")
	admin.Post("/", handler.Create)
	// admin.Get("/", handler.Index)
	// admin.Put("/:id", handler.Update)
	// admin.Delete("/:id", handler.Delete)
	// admin.Put("/:id/activate", handler.Activate)
}
