package usecases

import (
	"api/internal/domain"
	"api/internal/interfaces"
)

type TicketGetUseCase interface {
	Execute(ticketID int) (domain.Ticket, error)
}

type ticketGetUseCase struct {
	ticketRepo interfaces.TicketRepository
}

func NewTicketGetUseCase(ticketRepo interfaces.TicketRepository) TicketGetUseCase {
	return &ticketGetUseCase{
		ticketRepo: ticketRepo,
	}
}

func (uc *ticketGetUseCase) Execute(ticketID int) (domain.Ticket, error) {
	ticket, err := uc.ticketRepo.GetTicket(ticketID)
	if err != nil {
		return domain.Ticket{}, err
	}

	return ticket, nil
}
