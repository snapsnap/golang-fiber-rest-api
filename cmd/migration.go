package cmd

import (
	"database/sql"
	"fmt"
	"log"
)

// RunMigrations menjalankan semua migration
func RunMigrations(db *sql.DB) {
	fmt.Println("Running migrations...")

	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT NULL,
			deleted_at TIMESTAMP DEFAULT NULL
		)`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
	}

	fmt.Println("Migrations completed successfully!")
}
