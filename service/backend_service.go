package service

import (
	"encoding/json"
	"fmt"
	"log"

	s3client "github.com/zij/interstate/pkg/s3client"
)

// BackendService handles the business logic for backend operations
type BackendService struct {
	s3Client *s3client.S3Client
}

// NewBackendService creates a new BackendService
func NewBackendService(s3Client *s3client.S3Client) *BackendService {
	return &BackendService{
		s3Client: s3Client,
	}
}

// ListBackends returns a list of all backends
func (s *BackendService) ListBackends() ([]string, error) {
	// TODO: Implement listing backends from S3
	return []string{"backend1", "backend2"}, nil
}

// CreateBackend creates a new backend
func (s *BackendService) CreateBackend(name string) error {
	folderKey := fmt.Sprintf("%s/", name)
	err := s.s3Client.CreateFolder(folderKey)
	if (err != nil) {
		log.Printf("Error creating folder in S3: %s", err)
	}
	return err
}

// GetBackend retrieves details of a specific backend
func (s *BackendService) GetBackend(id string) (string, error) {
	// TODO: Implement getting backend details from S3
	return fmt.Sprintf("Backend %s details", id), nil
}

// UpdateBackend updates an existing backend
func (s *BackendService) UpdateBackend(id string, newName string) error {
	// TODO: Implement updating backend in S3
	return nil
}

// DeleteBackend deletes a specific backend
func (s *BackendService) DeleteBackend(id string) error {
	// TODO: Implement deleting backend folder from S3
	return nil
}

// GetBackendState retrieves the state of a specific backend
func (s *BackendService) GetBackendState(id string) (map[string]interface{}, error) {
	key := fmt.Sprintf("%s/terraform.tfstate", id)
	data, err := s.s3Client.ReadFile(key)
	if err != nil {
		return nil, fmt.Errorf("failed to read backend state: %w", err)
	}

	var state map[string]interface{}
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, fmt.Errorf("failed to parse backend state: %w", err)
	}

	return state, nil
}

// UpdateBackendState updates the state of a specific backend
func (s *BackendService) UpdateBackendState(id string, state map[string]interface{}) error {
	data, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("failed to marshal backend state: %w", err)
	}

	key := fmt.Sprintf("%s/terraform.tfstate", id)
	if err := s.s3Client.WriteFile(key, data); err != nil {
		return fmt.Errorf("failed to write backend state: %w", err)
	}

	return nil
}

// ResetBackendState resets the state of a specific backend
func (s *BackendService) ResetBackendState(id string) error {
	// TODO: Implement resetting backend state in S3
	return nil
}