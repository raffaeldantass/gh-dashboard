package github

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/raffaeldantass/gh-dashboard/internal/models"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/go-github/v45/github"
)

type Client struct {
	appID          string
	installationID string
	privateKey     []byte
	client         *github.Client
}

func NewClient(appID, installationID string, privateKey []byte) *Client {
	return &Client{
		appID:          appID,
		installationID: installationID,
		privateKey:     privateKey,
	}
}

func (c *Client) generateJWT() (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(c.privateKey)
	if err != nil {
		return "", fmt.Errorf("parsing private key: %w", err)
	}

	now := time.Now()
	claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(10 * time.Minute)),
		Issuer:    c.appID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signed, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("signing token: %w", err)
	}

	return signed, nil
}

func (c *Client) getInstallationToken(ctx context.Context) (string, error) {
	jwt, err := c.generateJWT()
	if err != nil {
		return "", err
	}

	// Create a temporary client using the JWT
	tc := github.NewTokenClient(ctx, jwt)

	// Get an installation token
	token, _, err := tc.Apps.CreateInstallationToken(
		ctx,
		c.installationID,
		&github.InstallationTokenOptions{},
	)
	if err != nil {
		return "", fmt.Errorf("creating installation token: %w", err)
	}

	return token.GetToken(), nil
}

func (c *Client) GetRepositories(ctx context.Context, page, perPage int, query, org string) (*models.PaginatedResponse, error) {
	token, err := c.getInstallationToken(ctx)
	if err != nil {
		return nil, err
	}

	client := github.NewTokenClient(ctx, token)

	opts := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: perPage,
		},
	}

	var repos []*github.Repository
	if org != "" {
		repos, _, err = client.Repositories.ListByOrg(ctx, org, opts)
	} else {
		repos, _, err = client.Repositories.List(ctx, "", opts)
	}
	if err != nil {
		return nil, fmt.Errorf("listing repositories: %w", err)
	}

	// Convert GitHub repos to our model
	result := make([]models.Repository, 0, len(repos))
	for _, repo := range repos {
		if query != "" && !containsString(repo.GetName(), query) {
			continue
		}

		var org *models.Organization
		if repo.Organization != nil {
			org = &models.Organization{
				Login:     repo.Organization.GetLogin(),
				AvatarURL: repo.Organization.GetAvatarURL(),
			}
		}

		result = append(result, models.Repository{
			ID:           repo.GetID(),
			Name:         repo.GetName(),
			FullName:     repo.GetFullName(),
			Description:  repo.GetDescription(),
			Private:      repo.GetPrivate(),
			UpdatedAt:    repo.GetUpdatedAt().Time,
			Language:     repo.GetLanguage(),
			Organization: org,
		})
	}

	return &models.PaginatedResponse{
		Data:       result,
		Total:      len(result),
		Page:       page,
		PerPage:    perPage,
		TotalPages: (len(result) + perPage - 1) / perPage,
	}, nil
}

func containsString(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
