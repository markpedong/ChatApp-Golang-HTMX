package routes

import (
	"chat-app/golang-htmx/controllers.go"
	"chat-app/golang-htmx/internal"
	"chat-app/golang-htmx/templates"
	"context"
	"net/http"
)

func CreateRoutes(s *http.ServeMux, manager *internal.Manager, ctx context.Context) {
	files := http.FileServer(http.Dir("./static"))

	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := templates.Index()
		component.Render(r.Context(), w)
	})

	s.HandleFunc("/ws/chat", manager.Handle)
	s.HandleFunc("/components", controllers.HandleComponent)

	s.Handle("/static/", http.StripPrefix("/static/", files))
}
