package appwrite

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"file-manager/internal/config"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/file"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/permission"
	"github.com/appwrite/sdk-for-go/role"
	"github.com/appwrite/sdk-for-go/storage"
)

var (
	ErrFileNotFound     = errors.New("file not found")
	ErrPermissionDenied = errors.New("permission denied")
)

type StorageService struct {
	cfg config.Config
}

func NewStorageService(cfg config.Config) *StorageService {
	return &StorageService{
		cfg: cfg,
	}
}

func (s *StorageService) GetConfig() config.Config {
	return s.cfg
}

func (s *StorageService) newAdminClient() client.Client {
	return appwrite.NewClient(
		appwrite.WithEndpoint(s.cfg.AppwriteEndpoint),
		appwrite.WithProject(s.cfg.AppwriteProjectId),
		appwrite.WithKey(s.cfg.AppwriteApiKey),
	)
}

func (s *StorageService) hasReadPermission(permissions []string, userID string) bool {
	target := fmt.Sprintf(`"user:%s"`, userID)
	for _, p := range permissions {
		if strings.Contains(p, target) && strings.Contains(p, "read(") {
			return true
		}
	}
	return false
}

func (s *StorageService) hasWritePermission(permissions []string, userID string) bool {
	target := fmt.Sprintf(`"user:%s"`, userID)
	for _, p := range permissions {
		if strings.Contains(p, target) && (strings.Contains(p, "write(") || strings.Contains(p, "delete(")) {
			return true
		}
	}
	return false
}

type FileResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	SizeOriginal int    `json:"sizeOriginal,omitempty"`
	MimeType     string `json:"mimeType"`
	CreatedAt    string `json:"createdAt"`
	DownloadURL  string `json:"downloadUrl,omitempty"`
}

func (s *StorageService) UploadFile(bucketID string, fileName string, fileContent []byte, userID string) (*FileResponse, error) {

	tmpFile, err := os.CreateTemp("", "upload-*")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(fileContent); err != nil {
		tmpFile.Close()
		return nil, err
	}
	tmpFile.Close()

	adminClient := s.newAdminClient()
	svc := storage.New(adminClient)

	resp, err := svc.CreateFile(
		bucketID,
		id.Unique(),
		file.NewInputFile(tmpFile.Name(), fileName),
		svc.WithCreateFilePermissions([]string{
			permission.Read(role.User(userID, "")),
			permission.Write(role.User(userID, "")),
		}),
	)
	if err != nil {
		return nil, err
	}

	return &FileResponse{
		ID:           resp.Id,
		Name:         resp.Name,
		SizeOriginal: resp.SizeOriginal,
		MimeType:     resp.MimeType,
		CreatedAt:    resp.CreatedAt,
	}, nil
}

func (s *StorageService) ListFiles(bucketID string, userID string) ([]FileResponse, error) {
	adminClient := s.newAdminClient()
	svc := storage.New(adminClient)

	resp, err := svc.ListFiles(bucketID)
	if err != nil {
		return nil, err
	}

	files := make([]FileResponse, 0, len(resp.Files))
	for _, f := range resp.Files {
		if !s.hasReadPermission(f.Permissions, userID) {
			continue
		}
		files = append(files, FileResponse{
			ID:           f.Id,
			Name:         f.Name,
			SizeOriginal: f.SizeOriginal,
			MimeType:     f.MimeType,
			CreatedAt:    f.CreatedAt,
		})
	}

	return files, nil
}

func (s *StorageService) GetFile(bucketID string, fileID string) (*FileResponse, []string, error) {
	adminClient := s.newAdminClient()
	svc := storage.New(adminClient)

	resp, err := svc.GetFile(bucketID, fileID)
	if err != nil {
		var appErr *client.AppwriteError
		if errors.As(err, &appErr) && appErr.GetStatusCode() == http.StatusNotFound {
			return nil, nil, ErrFileNotFound
		}
		return nil, nil, err
	}

	return &FileResponse{
		ID:           resp.Id,
		Name:         resp.Name,
		SizeOriginal: resp.SizeOriginal,
		MimeType:     resp.MimeType,
		CreatedAt:    resp.CreatedAt,
	}, resp.Permissions, nil
}

func (s *StorageService) DeleteFile(bucketID string, fileID string, userID string) error {
	_, permissions, err := s.GetFile(bucketID, fileID)
	if err != nil {
		return err
	}

	if !s.hasWritePermission(permissions, userID) {
		return ErrPermissionDenied
	}

	adminClient := s.newAdminClient()
	svc := storage.New(adminClient)

	_, err = svc.DeleteFile(bucketID, fileID)
	return err
}

func (s *StorageService) DownloadFile(bucketID, fileID, userID, localPath string) error {
	_, permissions, err := s.GetFile(bucketID, fileID)
	if err != nil {
		return err
	}

	if !s.hasReadPermission(permissions, userID) {
		return ErrPermissionDenied
	}

	adminClient := s.newAdminClient()
	svc := storage.New(adminClient)

	data, err := svc.GetFileDownload(bucketID, fileID)
	if err != nil {
		return err
	}

	outFile, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = outFile.Write(*data)
	return err
}

func (s *StorageService) GetFileBytes(bucketID, fileID, userID string) ([]byte, error) {
	_, permissions, err := s.GetFile(bucketID, fileID)
	if err != nil {
		return nil, err
	}

	if !s.hasReadPermission(permissions, userID) {
		return nil, ErrPermissionDenied
	}

	adminClient := s.newAdminClient()
	svc := storage.New(adminClient)

	data, err := svc.GetFileDownload(bucketID, fileID)
	if err != nil {
		return nil, err
	}

	return *data, nil
}
