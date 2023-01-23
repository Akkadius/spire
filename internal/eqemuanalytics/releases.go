package eqemuanalytics

import (
	"context"
	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"time"
)

type Releases struct {
	client *github.Client
}

func NewReleases() *Releases {
	client := github.NewClient(nil)
	if len(os.Getenv("GITHUB_TOKEN")) > 0 {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
		)
		tc := &http.Client{
			Timeout: 5 * time.Second,
			Transport: &oauth2.Transport{
				Source: ts,
			},
		}
		client = github.NewClient(tc)
	}

	return &Releases{client: client}
}

func (r *Releases) getReleases() ([]*github.RepositoryRelease, error) {
	releases, _, err := r.client.Repositories.ListReleases(
		context.Background(),
		"EQEmu",
		"server",
		&github.ListOptions{
			Page:    0,
			PerPage: 100,
		},
	)
	if err != nil {
		return nil, err
	}

	return releases, nil
}
