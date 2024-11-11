package usecases

import (
	"api/internal/domain"
	"api/internal/interfaces"
	"errors"
)

type TicketPurchaseUseCase interface {
	Execute(ticketID int, req domain.PurchaseRequest) error
}

type ticketPurchaseUseCase struct {
	ticketRepo interfaces.TicketRepository
}

func NewTicketPurchaseUseCase(ticketRepo interfaces.TicketRepository) TicketPurchaseUseCase {
	return &ticketPurchaseUseCase{
		ticketRepo: ticketRepo,
	}
}

func (uc *ticketPurchaseUseCase) Execute(ticketID int, req domain.PurchaseRequest) error {
	ticket, err := uc.ticketRepo.GetTicket(ticketID)
	if err != nil {
		return err
	}

	if ticket.Allocation < req.Quantity {
		return errors.New("not enough tickets available")
	}

	ticket.Allocation -= req.Quantity

	err = uc.ticketRepo.UpdateTicket(ticket)
	if err != nil {
		return err
	}

	return nil
}
