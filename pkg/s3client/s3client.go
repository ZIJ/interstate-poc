package s3client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

// S3Client wraps the S3 client and provides methods for interacting with S3
type S3Client struct {
	client *s3.Client
	bucket string
}

// NewS3Client creates a new S3Client
func NewS3Client() (*S3Client, error) {
	bucket := os.Getenv("S3_BUCKET")
	if bucket == "" {
		return nil, fmt.Errorf("S3_BUCKET environment variable is not set")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS SDK config: %w", err)
	}

	client := s3.NewFromConfig(cfg)

	return &S3Client{
		client: client,
		bucket: bucket,
	}, nil
}

// CreateFolder creates a new "folder" in S3 (which is actually just an empty object with a key ending in '/')
func (s *S3Client) CreateFolder(key string) error {
	_, err := s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &s.bucket,
		Key:    &key,
	})
	if err != nil {
		return fmt.Errorf("failed to create folder in S3: %w", err)
	}
	return nil
}

// ListFolders returns a list of all "folders" in the S3 bucket
func (s *S3Client) ListFolders() ([]string, error) {
	resp, err := s.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &s.bucket,
		Delimiter: aws.String("/"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list folders in S3: %w", err)
	}

	var folders []string
	for _, prefix := range resp.CommonPrefixes {
		folders = append(folders, *prefix.Prefix)
	}
	return folders, nil
}

// DeleteFolder deletes a "folder" and all its contents from S3
func (s *S3Client) DeleteFolder(key string) error {
	// List all objects in the folder
	resp, err := s.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &s.bucket,
		Prefix: &key,
	})
	if err != nil {
		return fmt.Errorf("failed to list objects in folder: %w", err)
	}

	// Delete all objects in the folder
	for _, obj := range resp.Contents {
		_, err := s.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
			Bucket: &s.bucket,
			Key:    obj.Key,
		})
		if err != nil {
			return fmt.Errorf("failed to delete object %s: %w", *obj.Key, err)
		}
	}

	return nil
}

// PutObject puts an object into the S3 bucket
func (s *S3Client) PutObject(key string, data []byte) error {
	_, err := s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &s.bucket,
		Key:    &key,
		Body:   bytes.NewReader(data),
	})
	if err != nil {
		return fmt.Errorf("failed to put object in S3: %w", err)
	}
	return nil
}

// GetObject retrieves an object from the S3 bucket
func (s *S3Client) GetObject(key string) ([]byte, error) {
	resp, err := s.client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &s.bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get object from S3: %w", err)
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// DeleteObject deletes an object from the S3 bucket
func (s *S3Client) DeleteObject(key string) error {
	_, err := s.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &s.bucket,
		Key:    &key,
	})
	if err != nil {
		return fmt.Errorf("failed to delete object from S3: %w", err)
	}
	return nil
}