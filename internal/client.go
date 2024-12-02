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

			// Render the message component with the received text
			component := components.Message(text)
			buffer := &bytes.Buffer{}
			err := component.Render(r.Context(), buffer)
			if err != nil {
				fmt.Printf("Error rendering component: %v\n", err)
				continue
			}

			// Send the rendered message to the WebSocket
			for _, client := range c.Manager.ClientList {
				messageHTML := buffer.String() // Get the HTML content to send

				// Send the rendered message to the client
				err := client.Conn.WriteMessage(websocket.TextMessage, []byte(messageHTML))
				if err != nil {
					fmt.Printf("Error sending message to client %s: %s\n", client.ID, err)
					client.Conn.Close() // Cleanup
					continue
				}
				fmt.Printf("Message sent to client %s: %s\n", client.ID, messageHTML) // Debug log
			}

		case <-r.Context().Done():
			return
		}
	}
}
