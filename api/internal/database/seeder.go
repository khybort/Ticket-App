package database

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	_ "gorm.io/driver/postgres"
)

func SeedDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(&Ticket{}); err != nil {
		return fmt.Errorf("failed to migrate the tickets table: %w", err)
	}

	var count int64
	if err := db.Model(&Ticket{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to check tickets count: %w", err)
	}

	if count == 0 {
		log.Println("No tickets found, inserting initial seed data.")
		tickets := []Ticket{
			{Name: "Concert A", Desc: "Desc A", Allocation: 100},
			{Name: "Concert B", Desc: "Desc B", Allocation: 100},
			{Name: "Concert C", Desc: "Desc C", Allocation: 100},
		}
		if err := db.Create(&tickets).Error; err != nil {
			return fmt.Errorf("failed to seed tickets: %w", err)
		}
		log.Println("Seed data inserted successfully.")
	} else {
		log.Println("Tickets table already contains data. Skipping seed.")
	}

	return nil
}
