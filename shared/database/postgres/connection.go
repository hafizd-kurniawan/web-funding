package postgres

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func SetupDB() *sqlx.DB {
	var err error
	// Example PostgreSQL connection string
	connStr := "user=root password=root dbname=funding sslmode=disable host=localhost port=5433"

	// sqlx.Connect opens a database and tests the connection with a Ping
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Configure connection pool settings (optional, but recommended)
	db.SetMaxOpenConns(25)                 // Set maximum number of open connections
	db.SetMaxIdleConns(5)                  // Set maximum number of idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // Set maximum lifetime of connections

	log.Println("Database connection established successfully!")
	return db
}
