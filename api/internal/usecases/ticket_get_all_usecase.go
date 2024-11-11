package usecases

import (
	"api/internal/domain"
	"api/internal/interfaces"
)

type TicketGetAllUseCase interface {
	Execute() ([]domain.Ticket, error)
}

type ticketGetAllUseCase struct {
	ticketRepo interfaces.TicketRepository
}

func NewTicketGetAllUseCase(ticketRepo interfaces.TicketRepository) TicketGetAllUseCase {
	return &ticketGetAllUseCase{
		ticketRepo: ticketRepo,
	}
}

func (uc *ticketGetAllUseCase) Execute() ([]domain.Ticket, error) {
	return uc.ticketRepo.GetAllTickets()
}
