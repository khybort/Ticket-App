package main

import (
	"fmt"
	"log"
	"time"
	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"api/internal/config"
	"api/internal/database"
	"api/middleware"
	"api/internal/routers"
	_ "api/docs"
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
	router.SetUpTicketRoutes(r, db)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Backend listening on port %s...", cfg.Port)
	log.Fatal(r.Run(fmt.Sprintf(":%s", cfg.Port)))
}

const (
	maxRetries = 20
	retryDelay = 5 * time.Second
)

func connectWithRetry(cfg *config.Config) *gorm.DB {
	var db *gorm.DB
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
