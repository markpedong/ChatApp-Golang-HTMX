package main

import (
	"chat-app/golang-htmx/helper"
	"chat-app/golang-htmx/routes"
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	router := http.NewServeMux()

	stack := helper.CreateStack(
		helper.Logging,
	)

	routes.CreateRoutes(router)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: stack(router),
	}

	fmt.Printf("Listening on port %s\n", os.Getenv("PORT"))
	srv.ListenAndServe()
}
