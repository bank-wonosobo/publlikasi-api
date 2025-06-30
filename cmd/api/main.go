package main

import (
	"log"

	"github.com/bank-wonosobo/publlikasi-api.git/config"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/routes"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/repositories"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/database"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/storage"
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

	// init s3
	s3Storage := storage.NewS3Storage(
		cfg.S3.BucketName,
		cfg.S3.Region,
		cfg.S3.Endpoint,
		cfg.S3.AccessKey,
		cfg.S3.SecretKey,
	)

	// auto migrate entity
	if err := db.AutoMigrate(
		entities.ReportType{},
		entities.Report{},
		entities.News{},
		entities.Announcement{},
		entities.Product{},
		entities.Banner{},
		entities.Office{},
		entities.ComplaintType{},
		entities.Complaint{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// init fiber app
	app := fiber.New(fiber.Config{
		AppName: cfg.App.Name,
	})

	// middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: false,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	// app.Use(cors.New(cors.Config{
	// 	// AllowOriginsFunc: func(origin string) bool {
	// 	// 	allowedOrigins := []string{
	// 	// 		"http://localhost:3000",
	// 	// 		"http://localhost:5173",
	// 	// 		"http://bankwonosobo.co.id",
	// 	// 		"http://adminer.bankwonosobo.co.id",
	// 	// 	}
	// 	// 	for _, o := range allowedOrigins {
	// 	// 		if origin == o {
	// 	// 			return true
	// 	// 		}
	// 	// 	}
	// 	// 	return false
	// 	// },
	// 	// AllowOrigins: "",
	// 	// AllowMethods:     "GET,POST,PUT,DELETE",
	// 	// AllowHeaders:     "*",
	// 	// AllowCredentials: true,
	// }))

	// init repo
	reportTypeRepo := repositories.NewReportType(db)
	reportRepo := repositories.NewReport(db)
	newsRepo := repositories.NewNews(db)
	announcementRepo := repositories.NewAnnouncement(db)
	productRepo := repositories.NewProduct(db)
	bannerRepo := repositories.NewBanner(db)
	officeRepo := repositories.NewOffice(db)
	complaintTypeRepo := repositories.NewComplaintType(db)

	// init service
	reportTypeService := services.NewReportType(reportTypeRepo)
	reportService := services.NewReport(reportRepo, reportTypeRepo, s3Storage)
	newsService := services.NewNews(newsRepo, s3Storage)
	announcementServce := services.NewAnnouncement(announcementRepo, s3Storage)
	productService := services.NewProduct(productRepo, s3Storage)
	bannerService := services.NewBanner(bannerRepo, s3Storage)
	officeSerice := services.NewOffice(officeRepo, s3Storage)
	complaintTypeSerice := services.NewComplaintType(complaintTypeRepo)

	// validator
	validator := validator.NewValidator()
	validator.RegisterCustomValidor()

	// API Router
	api := app.Group("/api")
	v1 := api.Group("/v1")
	defaultRoute(v1)

	// init router
	defaultRoute(app)
	routes.RegisterReportTypeRouter(v1, reportTypeService, validator)
	routes.RegisterReportRouter(v1, reportService, validator)
	routes.RegisterNewsRouter(v1, newsService, validator)
	routes.RegisterAnnouncementRouter(v1, announcementServce, validator)
	routes.RegisterProductRouter(v1, productService, validator)
	routes.RegisterBannerRouter(v1, bannerService, validator)
	routes.RegisterOfficeRouter(v1, officeSerice, validator)
	routes.RegisterComplaintTypeRouter(v1, complaintTypeSerice, validator)

	// make server
	log.Printf("Server running on port %s", cfg.App.Port)
	if err := app.Listen(":" + cfg.App.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func defaultRoute(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Default Route Publikasi API")
	})
}
