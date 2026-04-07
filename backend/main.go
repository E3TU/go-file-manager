package main

import (
	"log"
	"os"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("no .env loaded: %v", err)
	}

	client := appwrite.NewClient(
		appwrite.WithEndpoint(os.Getenv("APPWRITE_ENDPOINT")),
		appwrite.WithProject(os.Getenv("APPWRITE_PROJECT_ID")),
		appwrite.WithKey(os.Getenv("APPWRITE_API_KEY")),
	)

	databases := appwrite.NewDatabases(client)

	res, err := databases.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, db := range res.Databases {
		log.Printf("id=%s name=%s", db.Id, db.Name)
	}
}
