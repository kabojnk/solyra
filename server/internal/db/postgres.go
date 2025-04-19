package db

import (
	"database/sql"
	"fmt"
	"github.com/kabojnk/solyra/server/internal/config"
	"github.com/kabojnk/solyra/server/internal/models"
	_ "github.com/lib/pq"
)

// PostgresDB represents a PostgreSQL database connection
type PostgresDB struct {
	db *sql.DB
}

// NewPostgresDB creates a new PostgreSQL database connection
func NewPostgresDB(cfg config.DatabaseConfig) (*PostgresDB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{db: db}, nil
}

// Close closes the database connection
func (p *PostgresDB) Close() error {
	return p.db.Close()
}

// GetApplicationByClientID retrieves an application by its client ID
func (p *PostgresDB) GetApplicationByClientID(clientID string) (*models.Application, error) {
	query := `SELECT id, client_id, client_secret FROM applications WHERE client_id = $1`

	var app models.Application
	err := p.db.QueryRow(query, clientID).Scan(&app.ID, &app.ClientID, &app.ClientSecret)
	if err != nil {
		return nil, err
	}

	return &app, nil
}

// CreateApplication creates a new application
func (p *PostgresDB) CreateApplication(app *models.Application) error {
	query := `INSERT INTO applications (id, client_id, client_secret) VALUES ($1, $2, $3)`

	_, err := p.db.Exec(query, app.ID, app.ClientID, app.ClientSecret)
	return err
}

// DeleteApplication deletes an application by its ID
func (p *PostgresDB) DeleteApplication(id string) error {
	query := `DELETE FROM applications WHERE id = $1`

	_, err := p.db.Exec(query, id)
	return err
}
