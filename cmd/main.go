package main

import (
	"chat-app/golang-htmx/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	p := os.Getenv("PORT")
	route := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static"))
	route.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	routes.CreateRoutes(route)

	s := http.Server{
		Addr:    fmt.Sprintf(":%s", p),
		Handler: route,
	}

	log.Printf("Server started and listening n port :%s ", p)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
