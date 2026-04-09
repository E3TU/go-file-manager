package config

import "os"

type Config struct {
	AppwriteEndpoint  string
	AppwriteProjectId string
	AppwriteApiKey    string
	BucketID          string
	Port              string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return Config{
		AppwriteEndpoint:  os.Getenv("APPWRITE_ENDPOINT"),
		AppwriteProjectId: os.Getenv("APPWRITE_PROJECT_ID"),
		AppwriteApiKey:    os.Getenv("APPWRITE_API_KEY"),
		BucketID:          os.Getenv("APPWRITE_BUCKET_ID"),
		Port:              port,
	}
}
