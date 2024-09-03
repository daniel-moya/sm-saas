package database

import (
    "database/sql"
    "fmt"
    "log"
    "user-service/internal/config"
    _ "github.com/lib/pq"
)

func NewDatabase(cfg *config.Config) (*sql.DB, error) {
    dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
        cfg.DBUser,
        cfg.DBPassword,
        cfg.DBName,
        cfg.DBHost,
        cfg.DBPort,
    )

    log.Printf("dsn: %v \n", dsn)

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, fmt.Errorf("error opening database connection: %w", err)
    }

    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("error pinging database: %w", err)
    }

    log.Println("Successfully connected to the database")
    return db, nil
}
