package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"api/internal/config"
	"api/internal/controllers"
	"api/internal/database"
	"api/internal/repositories"
	"api/internal/usecases"
	"api/middleware"

	_ "api/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Ticket API
// @version		1.0
// @description	This is a simple API to manage tickets
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8000
// @BasePath		/api/v1
func main() {
	cfg := config.LoadConfig()

	db := connectWithRetry(cfg)
	if err := database.SeedDatabase(db); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	ticketRepo := repositories.NewTicketRepo(db)
	ticketGetAllUseCase := usecases.NewTicketGetAllUseCase(ticketRepo)
	ticketCreateUseCase := usecases.NewTicketCreateUseCase(ticketRepo)
	ticketGetByIDUseCase := usecases.NewTicketGetUseCase(ticketRepo)
	ticketPurchaseUseCase := usecases.NewTicketPurchaseUseCase(ticketRepo)

	ticketHandler := controllers.NewTicketHandler(
		ticketGetAllUseCase,
		ticketCreateUseCase,
		ticketGetByIDUseCase,
		ticketPurchaseUseCase,
	)

	r := gin.Default()
	r.SetTrustedProxies([]string{cfg.UIAddress})

	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.ErrorHandlerMiddleware())
	// CORS settings
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.UIAddress},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// API routes
	api := r.Group("/api/v1")
	{
		api.GET("/tickets", func(c *gin.Context) {
			ticketHandler.GetAllTickets(c)
		})
		api.GET("/tickets/:id", func(c *gin.Context) {
			ticketHandler.GetTicket(c)
		})
		api.POST("/tickets", func(c *gin.Context) {
			ticketHandler.CreateTicket(c)
		})
		api.POST("/tickets/:id/purchases", func(c *gin.Context) {
			ticketHandler.PurchaseTicket(c)
		})
	}

	// Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Backend listening on port %s...", cfg.Port)
	log.Fatal(r.Run(fmt.Sprintf(":%s", cfg.Port)))
}

const (
	maxRetries = 20
	retryDelay = 5 * time.Second
)

func connectWithRetry(cfg *config.Config) *sql.DB {
	var db *sql.DB
	var err error

	for attempts := 0; attempts < maxRetries; attempts++ {
		db, err = database.ConnectDB(cfg)
		if err == nil {
			log.Printf("Connected to the database successfully with %d attempts.", attempts)
			return db
		}

		log.Printf("Database connection attempt %d failed: %v. Retrying in %v...", attempts+1, err, retryDelay)
		time.Sleep(retryDelay)
	}

	log.Fatal("Max retries reached, could not connect to the database")
	return nil
}
