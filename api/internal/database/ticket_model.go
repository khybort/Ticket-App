package database

type Ticket struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Desc      string `gorm:"not null"`
	Allocation int    `gorm:"default:100"`
}