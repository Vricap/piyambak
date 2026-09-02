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
	app.Get("/", handlers.Index)
	app.Get("/rooms", handlers.GetAllRooms)

	// uses POST because user also send the room password in body
	// but if user refresh the page, will return Method Not Allowed since we use POST
	// we must separate the responbilities of checking password with POST, and returning the page with GET. we must use session or some kind
	app.Post("/rooms/:id/join", handlers.GetRoom)
	app.Post("/rooms", handlers.CreateRoom)
	app.Get("/messages/:id", handlers.GetMessages)
}
