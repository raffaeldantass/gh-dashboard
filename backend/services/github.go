package services

import (
	"context"
	"github.com/google/go-github/v68/github"
	"github.com/raffaeldantass/gh-dashboard/models"
	"golang.org/x/oauth2"
)

type GitHubService struct {
	client *github.Client
}

func NewGitHubService(token string) *GitHubService {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(context.Background(), ts)

	return &GitHubService{
		client: github.NewClient(tc),
	}
}

func (s *GitHubService) GetAllRepositories(ctx context.Context) ([]models.Repository, error) {
	var allRepos []models.Repository

	// Get user's repositories
	repos, _, err := s.client.Repositories.ListByAuthenticatedUser(ctx, &github.RepositoryListByAuthenticatedUserOptions{
		Type: "private",
	})
	if err != nil {
		return nil, err
	}

	// Get user's organizations
	orgs, _, err := s.client.Organizations.List(ctx, "", nil)
	if err != nil {
		return nil, err
	}

	// Add user's repositories
	for _, repo := range repos {
		if repo.Private != nil && *repo.Private {
			allRepos = append(allRepos, models.Repository{
				Name:        *repo.Name,
				Description: getDescription(repo.Description),
				LastUpdate:  repo.UpdatedAt.Time,
				IsPrivate:   true,
				Owner:       *repo.Owner.Login,
			})
		}
	}

	// Add organizations' repositories
	for _, org := range orgs {
		orgRepos, _, err := s.client.Repositories.ListByOrg(ctx, *org.Login, &github.RepositoryListByOrgOptions{
			Type: "private",
		})
		if err != nil {
			continue
		}

		for _, repo := range orgRepos {
			if repo.Private != nil && *repo.Private {
				allRepos = append(allRepos, models.Repository{
					Name:        *repo.Name,
					Description: getDescription(repo.Description),
					LastUpdate:  repo.UpdatedAt.Time,
					IsPrivate:   true,
					Owner:       *org.Login,
				})
			}
		}
	}

	return allRepos, nil
}

func getDescription(desc *string) string {
	if desc == nil {
		return ""
	}
	return *desc
}
