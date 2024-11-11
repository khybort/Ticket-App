package controllers

import (
	"strconv"
	"api/internal/domain"
	"api/internal/usecases"
	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	TicketGetAllUseCase   usecases.TicketGetAllUseCase
	TicketCreateUseCase   usecases.TicketCreateUseCase
	TicketGetUseCase      usecases.TicketGetUseCase
	TicketPurchaseUseCase usecases.TicketPurchaseUseCase
}

func NewTicketHandler(
	ticketGetAllUseCase usecases.TicketGetAllUseCase,
	ticketCreateUseCase usecases.TicketCreateUseCase,
	ticketGetUseCase usecases.TicketGetUseCase,
	ticketPurchaseUseCase usecases.TicketPurchaseUseCase,
) *TicketHandler {
	return &TicketHandler{
		TicketGetAllUseCase:   ticketGetAllUseCase,
		TicketCreateUseCase:   ticketCreateUseCase,
		TicketGetUseCase:      ticketGetUseCase,
		TicketPurchaseUseCase: ticketPurchaseUseCase,
	}
}

// GetAllTickets godoc
//	@Summary		Get all tickets
//	@Description	Retrieves all tickets
//	@Tags			tickets
//	@Produce		json
//	@Success		200	{array}		domain.Ticket
//	@Failure		500	{object}	domain.JSONResponse	"Internal Server Error"
//	@Router			/tickets [get]
func (h *TicketHandler) GetAllTickets(c *gin.Context) {
	tickets, err := h.TicketGetAllUseCase.Execute()
	if err != nil {
		c.JSON(500, domain.JSONResponse{Message: err.Error()})
		return
	}

	c.JSON(200, tickets)
}

// GetTicket godoc
//	@Summary		Get a ticket
//	@Description	Retrieves a ticket by ID
//	@Tags			tickets
//	@Param			id	path	int	true	"Ticket ID"
//	@Produce		json
//	@Success		200	{object}	domain.Ticket
//	@Failure		400	{object}	domain.JSONResponse	"Bad Request"
//	@Failure		404	{object}	domain.JSONResponse	"Ticket not found"
//	@Router			/tickets/{id} [get]
func (h *TicketHandler) GetTicket(c *gin.Context) {
	ticketID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, domain.JSONResponse{Message: "Invalid ticket ID"})
		return
	}

	ticket, err := h.TicketGetUseCase.Execute(ticketID)
	if err != nil {
		c.JSON(404, domain.JSONResponse{Message: "Ticket not found"})
		return
	}

	c.JSON(200, ticket)
}

// CreateTicket godoc
//	@Summary		Create a ticket
//	@Description	Creates a new ticket
//	@Tags			tickets
//	@Accept			json
//	@Produce		json
//	@Param			ticket	body		domain.CreateTicketRequest	true	"Ticket creation data"
//	@Success		201		{object}	domain.Ticket
//	@Failure		400		{object}	domain.JSONResponse	"Invalid request body"
//	@Router			/tickets [post]
func (h *TicketHandler) CreateTicket(c *gin.Context) {
	var createReq domain.CreateTicketRequest
	if err := c.ShouldBindJSON(&createReq); err != nil {
		c.JSON(400, domain.JSONResponse{Message: "Invalid request body"})
		return
	}

	ticket, err := h.TicketCreateUseCase.Execute(createReq)
	if err != nil {
		c.JSON(400, domain.JSONResponse{Message: err.Error()})
		return
	}

	c.JSON(201, ticket)
}

// PurchaseTicket godoc
//	@Summary		Purchase a ticket
//	@Description	Purchases a ticket by ID
//	@Tags			tickets
//	@Accept			json
//	@Param			id			path	int						true	"Ticket ID"
//	@Param			purchase	body	domain.PurchaseRequest	true	"Purchase request data"
//	@Success		204
//	@Failure		400	{object}	domain.JSONResponse	"Invalid request body or ticket ID"
//	@Router			/tickets/{id}/purchases [post]
func (h *TicketHandler) PurchaseTicket(c *gin.Context) {
	ticketID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, domain.JSONResponse{Message: "Invalid ticket ID"})
		return
	}

	var purchaseReq domain.PurchaseRequest
	if err := c.ShouldBindJSON(&purchaseReq); err != nil {
		c.JSON(400, domain.JSONResponse{Message: "Invalid request body"})
		return
	}

	err = h.TicketPurchaseUseCase.Execute(ticketID, purchaseReq)
	if err != nil {
		c.JSON(400, domain.JSONResponse{Message: err.Error()})
		return
	}

	c.Status(204)
}
