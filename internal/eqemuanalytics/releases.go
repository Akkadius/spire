package eqemuanalytics

import (
	"context"
	"github.com/google/go-github/v41/github"
	gocache "github.com/patrickmn/go-cache"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"time"
)

type Releases struct {
	client *github.Client
	cache  *gocache.Cache
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

	return &Releases{client: client, cache: gocache.New(5*time.Minute, 10*time.Minute)}
}

func (r *Releases) getReleases() ([]*github.RepositoryRelease, error) {
	// return cached releases if they exist
	if x, found := r.cache.Get("releases"); found {
		return x.([]*github.RepositoryRelease), nil
	}

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

	// cache releases
	r.cache.Set("releases", releases, time.Minute*1)

	return releases, nil
}
