package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/joho/godotenv"
	"golang.ngrok.com/ngrok/v2"
	// "github.com/vricap/ssshh/handlers"
)

const HOST = "http://localhost"
const PORT = ":3000"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	// run ngrok
	ctx := context.Background()
	forwardHTTPSUrl, err := runNgrok(ctx)
	if err != nil {
		log.Fatalf("Error running ngrok: %v", err)
	}

	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Static("/static", "./static")

	app.Get("/ping", func(ctx *fiber.Ctx) error {
		fmt.Println("ping")
		return ctx.SendString("Welcome to fiber\n")
	})

	appHandler := NewHandler()

	app.Get("/", appHandler.HandlerGetIndex(forwardHTTPSUrl))

	// create new websocket
	server := NewWebSocket()
	app.Get("/ws", websocket.New(func(ctx *websocket.Conn) {
		server.HandleWebSocket(ctx)
	}))

	go server.HandleMessage()

	app.Listen(PORT)
}

func runNgrok(ctx context.Context) (string, error) {
	agent, err := ngrok.NewAgent(ngrok.WithAuthtoken(os.Getenv("NGROK_AUTHTOKEN")))
	if err != nil {
		return "", err
	}

	ln, err := agent.Forward(ctx,
		ngrok.WithUpstream(HOST+PORT),
	)

	if err != nil {
		return "", err
	}

	fmt.Println("Endpoint online: forwarding from", ln.URL(), "to", HOST+PORT)
	return ln.URL().String(), nil
}
