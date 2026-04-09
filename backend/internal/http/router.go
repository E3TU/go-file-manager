package router

import (
	"file-manager/internal/http/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *handlers.Handler) {
	api := r.Group("/api")

	api.POST("/auth/session", h.CreateSession)
	api.GET("/auth/session", h.GetSession)
	api.DELETE("/auth/session", h.DeleteSession)

	api.POST("/storage/files", h.UploadFile)
	api.GET("/storage/files", h.ListFiles)
	api.DELETE("/storage/files/:id", h.DeleteFile)
}
