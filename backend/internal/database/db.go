package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// DB wraps the database connection
type DB struct {
	*sql.DB
}

// Connect creates a new database connection
func Connect() (*DB, error) {
	connStr := os.Getenv("SUPABASE_CONNECTION_STRING")
	if connStr == "" {
		// Build connection string from individual parts
		host := os.Getenv("SUPABASE_HOST")
		port := os.Getenv("SUPABASE_PORT")
		user := os.Getenv("SUPABASE_USER")
		password := os.Getenv("SUPABASE_PASSWORD")
		dbname := os.Getenv("SUPABASE_DB")

		if host == "" {
			return nil, fmt.Errorf("database connection info not provided")
		}

		if port == "" {
			port = "5432"
		}

		connStr = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
			host, port, user, password, dbname,
		)
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	return &DB{db}, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.DB.Close()
}
