package main

import (
	"log"

	"github.com/bank-wonosobo/publlikasi-api.git/config"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/routes"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/repositories"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/database"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// connect database psql
	db, err := database.NewPsql(cfg)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// auto migrate entity
	if err := db.AutoMigrate(entities.ReportType{}, entities.Report{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// init fiber app
	app := fiber.New(fiber.Config{
		AppName: cfg.App.Name,
	})

	// middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// init repo
	reportTypeRepo := repositories.NewReportType()

	// init service
	reportTypeService := services.NewReportType(db, reportTypeRepo)

	// validator
	validator := validator.NewValidator()

	// API Router
	api := app.Group("/api")
	v1 := api.Group("/v1")
	defaultRoute(v1)

	// init router
	defaultRoute(app)
	routes.RegisterReportTypeRouter(v1, reportTypeService, validator)

	// make server
	log.Printf("Server running on port %s", cfg.App.Port)
	if err := app.Listen(":" + cfg.App.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func defaultRoute(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Default Route Diengs.id API")
	})
}
