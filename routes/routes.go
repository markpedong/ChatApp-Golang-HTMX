package routes

import (
	"chat-app/golang-htmx/internal"
	"chat-app/golang-htmx/templates"
	"context"
	"net/http"
)

func CreateRoutes(s *http.ServeMux) {
	files := http.FileServer(http.Dir("./static"))
	manager := internal.NewManager()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go manager.HandleClientListEventChannel(ctx)

	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := templates.Index()
		component.Render(r.Context(), w)
	})

	s.HandleFunc("/ws/chat", manager.Handle)

	s.Handle("/static/", http.StripPrefix("/static/", files))
}
