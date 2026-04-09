// internal/services/auth.go
package services

import (
	"github.com/appwrite/sdk-for-go/client"
)

type AuthService struct {
	client client.Client
}

func NewAuthService(c client.Client) *AuthService {
	return &AuthService{client: c}
}
