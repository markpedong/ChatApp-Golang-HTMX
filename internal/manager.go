package internal

import (
	"bytes"
	"chat-app/golang-htmx/templates/components"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
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

func (m *Manager) Handle(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	component := components.Message("Hello Client!")
	buffer := &bytes.Buffer{}
	component.Render(c, buffer)
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
