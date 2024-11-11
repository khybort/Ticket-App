package router

import (
	"api/internal/controllers"
	"api/internal/repositories"
	"api/internal/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpTicketRoutes(r *gin.Engine, db *gorm.DB) {
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
}
