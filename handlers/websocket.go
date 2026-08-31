package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/websocket/v2"
	"github.com/vricap/ssshh/database"
	"github.com/vricap/ssshh/models"
)

type WebSocketServer struct {
	clients   map[*websocket.Conn]bool
	broadcast chan *models.Message
}

func NewWebSocket() *WebSocketServer {
	return &WebSocketServer{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan *models.Message),
	}
}

func (s *WebSocketServer) HandleWebSocket(ctx *websocket.Conn) {
	// register new client
	s.clients[ctx] = true // FIX: RACE CONDITION ISSUE. HandleMessage could write to the map when at the same time we write new user to the map. Use MUTEX
	defer func() {
		delete(s.clients, ctx)
		ctx.Close()
	}()

	for {
		_, msg, err := ctx.ReadMessage()
		if err != nil {
			log.Println("Read Error:", err)
			break
		}

		// send the message to the broadcast channel
		var message *models.Message
		err = json.Unmarshal(msg, &message)
		if err != nil {
			log.Fatal("Error Unmarshalling")
		}
		s.broadcast <- message

		// fill the created_at field and store to DB
		message.CreatedAt = time.Now().UTC().Format("2006-01-02 15:04:05")
		err = storeMessageToDB(database.DbCtx.DB, message)
		if err != nil {
			log.Fatalf("Failed to store message to database: %v", err)
		}
	}
}

func (s *WebSocketServer) HandleMessage() {
	for {
		msg := <-s.broadcast

		// send the message to all Clients
		for client := range s.clients {
			msg, err := marshalMessage(msg)
			if err != nil {
				log.Printf("Error marshaling message: %v", err)
			}
			err = client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Printf("Write Error: %v", err)
				client.Close()
				delete(s.clients, client)
			}
		}
	}
}

func marshalMessage(msg *models.Message) ([]byte, error) {
	s, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func storeMessageToDB(DB *sql.DB, message *models.Message) error {
	query := `
	INSERT INTO messages (message, user) VALUES (?, ?)`

	_, err := DB.Exec(query, message.Text, message.User)
	if err != nil {
		return err
	}

	return nil
}
