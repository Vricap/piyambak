package main

import "github.com/gofiber/fiber/v2"

type AppHandler struct{}

func NewHandler() *AppHandler {
	return &AppHandler{}
}

func (a *AppHandler) HandlerGetIndex(ctx *fiber.Ctx) error {
	context := fiber.Map{}

	return ctx.Render("index", context)
}
