package internal

import (
	"bytes"
	"chat-app/golang-htmx/templates/components"
	"context"
	"fmt"
	"time"

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

var (
	pongWaitTime = time.Second * 10
	pingInterval = time.Second * 9
)

func NewClient(ws *websocket.Conn, manager *Manager) *Client {
	return &Client{
		Conn:           ws,
		ID:             uuid.New().String(),
		Chatroom:       "general",
		Manager:        manager,
		MessageChannel: make(chan string),
	}
}

func (c *Client) closeConnection() {
	c.Conn.Close()
	c.Manager.ClientListEventChannel <- &ClientListEvent{
		Client:    c,
		EventType: "REMOVE",
	}
}

func (c *Client) ReadMessages(ctx context.Context) {
	if err := c.Conn.SetReadDeadline(time.Now().Add(pongWaitTime)); err != nil {
		fmt.Printf("Error setting read deadline: %s\n", err)
		return
	}
	c.Conn.SetPongHandler(func(appData string) error {
		if err := c.Conn.SetReadDeadline(time.Now().Add(pongWaitTime)); err != nil {
			fmt.Printf("Error setting read deadline: %s\n", err)
			return err
		}
		fmt.Println("Pong received")
		return nil
	})
	defer c.closeConnection()

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

func (c *Client) WriteMessages(ctx context.Context) {
	defer c.closeConnection()

	ticker := time.NewTicker(pingInterval)
	for {
		select {
		case text, ok := <-c.MessageChannel:
			if !ok {
				return
			}

			component := components.Message(text)
			buffer := &bytes.Buffer{}
			component.Render(ctx, buffer)

			for _, client := range c.Manager.ClientList {
				err := client.Conn.WriteMessage(websocket.TextMessage, buffer.Bytes())
				if err != nil {
					fmt.Printf("Error sending message to client %s: %s\n", client.ID, err)
					continue
				}
			}

		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := c.Conn.WriteMessage(websocket.PingMessage, []byte("")); err != nil {
				fmt.Printf("Error sending ping message: %s\n", err)
				return
			}
			fmt.Println("Sent ping message")
		}
	}
}
