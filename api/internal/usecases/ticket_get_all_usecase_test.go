package usecases

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"api/internal/domain"
	"api/internal/repositories"
)

func TestTicketGetAllUseCase(t *testing.T) {
	mockRepo := new(repositories.MockTicketRepository)

	mockRepo.On("GetAllTickets").Return([]domain.Ticket{
		{ID: 1, Name: "Concert A", Desc: "Description A", Allocation: 100},
		{ID: 2, Name: "Concert B", Desc: "Description B", Allocation: 200},
	}, nil)

	useCase := NewTicketGetAllUseCase(mockRepo)

	tickets, err := useCase.Execute()
	assert.NoError(t, err)

	assert.Equal(t, 2, len(tickets))
	assert.Equal(t, "Concert A", tickets[0].Name)

	mockRepo.AssertExpectations(t)
}
