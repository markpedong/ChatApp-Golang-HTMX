package routes

import (
	"chat-app/golang-htmx/internal"
	"chat-app/golang-htmx/templates"
	"net/http"
)

func CreateRoutes(s *http.ServeMux) {
	manager := internal.NewManager()

	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := templates.Index()
		component.Render(r.Context(), w)
	})

	s.HandleFunc("/ws/chat", manager.Handle)
}
