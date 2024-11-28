package main

import (
	"chat-app/golang-htmx/templates"
	"context"
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	files := http.FileServer(http.Dir("./static"))
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := templates.Index()
		component.Render(context.Background(), w)
	})

	router.Handle("/static/", http.StripPrefix("/static/", files))

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: router,
	}

	fmt.Printf("Listening on port %s\n", os.Getenv("PORT"))
	srv.ListenAndServe()
}
