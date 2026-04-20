package main

import (
	"file-manager/internal/config"
	"file-manager/internal/handlers"
	router "file-manager/internal/router"
	appwriteSvc "file-manager/internal/services/appwrite"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cfg := config.Load()

	authSvc := appwriteSvc.NewAuthService(cfg)
	storageSvc := appwriteSvc.NewStorageService(cfg)

	h := handlers.NewHandler(authSvc, storageSvc)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:8080"},
		AllowMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:  []string{"Content-Type"},
	}))
	router.RegisterRoutes(r, h)

	log.Fatal(r.Run(":" + cfg.Port))

}
