package cmd

import (
	"api-dev/internal/utils"
	"database/sql"
	"fmt"
	"log"
	"time"
)

// RunSeeder menjalankan seeder untuk data awal
func RunSeeder(db *sql.DB) {
	fmt.Println("Running seeder...")
	SeedUser(db)
	fmt.Println("Seeder completed successfully!")
}

// SeedUser inserts initial data into the users table
func SeedUser(db *sql.DB) {

	users := []struct {
		name     string
		email    string
		password string
	}{
		{"John Doe", "john@example.com", "password123"},
		{"Jane Doe", "jane@example.com", "securepass"},
		{"Alice", "alice@example.com", "mypassword"},
	}

	for _, user := range users {
		hashedPassword, err := utils.HashPassword(user.password)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}

		// Format waktu ke dalam format timestamp SQL
		createdAt := time.Now().Format("2006-01-02 15:04:05")

		query := `INSERT INTO users (name, email, password, created_at) VALUES (?, ?, ?, ?)`
		_, err = db.Exec(query, user.name, user.email, hashedPassword, createdAt)
		if err != nil {
			log.Fatalf("Seeding failed: %v", err)
		}
	}
}
