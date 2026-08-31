package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/vricap/ssshh/utils"
)

type AppHandler struct{}

func NewHandler() *AppHandler {
	return &AppHandler{}
}

func (a *AppHandler) HandlerGetIndex(ctx *fiber.Ctx) error {
	// convert to wss

	forwardHTTPSUrl := utils.NgroxCtx.Url
	publicWSURL := strings.Replace(forwardHTTPSUrl, "https://", "wss://", 1)
	context := fiber.Map{
		"Public_URL":   forwardHTTPSUrl, // for user
		"Public_WSURL": publicWSURL,     // for request
	}
	return ctx.Render("index", context)
}
