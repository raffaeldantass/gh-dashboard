package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDescription(t *testing.T) {
	tests := []struct {
		name     string
		input    *string
		expected string
	}{
		{
			name:     "nil description",
			input:    nil,
			expected: "",
		},
		{
			name:     "non-nil description",
			input:    stringPtr("test description"),
			expected: "test description",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getDescription(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetAllRepositoriesIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	service := NewGitHubService("test-token")
	_, err := service.GetAllRepositories(context.Background(), 1, 10)

	// We expect an error because we're using an invalid token
	assert.Error(t, err)
}

func stringPtr(s string) *string {
	return &s
}
