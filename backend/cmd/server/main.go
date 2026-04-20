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

	// client := appwrite.NewClient(cfg)

	authSvc := appwriteSvc.NewAuthService(cfg)

	h := handlers.NewHandler(authSvc)

	r := gin.Default()
	r.Use(cors.Default())
	router.RegisterRoutes(r, h)

	log.Fatal(r.Run(":" + cfg.Port))

}
