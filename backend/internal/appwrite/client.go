package appwrite

import (
	"file-manager/internal/config"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
)

func NewClient(cfg config.Config) client.Client {

	clt := appwrite.NewClient(
		appwrite.WithEndpoint(cfg.AppwriteEndpoint),
		appwrite.WithKey(cfg.AppwriteApiKey),
	)

	return clt
}
