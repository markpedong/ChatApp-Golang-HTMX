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
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}
		fmt.Printf("Received message: %s\n", msg) // Log received message
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
			fmt.Printf("Sending message: %s\n", text) // Log sent message
			component := components.Message(text)
			buffer := &bytes.Buffer{}
			component.Render(r.Context(), buffer)

			for _, client := range c.Manager.ClientList {
				err := client.Conn.WriteMessage(websocket.TextMessage, buffer.Bytes())
				if err != nil {
					fmt.Printf("Error sending message: %s\n", err)
					return
				}
			}

		case <-r.Context().Done():
			return
		}

	}

}
