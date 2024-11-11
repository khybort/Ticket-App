package repositories

import (
	"api/internal/domain"
	"api/internal/database"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type TicketRepository struct {
	DB *gorm.DB
}

func NewTicketRepo(db *gorm.DB) *TicketRepository {
	return &TicketRepository{DB: db}
}

func (r *TicketRepository) GetAllTickets() ([]domain.Ticket, error) {
	var ticketModels []database.Ticket
	if err := r.DB.Find(&ticketModels).Error; err != nil {
		return nil, err
	}

	var tickets []domain.Ticket
	for _, ticketModel := range ticketModels {
		tickets = append(tickets, domain.Ticket{
			ID:         ticketModel.ID,
			Name:       ticketModel.Name,
			Desc:       ticketModel.Desc,
			Allocation: ticketModel.Allocation,
		})
	}

	log.Print("Tickets: ", tickets)
	return tickets, nil
}

func (r *TicketRepository) GetTicket(id int) (domain.Ticket, error) {
	var ticketModel database.Ticket
	if err := r.DB.First(&ticketModel, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Ticket{}, fmt.Errorf("ticket with id %d not found", id)
		}
		log.Printf("Error retrieving ticket: %v", err)
		return domain.Ticket{}, err
	}
	ticket := domain.Ticket{
		ID:         ticketModel.ID,
		Name:       ticketModel.Name,
		Desc:       ticketModel.Desc,
		Allocation: ticketModel.Allocation,
	}
	return ticket, nil
}

func (r *TicketRepository) CreateTicket(ticket domain.Ticket) (domain.Ticket, error) {
	ticketModel := database.Ticket{
		Name:       ticket.Name,
		Desc:       ticket.Desc,
		Allocation: ticket.Allocation,
	}

	if err := r.DB.Create(&ticketModel).Error; err != nil {
		log.Printf("Error creating ticket: %v", err)
		return domain.Ticket{}, err
	}

	ticket.ID = ticketModel.ID
	return ticket, nil
}

func (r *TicketRepository) UpdateTicket(ticket domain.Ticket) error {
	var ticketModel database.Ticket
	if err := r.DB.First(&ticketModel, ticket.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("ticket with id %d not found", ticket.ID)
		}
		log.Printf("Error finding ticket: %v", err)
		return err
	}

	ticketModel.Name = ticket.Name
	ticketModel.Desc = ticket.Desc
	ticketModel.Allocation = ticket.Allocation

	if err := r.DB.Save(&ticketModel).Error; err != nil {
		log.Printf("Error updating ticket: %v", err)
		return err
	}

	return nil
}
