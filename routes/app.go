package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vricap/piyambak/handlers"
)

func SetupAppRoutes(app *fiber.App) {
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		fmt.Println("ping")
		return ctx.SendString("Welcome to fiber\n")
	})
	app.Get("/", handlers.GetAllRooms)
	app.Post("/rooms/:id", handlers.GetRoom) // uses POST because user also send the room password in body
	app.Post("/rooms", handlers.CreateRoom)
	app.Get("/messages/:id", handlers.GetMessages)
	// app.Post("/rooms", handlers.CreateRoom)
}
