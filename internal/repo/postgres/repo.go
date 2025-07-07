package pgrepo

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/abdullahnettoor/tastybites/internal/config"
	"github.com/abdullahnettoor/tastybites/internal/repo/interfaces"
	_ "github.com/jackc/pgx/v5/stdlib" 
)

type repository struct {
	DB *sql.DB
}

func NewRepository(cfg *config.DBConfig) (interfaces.Repository, error) {
	// Initialize the database connection here
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
	log.Println("Connecting to database with connection string:", connStr)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}
	return &repository{DB: db}, nil
}
