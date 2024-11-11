package usecases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"api/internal/domain"
	"api/internal/repositories"
)

func TestTicketCreateUseCase(t *testing.T) {
	t.Run("should return error when name or description is empty", func(t *testing.T) {
		mockRepo := new(repositories.MockTicketRepository)
		useCase := NewTicketCreateUseCase(mockRepo)
		
		_, err := useCase.Execute(domain.CreateTicketRequest{})
		assert.Error(t, err)
	})

	t.Run("should return created ticket when valid data is provided", func(t *testing.T) {
		mockRepo := new(repositories.MockTicketRepository)
		
		mockRepo.On("CreateTicket", mock.AnythingOfType("domain.Ticket")).Return(domain.Ticket{
			ID:         1,
			Name:       "Concert A",
			Desc:       "Description A",
			Allocation: 100,
		}, nil)

		useCase := NewTicketCreateUseCase(mockRepo)
		ticket, err := useCase.Execute(domain.CreateTicketRequest{
			Name:       "Concert A",
			Desc:       "Description A",
			Allocation: 100,
		})

		assert.NoError(t, err)
		assert.Equal(t, "Concert A", ticket.Name)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails to create ticket", func(t *testing.T) {
		mockRepo := new(repositories.MockTicketRepository)

		mockRepo.On("CreateTicket", mock.AnythingOfType("domain.Ticket")).Return(domain.Ticket{}, errors.New("repository error"))

		useCase := NewTicketCreateUseCase(mockRepo)
		_, err := useCase.Execute(domain.CreateTicketRequest{
			Name:       "Concert B",
			Desc:       "Description B",
			Allocation: 100,
		})

		assert.Error(t, err, "Expected an error when repository fails to create ticket")
		mockRepo.AssertExpectations(t)
	})
}
