package storage

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
)

// Default bucket name
const defaultBucketName = "youthcongressnepal"

// CreateBucket ensures the specified bucket exists or creates the default bucket if none is provided
func CreateBucket(bucketName ...string) {
	client := GetClient() // Ensure the MinIO client is initialized
	ctx := context.Background()

	// Use default bucket name if no bucket name is provided
	name := defaultBucketName
	if len(bucketName) > 0 {
		name = bucketName[0]
	}

	// Check if the bucket exists
	exists, err := client.BucketExists(ctx, name)
	if err != nil {
		log.Fatalf("Error checking if bucket exists: %v", err)
	}
	if exists {
		log.Printf("Bucket '%s' already exists", name)
		return
	}

	// Create the bucket
	err = client.MakeBucket(ctx, name, minio.MakeBucketOptions{})
	if err != nil {
		log.Fatalf("Error creating bucket: %v", err)
	}
	log.Printf("Bucket %s created successfully", name)
}
