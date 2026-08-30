package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/websocket/v2"
	// "github.com/vricap/ssshh/handlers"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Static("/static", "./static")

	app.Get("/ping", func(ctx *fiber.Ctx) error {
		fmt.Println("ping")
		return ctx.SendString("Welcome to fiber\n")
	})

	appHandler := NewHandler()

	app.Get("/", appHandler.HandlerGetIndex)

	// create new websocket
	server := NewWebSocket()
	app.Get("/ws", websocket.New(func(ctx *websocket.Conn) {
		server.HandleWebSocket(ctx)
	}))

	go server.HandleMessage()

	err := startNgrok()
	if err != nil {
		log.Fatalf("Error starting ngrok: %v", err)
	}

	app.Listen(":3000")
}

func startNgrok() error {
	cmd := exec.Command("ngrok", "http", "3000")
	// cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	err := cmd.Start()
	if err != nil {
		return err
	}

	return nil
}
