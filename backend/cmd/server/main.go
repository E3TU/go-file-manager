package main

import (
	"file-manager/internal/appwrite"
	"file-manager/internal/config"
	router "file-manager/internal/http"
	"file-manager/internal/http/handlers"
	"file-manager/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	client := appwrite.NewClient(cfg)

	authSvc := services.NewAuthService(client)
	storageSvc := services.NewStorageService(client, cfg.BucketID)

	h := handlers.NewHandler(authSvc, storageSvc)

	r := gin.Default()
	router.RegisterRoutes(r, h)

	log.Fatal(r.Run(":" + cfg.Port))

}
