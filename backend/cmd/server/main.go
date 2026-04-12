package main

import (
	"file-manager/internal/appwrite"
	"file-manager/internal/config"
	"file-manager/internal/handlers"
	router "file-manager/internal/router"
	appwriteSvc "file-manager/internal/services/appwrite"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cfg := config.Load()
	log.Printf("Loaded endpoint: %s", cfg.AppwriteEndpoint)
	log.Printf("Loaded project: %s", cfg.AppwriteProjectId)
	log.Printf("Loaded api key: %s", cfg.AppwriteApiKey)
	log.Printf("PORT env: %s", os.Getenv("PORT"))

	client := appwrite.NewClient(cfg)

	authSvc := appwriteSvc.NewAuthService(client, cfg)

	h := handlers.NewHandler(authSvc)

	r := gin.Default()
	r.Use(cors.Default())
	router.RegisterRoutes(r, h)

	log.Fatal(r.Run(":" + cfg.Port))

}
