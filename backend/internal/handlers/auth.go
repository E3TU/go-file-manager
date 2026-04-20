package handlers

import (
	"file-manager/internal/services/appwrite"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService     *appwrite.AuthService
	storageService  *appwrite.StorageService
}

func NewHandler(authSvc *appwrite.AuthService, storageSvc *appwrite.StorageService) *Handler {
	return &Handler{
		authService:    authSvc,
		storageService: storageSvc,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req appwrite.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1️⃣ Create the user
	user, err := h.authService.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2️⃣ Automatically create a session (login)
	session, err := h.authService.CreateSession(appwrite.CreateSessionRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session after registration"})
		return
	}

	// 3️⃣ Set the cookie exactly like the login flow
	cookieName := "a_session"
	isProd := false
	if c.Request.TLS != nil {
		isProd = true
	}

	c.SetCookie(
		cookieName,
		session.SessionSecret,
		3600, // 1 hour
		"/",
		"localhost",
		isProd, // Secure only in prod
		true,   // HttpOnly
	)

	// 4️⃣ Return the user + session info if needed
	c.JSON(http.StatusCreated, gin.H{
		"user":    user,
		"session": session,
	})
}

func (h *Handler) CreateSession(c *gin.Context) {
	var req appwrite.CreateSessionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Appwrite session
	session, err := h.authService.CreateSession(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Set cookie for session secret
	cookieName := "a_session"

	// Detect if we are in production (HTTPS) or local dev
	isProd := false
	if c.Request.TLS != nil {
		isProd = true
	}

	c.SetCookie(
		cookieName,
		session.SessionSecret,
		3600, // 1 hour expiry
		"/",
		"localhost", // domain (change in prod)
		isProd,      // Secure in prod
		true,        // HttpOnly
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
	if err != nil || !session.Valid {
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
