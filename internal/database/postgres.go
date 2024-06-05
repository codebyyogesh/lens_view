package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib" // PostgreSQL driver
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

// Data Source Name:string that contains all the necessary information to
// connect to a database

func (cfg *PostgresConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Database, cfg.Password, cfg.SSLMode)
}

// default setup config from docker compose
func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "bagheera",
		Password: "junglebook",
		Database: "lensview",
		SSLMode:  "disable",
	}
}

// PS: When Open() is called, it is the responsibility of the caller
// to close the DB connection with db.Close()
func Open(cfg PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	return db, nil
}
