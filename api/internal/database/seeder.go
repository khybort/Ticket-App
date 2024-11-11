package database

import (
	"database/sql"
	"fmt"
	"log"
)

func SeedDatabase(db *sql.DB) error {
	// Create the tickets table if it doesn't exist
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS tickets (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			"desc" VARCHAR(255) NOT NULL,
			allocation INTEGER NOT NULL DEFAULT 100
		);
	`

	// Check if the table exists and contains data
	var count int
	// Use a safer approach to check if the table exists and has data
	row := db.QueryRow(`
		SELECT COUNT(*) FROM information_schema.tables
		WHERE table_schema = 'public' AND table_name = 'tickets'
	`)
	var tableExists bool
	var err = row.Scan(&tableExists)
	if err != nil {
		return fmt.Errorf("failed to check tickets table existence: %w", err)
	}

	if !tableExists {
		log.Println("Tickets table does not exist. Creating it.")
		// If the table doesn't exist, create it
		_, err = db.Exec(createTableQuery)
		if err != nil {
			return fmt.Errorf("failed to create tickets table: %w", err)
		}
	}

	// Check for data in the table
	row = db.QueryRow("SELECT COUNT(*) FROM tickets")
	if err := row.Scan(&count); err != nil {
		return fmt.Errorf("failed to check tickets count: %w", err)
	}

	// If no tickets are found, insert seed data
	if count == 0 {
		log.Println("No tickets found, inserting initial seed data.")
		seedDataQuery := `
            INSERT INTO tickets (name, "desc", allocation) VALUES
            ('Concert A', 'Desc A', 100),
            ('Concert B', 'Desc B', 100),
            ('Concert C', 'Desc C', 100);
        `
		_, err = db.Exec(seedDataQuery)
		if err != nil {
			return fmt.Errorf("failed to seed tickets: %w", err)
		}
		log.Println("Seed data inserted successfully.")
	} else {
		log.Println("Tickets table already contains data. Skipping seed.")
	}

	return nil
}
