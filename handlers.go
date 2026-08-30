package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AppHandler struct{}

func NewHandler() *AppHandler {
	return &AppHandler{}
}

func (a *AppHandler) HandlerGetIndex(forwardHTTPSUrl string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// convert to wss
		publicWSURL := strings.Replace(forwardHTTPSUrl, "https://", "wss://", 1)
		context := fiber.Map{
			"Public_URL":   forwardHTTPSUrl, // for user
			"Public_WSURL": publicWSURL,     // for request
		}
		return ctx.Render("index", context)
	}
}
