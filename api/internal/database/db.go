package database

import (
    "database/sql"
    "fmt"
    "log"
    "api/internal/config"

    _ "github.com/lib/pq"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
    connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName,
    )

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, fmt.Errorf("could not open database connection: %w", err)
    }

    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("could not establish a connection to the database: %w", err)
    }

    log.Println("Connected to the database successfully.")
    return db, nil
}
