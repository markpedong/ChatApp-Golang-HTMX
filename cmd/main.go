package main

import (
	"chat-app/golang-htmx/routes"
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	files := http.FileServer(http.Dir("./static"))
	router := http.NewServeMux()

	router.Handle("/static/", http.StripPrefix("/static/", files))
	routes.CreateRoutes(router)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: router,
	}

	fmt.Printf("Listening on port %s\n", os.Getenv("PORT"))
	srv.ListenAndServe()
}
