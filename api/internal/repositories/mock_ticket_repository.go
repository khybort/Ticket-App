package repositories

import (
    "api/internal/domain"
    "github.com/stretchr/testify/mock"
)

type MockTicketRepository struct {
    mock.Mock
}

func (m *MockTicketRepository) GetAllTickets() ([]domain.Ticket, error) {
    args := m.Called()
    return args.Get(0).([]domain.Ticket), args.Error(1)
}

func (m *MockTicketRepository) CreateTicket(ticket domain.Ticket) (domain.Ticket, error) {
    args := m.Called(ticket)
    return args.Get(0).(domain.Ticket), args.Error(1)
}

func (m *MockTicketRepository) GetTicket(ticketID int) (domain.Ticket, error) {
    args := m.Called(ticketID)
    return args.Get(0).(domain.Ticket), args.Error(1)
}

func (m *MockTicketRepository) UpdateTicket(ticket domain.Ticket) error {
    args := m.Called(ticket)
    return args.Error(0)
}
