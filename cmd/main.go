package main

import (
	"chat-app/golang-htmx/helper"
	"chat-app/golang-htmx/internal"
	"chat-app/golang-htmx/routes"
	"context"
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := http.NewServeMux()
	m := internal.NewManager()

	c, cancel := context.WithCancel(context.Background())
	defer cancel()
	go m.HandleClientListEventChannel(c)

	stack := helper.CreateStack(
		helper.Logging,
	)
	routes.CreateRoutes(r, m, c)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: stack(r),
	}

	fmt.Printf("Listening on port %s\n", os.Getenv("PORT"))
	srv.ListenAndServe()
}
