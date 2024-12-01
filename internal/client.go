package internal

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	ID       string
	Chatroom string
	Manager  *Manager
}

func NewClient(ws *websocket.Conn, manager *Manager) *Client {
	return &Client{
		conn:     ws,
		ID:       uuid.New().String(),
		Chatroom: "general",
		Manager:  manager,
	}
}

func (c *Client) ReadMessage() {}

func (c *Client) WriteMessage() {}
