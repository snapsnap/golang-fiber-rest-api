package connection

import (
	"api-dev/internal/config"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetDatabase(conf config.Database) *sql.DB {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Port,
		conf.Name,
		conf.Tz,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err.Error())
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	// Test connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection test failed:", err.Error())
	}

	fmt.Println("Database connected successfully!")

	return db
}
