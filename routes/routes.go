package routes

import (
	templates "chat-app/golang-htmx/internal/templates"
	"context"
	"net/http"
)

type ErrJSON struct {
	Message string `json:"message"`
}

func CreateRoutes(s *http.ServeMux) {
	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := templates.Index()
		component.Render(context.Background(), w)
	})
}
