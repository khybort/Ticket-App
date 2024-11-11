package repositories

import (
	"api/internal/domain"
	"database/sql"
	"fmt"
	"log"
)

type TicketRepository struct {
	DB *sql.DB
}

func NewTicketRepo(db *sql.DB) *TicketRepository {
	return &TicketRepository{DB: db}
}

func (r *TicketRepository) GetAllTickets() ([]domain.Ticket, error) {
	rows, err := r.DB.Query(`SELECT id, name, "desc", allocation FROM tickets`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []domain.Ticket
	for rows.Next() {
		var ticket domain.Ticket
		if err := rows.Scan(&ticket.ID, &ticket.Name, &ticket.Desc, &ticket.Allocation); err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}
	fmt.Print("Ticket: ", tickets)

	return tickets, nil
}

func (r *TicketRepository) GetTicket(id int) (domain.Ticket, error) {
	query := `SELECT id, name, "desc", allocation FROM tickets WHERE id = $1`
	row := r.DB.QueryRow(query, id)

	var ticket domain.Ticket

	err := row.Scan(&ticket.ID, &ticket.Name, &ticket.Desc, &ticket.Allocation)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Ticket{}, fmt.Errorf("ticket with id %d not found", id)
		}
		log.Printf("Error scanning ticket: %v", err)
		return domain.Ticket{}, err
	}
	return ticket, nil
}

func (r *TicketRepository) CreateTicket(ticket domain.Ticket) (domain.Ticket, error) {
	query := `
        INSERT INTO tickets (name, "desc", allocation)
        VALUES ($1, $2, $3) RETURNING id
    `

	err := r.DB.QueryRow(query, ticket.Name, ticket.Desc, ticket.Allocation).Scan(&ticket.ID)
	if err != nil {
		log.Printf("Error inserting ticket: %v", err)
		return domain.Ticket{}, err
	}

	return ticket, nil
}

func (r *TicketRepository) UpdateTicket(ticket domain.Ticket) error {
	query := `
        UPDATE tickets
        SET name = $1, "desc" = $2, allocation = $3
        WHERE id = $4
    `

	res, err := r.DB.Exec(query, ticket.Name, ticket.Desc, ticket.Allocation, ticket.ID)
	if err != nil {
		log.Printf("Error updating ticket: %v", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("ticket with id %d not found", ticket.ID)
	}

	return nil
}
