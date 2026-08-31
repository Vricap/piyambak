package handlers

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/vricap/ssshh/database"
	"github.com/vricap/ssshh/models"
	"github.com/vricap/ssshh/utils"
)

func GetIndex(ctx *fiber.Ctx) error {
	// convert to wss
	forwardHTTPSUrl := utils.NgroxCtx.Url
	publicWSURL := strings.Replace(forwardHTTPSUrl, "https://", "wss://", 1)
	context := fiber.Map{
		"Public_URL":   forwardHTTPSUrl, // for user
		"Public_WSURL": publicWSURL,     // for request
	}
	return ctx.Render("index", context)
}

func GetAllMessages(ctx *fiber.Ctx) error {
	rows, err := database.DbCtx.DB.Query(`SELECT * FROM messages ORDER BY created_at ASC;`)
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
		return err
	}
	defer rows.Close()

	var messages []models.Message

	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.ID, &msg.User, &msg.Text, &msg.CreatedAt)
		if err != nil {
			log.Fatalf("Error scanning rows: %v", err)
			return err
		}
		msg.CreatedAt, err = utils.FormatTimestamp(msg.CreatedAt)
		messages = append(messages, msg)
	}

	if rows.Err() != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": rows.Err(),
		})
	}

	return ctx.JSON(messages)
}
