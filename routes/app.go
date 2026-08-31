package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vricap/ssshh/handlers"
)

func SetupAppRoutes(app *fiber.App) {
	appHandler := handlers.NewHandler()

	app.Get("/ping", func(ctx *fiber.Ctx) error {
		fmt.Println("ping")
		return ctx.SendString("Welcome to fiber\n")
	})
	app.Get("/", appHandler.HandlerGetIndex)
}
