package services

import (
	"context"
	"fmt"
	"log"
	"sort"

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

func (s *GitHubService) GetAllRepositories(ctx context.Context, page, perPage int) (*models.PaginatedResponse, error) {
	repoMap := make(map[string]models.Repository)

	// First, get all user repositories without pagination
	opt := &github.RepositoryListByAuthenticatedUserOptions{
		Type: "private",
		ListOptions: github.ListOptions{
			PerPage: 100, // Get maximum items per page
		},
	}

	// Get all user repositories using pagination
	for {
		repos, resp, err := s.client.Repositories.ListByAuthenticatedUser(ctx, opt)
		if err != nil {
			log.Printf("Error getting repositories: %v", err)
			return nil, err
		}

		for _, repo := range repos {
			if repo.Private != nil && *repo.Private {
				key := fmt.Sprintf("%s/%s", *repo.Owner.Login, *repo.Name)
				repoMap[key] = models.Repository{
					Name:        *repo.Name,
					Description: getDescription(repo.Description),
					LastUpdate:  repo.UpdatedAt.Time,
					IsPrivate:   true,
					Owner:       *repo.Owner.Login,
				}
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	// Get all organizations
	orgs, _, err := s.client.Organizations.List(ctx, "", nil)
	if err != nil {
		return nil, err
	}

	// For each organization, get all their repositories
	for _, org := range orgs {
		orgOpt := &github.RepositoryListByOrgOptions{
			Type: "private",
			ListOptions: github.ListOptions{
				PerPage: 100, // Get maximum items per page
			},
		}

		for {
			orgRepos, resp, err := s.client.Repositories.ListByOrg(ctx, *org.Login, orgOpt)
			if err != nil {
				break // Skip this org if there's an error
			}

			for _, repo := range orgRepos {
				if repo.Private != nil && *repo.Private {
					key := fmt.Sprintf("%s/%s", *org.Login, *repo.Name)
					if _, exists := repoMap[key]; !exists {
						repoMap[key] = models.Repository{
							Name:        *repo.Name,
							Description: getDescription(repo.Description),
							LastUpdate:  repo.UpdatedAt.Time,
							IsPrivate:   true,
							Owner:       *org.Login,
						}
					}
				}
			}

			if resp.NextPage == 0 {
				break
			}
			orgOpt.Page = resp.NextPage
		}
	}

	// Convert map to slice
	var allRepos []models.Repository
	for _, repo := range repoMap {
		allRepos = append(allRepos, repo)
	}

	// Sort repositories by LastUpdate
	sort.Slice(allRepos, func(i, j int) bool {
		return allRepos[i].LastUpdate.After(allRepos[j].LastUpdate)
	})

	// Calculate pagination values
	totalRepos := len(allRepos)
	totalPages := (totalRepos + perPage - 1) / perPage
	if totalPages < 1 {
		totalPages = 1
	}

	// Apply pagination to the full sorted list
	start := (page - 1) * perPage
	end := start + perPage
	if start >= len(allRepos) {
		// Return empty slice if start is beyond array bounds
		return &models.PaginatedResponse{
			Repositories: []models.Repository{},
			CurrentPage:  page,
			TotalPages:   totalPages,
			TotalItems:   totalRepos,
			ItemsPerPage: perPage,
		}, nil
	}
	if end > len(allRepos) {
		end = len(allRepos)
	}

	return &models.PaginatedResponse{
		Repositories: allRepos[start:end],
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalItems:   totalRepos,
		ItemsPerPage: perPage,
	}, nil
}

func getDescription(desc *string) string {
	if desc == nil {
		return ""
	}
	return *desc
}
