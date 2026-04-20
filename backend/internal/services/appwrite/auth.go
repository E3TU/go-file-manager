package appwrite

import (
	"log"

	"file-manager/internal/config"

	"github.com/appwrite/sdk-for-go/account"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/users"
)

// AuthService handles user authentication with Appwrite
type AuthService struct {
	cfg config.Config
}

// NewAuthService creates a new AuthService instance
func NewAuthService(cfg config.Config) *AuthService {
	return &AuthService{
		cfg: cfg,
	}
}

// Request/Response DTOs
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	UserID    string `json:"userId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}

type CreateSessionRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateSessionResponse struct {
	UserID        string `json:"userId"`
	SessionID     string `json:"sessionId"`
	SessionSecret string `json:"sessionSecret"`
}

type GetSessionResponse struct {
	UserID    string `json:"userId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Valid     bool   `json:"valid"`
	CreatedAt string `json:"createdAt"`
}

// -------------------------
// Client helpers
// -------------------------

func (s *AuthService) newAdminClient() client.Client {
	return appwrite.NewClient(
		appwrite.WithEndpoint(s.cfg.AppwriteEndpoint),
		appwrite.WithProject(s.cfg.AppwriteProjectId),
		appwrite.WithKey(s.cfg.AppwriteApiKey),
	)
}

func (s *AuthService) newSessionClient(sessionSecret string) client.Client {
	return appwrite.NewClient(
		appwrite.WithEndpoint(s.cfg.AppwriteEndpoint),
		appwrite.WithProject(s.cfg.AppwriteProjectId),
		appwrite.WithSession(sessionSecret),
	)
}

// -------------------------
// Auth methods
// -------------------------

func (s *AuthService) Register(req RegisterRequest) (*RegisterResponse, error) {
	adminClient := s.newAdminClient()
	service := users.New(adminClient)

	user, err := service.Create(
		id.Unique(),
		service.WithCreateName(req.Name),
		service.WithCreateEmail(req.Email),
		service.WithCreatePassword(req.Password),
	)
	if err != nil {
		return nil, err
	}

	return &RegisterResponse{
		UserID:    user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *AuthService) CreateSession(req CreateSessionRequest) (*CreateSessionResponse, error) {
	adminClient := s.newAdminClient()
	service := account.New(adminClient)

	session, err := service.CreateEmailPasswordSession(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &CreateSessionResponse{
		UserID:        session.UserId,
		SessionID:     session.Id,
		SessionSecret: session.Secret,
	}, nil
}

func (s *AuthService) GetSession(sessionSecret string) (*GetSessionResponse, error) {
	if sessionSecret == "" {
		return &GetSessionResponse{Valid: false}, nil
	}

	sessionClient := s.newSessionClient(sessionSecret)
	service := account.New(sessionClient)

	user, err := service.Get()
	if err != nil {
		log.Printf("GetSession error: %v", err)
		return &GetSessionResponse{Valid: false}, nil
	}

	return &GetSessionResponse{
		UserID:    user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Valid:     true,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *AuthService) DeleteSession(sessionId string, sessionSecret string) error {
	sessionClient := s.newSessionClient(sessionSecret)
	service := account.New(sessionClient)

	_, err := service.DeleteSession(sessionId)
	return err
}
