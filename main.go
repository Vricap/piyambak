package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vricap/piyambak/database"
	"github.com/vricap/piyambak/routes"
	"github.com/vricap/piyambak/utils"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins: fmt.Sprintf("http://localhost:3000, %s", utils.NgroxCtx.Url),
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))
	app.Static("/static", "./web/static")
	routes.SetupRoutes(app)

	app.Listen(PORT)
}
