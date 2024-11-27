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
	files := http.FileServer(http.Dir("./static"))
	p := os.Getenv("PORT")
	router := http.NewServeMux()
	routes.CreateRoutes(router)

	router.Handle("/static/", http.StripPrefix("/static/", files))

	s := http.Server{
		Addr:    fmt.Sprintf(":%s", p),
		Handler: router,
	}

	log.Printf("Server started and listening n port :%s ", p)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
