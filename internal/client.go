package internal

import "github.com/gorilla/websocket"

type Client struct {
	conn     *websocket.Conn
	ID       string
	Chatroom string
	Manager  *Manager
}
