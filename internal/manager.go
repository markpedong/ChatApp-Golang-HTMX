package internal

import (
	"bytes"
	"chat-app/golang-htmx/templates/components"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Manager struct {
	ClientList []*Client
}

var (
	upgrader = websocket.Upgrader{}
)

func NewManager() *Manager {
	return &Manager{
		ClientList: []*Client{},
	}
}

func (m *Manager) Handle(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	newClient := NewClient(ws, m)

	component := components.Message("Hello Client!")
	buffer := &bytes.Buffer{}
	component.Render(r.Context(), buffer)
	for {
		err := ws.WriteMessage(websocket.TextMessage, buffer.Bytes())
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return

		}

		time.Sleep(time.Second * 10)

		// Read
		// _, msg, err := ws.ReadMessage()
		// if err != nil {
		// 	c.AbortWithError(http.StatusInternalServerError, err)
		// 	return
		// }
		// fmt.Printf("%s\n", msg)
	}
}
