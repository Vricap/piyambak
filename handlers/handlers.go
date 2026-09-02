package handlers

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/vricap/piyambak/database"
	"github.com/vricap/piyambak/models"
	"github.com/vricap/piyambak/utils"
)

func GetAllRooms(ctx *fiber.Ctx) error {
	rows, err := database.DbCtx.DB.Query(`SELECT * FROM rooms ORDER BY created_at ASC;`)
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
		return err
	}
	defer rows.Close()

	var rooms []models.Room

	for rows.Next() {
		var room models.Room
		err := rows.Scan(&room.ID, &room.Name, &room.Password, &room.CreatedAt)
		if err != nil {
			log.Printf("Error scanning rows: %v", err)
			return ctx.Status(500).SendString("Error getting rooms.")
		}
		room.CreatedAt, err = utils.FormatTimestamp(room.CreatedAt)
		rooms = append(rooms, room)
	}

	if rows.Err() != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": rows.Err(),
		})
	}

	context := fiber.Map{
		"Rooms": rooms,
	}
	return ctx.Render("index", context)
}

func GetRoom(ctx *fiber.Ctx) error {
	roomID := ctx.Params("id")
	inputPassword := ctx.FormValue("password")

	var dbPassword string

	err := database.DbCtx.DB.QueryRow(
		`SELECT password FROM rooms WHERE id = ?`,
		roomID,
	).Scan(&dbPassword)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Room not found")
	}

	if inputPassword != dbPassword {
		return ctx.Status(fiber.StatusUnauthorized).SendString("Incorrect password")
	}

	// convert to wss
	forwardHTTPSUrl := utils.NgroxCtx.Url
	publicWSURL := strings.Replace(forwardHTTPSUrl, "https://", "wss://", 1)
	context := fiber.Map{
		"Public_URL":   forwardHTTPSUrl, // for user
		"Public_WSURL": publicWSURL,     // for request
		"RoomID":       roomID,
	}
	return ctx.Render("chatRoom", context)
}

func CreateRoom(ctx *fiber.Ctx) error {
	var room *models.Room
	err := ctx.BodyParser(&room)
	if err != nil {
		log.Fatalf("Error parsing body request: %v", err)
		return err
	}

	_, err = database.DbCtx.DB.Exec(`INSERT INTO rooms (name, password) VALUES (?, ?);`, room.Name, room.Password)
	if err != nil {
		log.Printf("Error inserting into database: %v", err)
		return ctx.Status(500).SendString("Error creating room.")
	}

	return ctx.Redirect("/", fiber.StatusSeeOther)
}

func GetMessages(ctx *fiber.Ctx) error {
	roomId := ctx.Params("id")
	rows, err := database.DbCtx.DB.Query(`SELECT * FROM messages WHERE  room_id = ? ORDER BY created_at ASC;`, roomId)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		return ctx.Status(500).SendString("Error getting messages.")
	}
	defer rows.Close()

	var messages []models.Message

	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.ID, &msg.RoomId, &msg.User, &msg.Text, &msg.CreatedAt)
		if err != nil {
			log.Printf("Error scanning rows: %v", err)
			return ctx.Status(500).SendString("Error getting rooms.")
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
