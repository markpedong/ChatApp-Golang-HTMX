package routes

import (
	"bytes"
	"chat-app/golang-htmx/templates"
	"chat-app/golang-htmx/templates/components"
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
)

func joinChat(c *gin.Context) {
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

		time.Sleep(time.Second * 1)

		// Read
		// _, msg, err := ws.ReadMessage()
		// if err != nil {
		// 	c.AbortWithError(http.StatusInternalServerError, err)
		// 	return
		// }
		// fmt.Printf("%s\n", msg)
	}

}

func CreateRoutes(c *gin.Engine) {
	c.GET("/", func(ctx *gin.Context) {
		component := templates.Index()
		component.Render(context.Background(), ctx.Writer)
	})

	c.GET("/ws/chat", joinChat)
}
