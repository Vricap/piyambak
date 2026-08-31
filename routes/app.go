package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vricap/ssshh/handlers"
)

func SetupAppRoutes(app *fiber.App) {
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		fmt.Println("ping")
		return ctx.SendString("Welcome to fiber\n")
	})
	app.Get("/", handlers.GetIndex)
	app.Get("/messages", handlers.GetAllMessages)
}
