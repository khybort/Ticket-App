package interfaces

import "api/internal/domain"

type TicketRepository interface {
	GetAllTickets() ([]domain.Ticket, error)
	GetTicket(ticketID int) (domain.Ticket, error)
	CreateTicket(ticket domain.Ticket) (domain.Ticket, error)
	UpdateTicket(ticket domain.Ticket) error
}
