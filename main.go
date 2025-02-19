package main

import (
	"api-dev/cmd"
	"api-dev/internal/api"
	"api-dev/internal/config"
	"api-dev/internal/connection"
	"api-dev/internal/repositories"
	"api-dev/internal/services"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)
	db := goqu.New("mysql", dbConnection)
	defer dbConnection.Close()

	app := fiber.New()

	// Menjalankan Migration atau Seeder berdasarkan argumen CLI
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			cmd.RunMigrations(dbConnection)
			fmt.Println("Migration completed!")
			return
		case "seed":
			cmd.RunSeeder(dbConnection)
			fmt.Println("Seeding completed!")
			return
		default:
			fmt.Println("Command not recognized")
			return
		}
	}

	// Buat folder logs jika belum ada
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0755)
		if err != nil {
			log.Fatalf("Error creating logs folder: %v", err)
		}
	}

	// Logging ke file per hari
	logFileName := fmt.Sprintf("logs/log_%s.log", time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer file.Close()

	// Middleware logger
	app.Use(logger.New(logger.Config{
		Output: file,
		Format: "[${time}] ${status} - ${method} ${path} | ${latency}\n",
	}))
	// Tambahkan log manual ke file
	loggerTest := log.New(file, "", log.LstdFlags)
	loggerTest.Println("Server is starting...")

	// Daftarkan repository
	userRepository := repositories.NewUser(db)

	// Daftarkan service
	userService := services.NewUser(userRepository)

	// Load API
	api.NewUser(app, userService)

	loggerTest.Println("Fiber server is running on : " + cnf.Server.Port)
	fiberlog.Info(app.Listen(cnf.Server.Host + ":" + cnf.Server.Port))

}
