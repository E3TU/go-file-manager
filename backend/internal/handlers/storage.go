package handlers

import (
	"errors"
	"io"
	"net/http"

	"file-manager/internal/services/appwrite"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UploadFile(c *gin.Context) {
	cookieName := "a_session"
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	session, err := h.authService.GetSession(cookie)
	if err != nil || !session.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read file"})
		return
	}

	bucketID := h.storageService.GetConfig().BucketID
	resp, err := h.storageService.UploadFile(bucketID, header.Filename, content, session.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp.DownloadURL = h.storageService.GetDownloadURL(bucketID, resp.ID)

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) ListFiles(c *gin.Context) {
	cookieName := "a_session"
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	session, err := h.authService.GetSession(cookie)
	if err != nil || !session.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
		return
	}

	bucketID := h.storageService.GetConfig().BucketID
	files, err := h.storageService.ListFiles(bucketID, session.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i := range files {
		files[i].DownloadURL = h.storageService.GetDownloadURL(bucketID, files[i].ID)
	}

	c.JSON(http.StatusOK, gin.H{"files": files})
}

func (h *Handler) DeleteFile(c *gin.Context) {
	cookieName := "a_session"
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	session, err := h.authService.GetSession(cookie)
	if err != nil || !session.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
		return
	}

	fileID := c.Param("id")
	if fileID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File ID is required"})
		return
	}

	bucketID := h.storageService.GetConfig().BucketID
	err = h.storageService.DeleteFile(bucketID, fileID, session.UserID)
	if err != nil {
		if errors.Is(err, appwrite.ErrFileNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
			return
		}
		if errors.Is(err, appwrite.ErrPermissionDenied) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File deleted"})
}

func (h *Handler) DownloadFile(c *gin.Context) {
	cookieName := "a_session"
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	session, err := h.authService.GetSession(cookie)
	if err != nil || !session.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
		return
	}

	fileID := c.Param("id")
	if fileID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File ID is required"})
		return
	}

	bucketID := h.storageService.GetConfig().BucketID
	downloadURL, err := h.storageService.GetFileDownloadURL(bucketID, fileID, session.UserID)
	if err != nil {
		if errors.Is(err, appwrite.ErrFileNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
			return
		}
		if errors.Is(err, appwrite.ErrPermissionDenied) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, downloadURL)
}