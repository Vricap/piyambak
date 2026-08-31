package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/vricap/ssshh/handlers"
)

func SetupWebsocketRoutes(app *fiber.App) {
	// create new websocket
	server := handlers.NewWebSocket()
	app.Get("/ws", websocket.New(func(ctx *websocket.Conn) {
		server.HandleWebSocket(ctx)
	}))

	go server.HandleMessage()
}
