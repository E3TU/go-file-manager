package handlers

import (
	"file-manager/internal/services/appwrite"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService *appwrite.AuthService
}

func NewHandler(authSvc *appwrite.AuthService) *Handler {
	return &Handler{
		authService: authSvc,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req appwrite.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *Handler) CreateSession(c *gin.Context) {
	var req appwrite.CreateSessionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := h.authService.CreateSession(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	cookieName := "a_session"
	c.SetCookie(
		cookieName,
		session.SessionSecret,
		0,
		"/",
		"",
		true,
		true,
	)

	c.JSON(http.StatusOK, session)
}

func (h *Handler) GetSession(c *gin.Context) {
	cookieName := "a_session"
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No session cookie"})
		return
	}

	session, err := h.authService.GetSession(cookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if !session.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
		return
	}

	c.JSON(http.StatusOK, session)
}

func (h *Handler) DeleteSession(c *gin.Context) {
	var req struct {
		SessionID string `json:"sessionId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cookieName := "a_session"
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No session cookie"})
		return
	}

	if err := h.authService.DeleteSession(req.SessionID, cookie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie(
		cookieName,
		"",
		-1,
		"/",
		"",
		true,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"message": "Session deleted"})
}
