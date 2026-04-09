// internal/services/storage.go
package services

import (
	"github.com/appwrite/sdk-for-go/client"
)

type StorageService struct {
	client   client.Client
	bucketID string
}

func NewStorageService(c client.Client, bucketID string) *StorageService {
	return &StorageService{client: c, bucketID: bucketID}
}
