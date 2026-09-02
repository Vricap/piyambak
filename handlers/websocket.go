package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/websocket/v2"
	"github.com/vricap/piyambak/database"
	"github.com/vricap/piyambak/models"
	"github.com/vricap/piyambak/utils"
)

type Room struct {
	Clients   map[*websocket.Conn]bool
	Broadcast chan *models.Message
}

type WebSocketServer struct {
	Rooms map[string]*Room
}

func NewWebSocket() *WebSocketServer {
	return &WebSocketServer{
		Rooms: make(map[string]*Room),
		// Broadcast: make(chan *models.Message),
	}
}

func (s *WebSocketServer) JoinRoom(roomID string, ctx *websocket.Conn) {
	if s.Rooms[roomID] == nil { // not initializing already exist room
		s.Rooms[roomID] = &Room{
			Clients:   make(map[*websocket.Conn]bool),
			Broadcast: make(chan *models.Message),
		}
	}
	// register new client to room
	s.Rooms[roomID].Clients[ctx] = true // FIX: RACE CONDITION ISSUE. HandleMessage could write to the map when at the same time we write new user to the map. Use MUTEX
	defer func() {
		delete(s.Rooms[roomID].Clients, ctx)
		ctx.Close()
	}()

	for {
		_, msg, err := ctx.ReadMessage()
		if err != nil {
			log.Println("Read Error:", err)
			break
		}
		// send the message to the Broadcast channel
		var message *models.Message
		err = json.Unmarshal(msg, &message)
		if err != nil {
			log.Println("Error Unmarshalling: %v", err)
			break
		}

		s.Rooms[roomID].Broadcast <- message

		// fill the created_at field and store to DB
		message.CreatedAt = time.Now().UTC().Format("2006-01-02 15:04:05")
		err = storeMessageToDB(database.DbCtx.DB, message)
		if err != nil {
			log.Printf("Failed to store message to database: %v\n", err)
			break
		}
	}
}

func (s *WebSocketServer) Broadcast(roomID string) {
	for {
		msg := <-s.Rooms[roomID].Broadcast
		msg.CreatedAt, _ = utils.FormatTimestamp(msg.CreatedAt)

		// send the message to all Clients
		for client := range s.Rooms[roomID].Clients {
			msg, err := MarshalMessage(msg)
			if err != nil {
				log.Printf("Error marshaling message: %v", err)
			}
			err = client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Printf("Write Error: %v", err)
				client.Close()
				delete(s.Rooms[roomID].Clients, client)
			}
		}
	}
}

func MarshalMessage(msg *models.Message) ([]byte, error) {
	s, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func storeMessageToDB(DB *sql.DB, message *models.Message) error {
	query := `
	INSERT INTO messages (message, user, room_id) VALUES (?, ?, ?)`

	_, err := DB.Exec(query, message.Text, message.User, message.RoomId)
	if err != nil {
		return err
	}

	return nil
}
