package routes

import (
	"bytes"
	"chat-app/golang-htmx/templates"
	"chat-app/golang-htmx/templates/components"
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type ErrJSON struct {
	Message string `json:"message"`
}

func joinChat(w http.ResponseWriter, r *http.Request) {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		component := components.Message("Hello Client!")
		buffer := &bytes.Buffer{}
		component.Render(context.Background(), buffer)

		for {
			// Write
			err := websocket.Message.Send(ws, buffer.Bytes())
			if err != nil {
				fmt.Printf("Error sending message: %s\n", err)
				return
			}

			time.Sleep(time.Second * 10)
			// Read
			// msg := ""
			// err = websocket.Message.Receive(ws, &msg)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// fmt.Printf("%s\n", msg)
		}
	}).ServeHTTP(w, r)
}

func CreateRoutes(s *http.ServeMux) {
	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := templates.Index()
		component.Render(context.Background(), w)
	})

	s.HandleFunc("/ws/chat", joinChat)
}
