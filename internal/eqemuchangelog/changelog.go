package eqemuchangelog

import (
	"context"
	"fmt"
	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

type Changelog struct {
	client *github.Client
}

func NewChangelog() *Changelog {
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

	return &Changelog{client: client}
}

func (c *Changelog) getCommitsDaysBack(days time.Duration) []*github.RepositoryCommit {
	var allCommits []*github.RepositoryCommit
	for i := 1; i < 10; i++ {
		commits, _, err := c.client.Repositories.ListCommits(
			context.Background(),
			"EQEmu",
			"server",
			&github.CommitsListOptions{
				Since: time.Now().Add(-time.Hour * 24 * days),
				ListOptions: github.ListOptions{
					Page:    i,
					PerPage: 100,
				},
			},
		)

		if err != nil {
			log.Println(err)
		}

		if len(commits) == 0 {
			break
		}

		allCommits = append(allCommits, commits...)

		if len(commits) != 100 {
			break
		}
	}

	return allCommits
}

func (c *Changelog) getCommits() []*github.RepositoryCommit {
	var allCommits []*github.RepositoryCommit
	for i := 0; i < 10; i++ {
		commits, _, err := c.client.Repositories.ListCommits(
			context.Background(),
			"EQEmu",
			"server",
			&github.CommitsListOptions{
				Since: time.Now().Add(-time.Hour * 24 * 90),
				Until: time.Time{},
				ListOptions: github.ListOptions{
					Page:    i,
					PerPage: 100,
				},
			},
		)

		if err != nil {
			log.Println(err)
		}

		if len(commits) == 0 {
			break
		}

		allCommits = append(allCommits, commits...)
	}

	return allCommits
}

type ChangelogEntry struct {
	Author      string
	Message     string
	Category    string
	PullRequest string
	Time        time.Time
}

func (c *Changelog) BuildChangelog(commits []*github.RepositoryCommit) string {
	var entries []ChangelogEntry
	var categories []string
	for _, commit := range commits {

		// username, message
		username := *commit.Author.Login
		message := *commit.Commit.Message
		pullRequest := ""
		firstLine := strings.Split(*commit.Commit.Message, "\n")
		if len(firstLine) > 0 {
			message = strings.TrimSpace(firstLine[0])
		}

		if strings.Contains(message, "(#") {
			s := strings.Split(message, "(#")
			if len(s) > 0 {
				s2 := strings.Split(s[1], ")")
				if len(s2) > 0 {
					pullRequest = strings.TrimSpace(s2[0])
					message = strings.ReplaceAll(message, fmt.Sprintf("(#%v)", pullRequest), "")
				}
			}
		}

		// categories
		category := ""
		firstWordSplit := strings.Split(message, " ")
		if len(firstWordSplit) > 0 {
			firstWord := strings.TrimSpace(firstWordSplit[0])
			if strings.Contains(firstWord, "]") && strings.Contains(firstWord, "[") {
				category = firstWord
				category = strings.ReplaceAll(category, "[", "")
				category = strings.ReplaceAll(category, "]", "")
				message = strings.TrimSpace(strings.ReplaceAll(message, firstWord, ""))

				// one-off find replace fixes
				replacements := make(map[string]string, 0)
				replacements["Cleanup"] = "Code Cleanup"
				replacements["Bot"] = "Bots"
				replacements["Command"] = "Commands"
				replacements["Repository"] = "Repositories"
				replacements["Rule"] = "Rules"
				replacements["Fix"] = "Bug Fix"
				replacements["INT64"] = "int64"
				replacements["Hotfox"] = "Hotfix"
				replacements["HotFix"] = "Hotfix"
				replacements["Commmands"] = "Hotfix"

				for find, replacement := range replacements {
					if category == find {
						category = strings.ReplaceAll(category, find, replacement)
					}
				}

				if !contains(categories, category) {
					categories = append(categories, category)
				}
			}
		}

		if len(category) > 0 {
			hasEntry := false
			for _, e := range entries {
				if e.Message == message {
					hasEntry = true
				}
			}

			if !hasEntry {
				entries = append(
					entries, ChangelogEntry{
						Author:      username,
						Message:     message,
						Category:    category,
						Time:        commit.Commit.Author.GetDate(),
						PullRequest: pullRequest,
					},
				)
			}
		}

		// sort changelog entries
		sort.Slice(entries[:], func(i, j int) bool {
			if entries[i].Category == entries[j].Category {
				return entries[i].Message < entries[j].Message
			}

			return entries[i].Category < entries[j].Category
		})

		// sort categories
		sort.Slice(categories[:], func(i, j int) bool {
			return categories[i] < categories[j]
		})
	}

	msg := ""

	// build message
	for _, category := range categories {
		hasCategory := false
		for _, e := range entries {
			if e.Category == category {
				hasCategory = true
			}
		}

		if hasCategory {
			msg += fmt.Sprintf("\n### %v\n\n", category)

			for _, e := range entries {
				if e.Category == category {
					pr := ""
					if len(e.PullRequest) > 0 {
						pr = fmt.Sprintf(
							"([#%v](https://github.com/EQEmu/Server/pull/%v))",
							e.PullRequest,
							e.PullRequest,
						)
					}

					msg += fmt.Sprintf(
						"* %v %v ([%v](https://github.com/%v)) %v\n",
						e.Message,
						pr,
						e.Author,
						e.Author,
						e.Time.Format("2006-01-02"),
					)
				}
			}
		}

	}

	return strings.TrimSpace(msg)
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func contains(slice []string, val string) bool {
	for _, item := range slice {
		if strings.Contains(val, item) {
			return true
		}
	}
	return false
}
