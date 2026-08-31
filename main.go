package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/vricap/ssshh/database"
	"github.com/vricap/ssshh/routes"
	"github.com/vricap/ssshh/utils"

	_ "github.com/mattn/go-sqlite3"
)

const HOST = "http://localhost"
const PORT = ":3000"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	DB, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to load database: %v", err)
	}

	err = database.RunMigrations(DB)
	if err != nil {
		log.Fatalf("Error running database migrations: %v", err)
	}
	defer DB.Close()

	// run ngrok
	err = utils.StartNgrok(HOST, PORT)
	if err != nil {
		log.Fatalf("Error running ngrok: %v", err)
	}

	// setup gofiber
	app := fiber.New(fiber.Config{
		Views: html.New("./web/views", ".html"),
	})
	app.Static("/static", "./web/static")
	routes.SetupRoutes(app)

	app.Listen(PORT)
}
