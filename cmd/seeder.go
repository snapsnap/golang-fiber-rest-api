package cmd

import (
	"database/sql"
	"fmt"
	"log"
)

// RunSeeder menjalankan seeder untuk data awal
func RunSeeder(db *sql.DB) {
	fmt.Println("Running seeder...")

	query := `INSERT INTO users (name, email) VALUES 
		('Admin', 'admin@example.com'),
		('User1', 'user1@example.com')`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Seeder failed: %v", err)
	}

	fmt.Println("Seeder completed successfully!")
}
