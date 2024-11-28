package main

import (
	"chat-app/golang-htmx/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	routes.CreateRoutes(router)

	fmt.Printf("Listening on port %s\n", os.Getenv("PORT"))
	router.Run(":" + os.Getenv("PORT"))
}
