package internal

import (
	"bytes"
	"chat-app/golang-htmx/templates/components"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn           *websocket.Conn
	ID             string
	Chatroom       string
	Manager        *Manager
	MessageChannel chan string
}

func NewClient(ws *websocket.Conn, manager *Manager) *Client {
	return &Client{
		Conn:           ws,
		ID:             uuid.New().String(),
		Chatroom:       "general",
		Manager:        manager,
		MessageChannel: make(chan string),
	}
}

func (c *Client) ReadMessages(r *http.Request) {
	defer func() {
		c.Conn.Close()
		c.Manager.ClientListEventChannel <- &ClientListEvent{
			Client:    c,
			EventType: "REMOVE",
		}
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Printf("Error reading message from WebSocket: %s\n", err)
			return
		}
		fmt.Printf("Received message: %s\n", msg)
		c.MessageChannel <- string(msg)
	}
}

func (c *Client) WriteMessages(r *http.Request) {
	defer func() {
		c.Conn.Close()
		c.Manager.ClientListEventChannel <- &ClientListEvent{
			Client:    c,
			EventType: "REMOVE",
		}
	}()

	for {
		select {
		case text, ok := <-c.MessageChannel:
			if !ok {
				return
			}

			component := components.Message(text)
			buffer := &bytes.Buffer{}
			err := component.Render(r.Context(), buffer)
			if err != nil {
				fmt.Printf("Error rendering component: %v\n", err)
				continue
			}

			for _, client := range c.Manager.ClientList {
				message := fmt.Sprintf("Message from client %s: %s", c.ID, text)
				fmt.Printf("Broadcasting: %s\n", message) // Debug log

				err := client.Conn.WriteMessage(websocket.TextMessage, []byte(message))
				if err != nil {
					fmt.Printf("Error sending message to client %s: %s\n", client.ID, err)
					client.Conn.Close() // Cleanup
					continue
				}
			}

		case <-r.Context().Done():
			return
		}

	}

}
