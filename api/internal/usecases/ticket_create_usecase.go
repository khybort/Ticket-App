package usecases

import (
	"api/internal/domain"
	"api/internal/interfaces"
	"errors"
)

type TicketCreateUseCase interface {
	Execute(req domain.TicketRequest) (domain.Ticket, error)
}

type ticketCreateUseCase struct {
	ticketRepo interfaces.TicketRepository
}

func NewTicketCreateUseCase(ticketRepo interfaces.TicketRepository) TicketCreateUseCase {
	return &ticketCreateUseCase{
		ticketRepo: ticketRepo,
	}
}

func (uc *ticketCreateUseCase) Execute(req domain.TicketRequest) (domain.Ticket, error) {
	if req.Name == "" || req.Desc == "" {
		return domain.Ticket{}, errors.New("ticket name and description are required")
	}

	ticket := domain.Ticket{
		Name:       req.Name,
		Desc:       req.Desc,
		Allocation: req.Allocation,
	}

	ticketDTO, err := uc.ticketRepo.CreateTicket(ticket)
	if err != nil {
		return domain.Ticket{}, err
	}

	return ticketDTO, nil
}
