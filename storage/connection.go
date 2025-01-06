package storage

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Package-level variable to hold the MinIO client
var minioClient *minio.Client

// Connection initializes the MinIO client
func Connection() {
	// MinIO server details
	endpoint := "161.97.141.69:9100"        // Replace with your MinIO endpoint
	accessKeyID := "chetanbudathoki" // Replace with your MinIO root username
	secretAccessKey := "Online248" // Replace with your MinIO root password
	useSSL := false                     // Use true if MinIO is configured with SSL

	// Initialize MinIO client
	var err error
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("Error initializing MinIO client: %v", err)
	}

	log.Println("Connected to the storage 'youthcongressnepal'")
	
	CreateBucket()	
}

// GetClient retrieves the initialized MinIO client
func GetClient() *minio.Client {
	if minioClient == nil {
		log.Fatalf("MinIO client is not initialized. Call storage.Connection() first.")
	}
	return minioClient
}
