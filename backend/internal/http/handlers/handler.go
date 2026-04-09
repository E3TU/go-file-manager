// internal/http/handlers/handler.go
package handlers

import "file-manager/internal/services"

type Handler struct {
	auth    *services.AuthService
	storage *services.StorageService
}

func NewHandler(a *services.AuthService, s *services.StorageService) *Handler {
	return &Handler{auth: a, storage: s}
}
