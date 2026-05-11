package routes

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/handlers"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func RegisterVisionMissionRouter(router fiber.Router, service services.VisionMissionService, validator validator.Validator) {
	handler := handlers.NewVisionMission(service, validator)

	items := router.Group("/vision-missions")
	items.Post("/", handler.Create)
	items.Get("/", handler.GetAll)
	items.Put("/:id", handler.Update)
	items.Delete("/:id", handler.Delete)
}
