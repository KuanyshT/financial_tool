package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/KuanyshT/financial-tool/backend/database"
	"github.com/KuanyshT/financial-tool/backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	godotenv.Load()
    database.InitDB()

	router := routes.NewRouter()

	// ✅ Оборачиваем маршруты в CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // или указать конкретные домены на проде
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}
