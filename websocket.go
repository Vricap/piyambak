package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"

	"github.com/gofiber/websocket/v2"
)

type WebSocketServer struct {
	clients   map[*websocket.Conn]bool
	broadcast chan *Message
}

func NewWebSocket() *WebSocketServer {
	return &WebSocketServer{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan *Message),
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
		var message Message
		err = json.Unmarshal(msg, &message)
		if err != nil {
			log.Fatal("Error Unmarshalling")
		}
		s.broadcast <- &message
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

func marshalMessage(msg *Message) ([]byte, error) {
	s, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func getMessageTemplate(msg *Message) []byte {
	base := `
<div id="messages">
    <p class="text-small">{{ .Text }}</p>
</div>
`

	tmpl, err := template.New("msg").Parse(base)
	if err != nil {
		log.Fatalf("template parsing: %s", err)
	}

	// render the template with the message as data
	var renderedMessage bytes.Buffer
	err = tmpl.Execute(&renderedMessage, msg)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

	return renderedMessage.Bytes()
}
