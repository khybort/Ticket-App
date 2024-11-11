package usecases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"api/internal/domain"
	"api/internal/repositories"
)

func TestTicketPurchaseUseCase(t *testing.T) {
	t.Run("should successfully purchase ticket when enough allocation is available", func(t *testing.T) {
		mockRepo := new(repositories.MockTicketRepository)

		mockRepo.On("GetTicket", 1).Return(domain.Ticket{
			ID:         1,
			Name:       "Concert A",
			Desc:       "Description A",
			Allocation: 100,
		}, nil)

		mockRepo.On("UpdateTicket", domain.Ticket{
			ID:         1,
			Name:       "Concert A",
			Desc:       "Description A",
			Allocation: 95,
		}).Return(nil)

		useCase := NewTicketPurchaseUseCase(mockRepo)

		err := useCase.Execute(1, domain.PurchaseRequest{Quantity: 5})

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when not enough tickets are available", func(t *testing.T) {
		mockRepo := new(repositories.MockTicketRepository)

		mockRepo.On("GetTicket", 1).Return(domain.Ticket{
			ID:         1,
			Name:       "Concert B",
			Desc:       "Description B",
			Allocation: 3,
		}, nil)

		useCase := NewTicketPurchaseUseCase(mockRepo)
		err := useCase.Execute(1, domain.PurchaseRequest{Quantity: 5})

		assert.Error(t, err)
		assert.Equal(t, "not enough tickets available", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when GetTicket fails", func(t *testing.T) {
		mockRepo := new(repositories.MockTicketRepository)

		mockRepo.On("GetTicket", 1).Return(domain.Ticket{}, errors.New("ticket not found"))

		useCase := NewTicketPurchaseUseCase(mockRepo)
		err := useCase.Execute(1, domain.PurchaseRequest{Quantity: 2})

		assert.Error(t, err)
		assert.Equal(t, "ticket not found", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when UpdateTicket fails", func(t *testing.T) {
		mockRepo := new(repositories.MockTicketRepository)

		mockRepo.On("GetTicket", 1).Return(domain.Ticket{
			ID:         1,
			Name:       "Concert A",
			Desc:       "Description A",
			Allocation: 50,
		}, nil)

		mockRepo.On("UpdateTicket", domain.Ticket{
			ID:         1,
			Name:       "Concert A",
			Desc:       "Description A",
			Allocation: 48,
		}).Return(errors.New("update failed"))

		useCase := NewTicketPurchaseUseCase(mockRepo)
		err := useCase.Execute(1, domain.PurchaseRequest{Quantity: 2})

		assert.Error(t, err)
		assert.Equal(t, "update failed", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
