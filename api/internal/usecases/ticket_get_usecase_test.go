package usecases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"api/internal/domain"
	"api/internal/repositories"
)

func TestTicketGetUseCase(t *testing.T) {
	t.Run("should return ticket when found", func(t *testing.T) {
		mockRepo := new(repositories.MockTicketRepository)

		mockRepo.On("GetTicket", 1).Return(domain.Ticket{
			ID:         1,
			Name:       "Concert A",
			Desc:       "Description A",
			Allocation: 100,
		}, nil)

		useCase := NewTicketGetUseCase(mockRepo)

		ticket, err := useCase.Execute(1)

		assert.NoError(t, err)
		assert.Equal(t, 1, ticket.ID)
		assert.Equal(t, "Concert A", ticket.Name)
		assert.Equal(t, "Description A", ticket.Desc)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when ticket not found", func(t *testing.T) {
		mockRepo := new(repositories.MockTicketRepository)

		mockRepo.On("GetTicket", 1).Return(domain.Ticket{}, errors.New("ticket not found"))

		useCase := NewTicketGetUseCase(mockRepo)

		ticket, err := useCase.Execute(1)

		assert.Error(t, err, "Expected an error when ticket not found")
		assert.Equal(t, "ticket not found", err.Error())
		assert.Equal(t, domain.Ticket{}, ticket)

		mockRepo.AssertExpectations(t)
	})
}
