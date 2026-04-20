package appwrite

import (
	"fmt"
	"os"

	"file-manager/internal/config"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/file"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/storage"
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

type FileResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	SizeOriginal int    `json:"sizeOriginal,omitempty"`
	MimeType     string `json:"mimeType"`
	CreatedAt    string `json:"createdAt"`
	DownloadURL  string `json:"downloadUrl,omitempty"`
}

func (s *StorageService) UploadFile(bucketID string, fileName string, fileContent []byte) (*FileResponse, error) {

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

func (s *StorageService) ListFiles(bucketID string) ([]FileResponse, error) {
	adminClient := s.newAdminClient()
	svc := storage.New(adminClient)

	resp, err := svc.ListFiles(bucketID)
	if err != nil {
		return nil, err
	}

	files := make([]FileResponse, 0, len(resp.Files))
	for _, f := range resp.Files {
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

func (s *StorageService) GetFile(bucketID string, fileID string) (*FileResponse, error) {
	adminClient := s.newAdminClient()
	svc := storage.New(adminClient)

	resp, err := svc.GetFile(bucketID, fileID)
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

func (s *StorageService) DeleteFile(bucketID string, fileID string) error {
	adminClient := s.newAdminClient()
	svc := storage.New(adminClient)

	_, err := svc.DeleteFile(bucketID, fileID)
	return err
}

func (s *StorageService) GetDownloadURL(bucketID string, fileID string) string {
	return fmt.Sprintf("%s/v1/storage/buckets/%s/files/%s/download",
		s.cfg.AppwriteEndpoint, bucketID, fileID)
}
