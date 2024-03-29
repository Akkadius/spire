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

func (c *Changelog) getCommitsDaysBack() []*github.RepositoryCommit {
	daysSinceLastRelease := 0
	r, _, _ := c.client.Repositories.GetLatestRelease(context.Background(), "EQEmu", "Server")
	if *r.ID > 0 {
		date := time.Now()
		diff := date.Sub(r.CreatedAt.Time)
		daysSinceLastRelease = int(diff.Hours()/24) + 1
	}

	var allCommits []*github.RepositoryCommit
	for i := 1; i < 10; i++ {
		commits, _, err := c.client.Repositories.ListCommits(
			context.Background(),
			"EQEmu",
			"server",
			&github.CommitsListOptions{
				Since: time.Now().Add(-time.Hour * 24 * time.Duration(daysSinceLastRelease)),
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

type ChangelogEntry struct {
	Author      string
	Message     string
	Category    string
	PullRequest string
	Time        time.Time
}

func (c *Changelog) BuildChangelog(commits []*github.RepositoryCommit) string {
	var lastReleaseNotes string
	r, _, _ := c.client.Repositories.GetLatestRelease(context.Background(), "EQEmu", "Server")
	if *r.ID > 0 {
		lastReleaseNotes = r.GetBody()
	}

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
		if strings.Contains(message, "]") && strings.Contains(message, "[") {
			category = c.GetStringInBetween(message, "[", "]")
			if len(category) < 20 {
				message = strings.TrimSpace(
					strings.ReplaceAll(
						message,
						fmt.Sprintf("[%v]", category),
						"",
					),
				)

				// one-off find replace fixes
				replacements := make(map[string]string, 0)
				replacements["Code Cleanup"] = "Code"
				replacements["Cleanup"] = "Code"
				replacements["Bot"] = "Bots"
				replacements["Command"] = "Commands"
				replacements["Repository"] = "Repositories"
				replacements["Rule"] = "Rules"
				replacements["Bug Fix"] = "Fixes"
				replacements["Fix"] = "Fixes"
				replacements["INT64"] = "int64"
				replacements["Hotfox"] = "Hotfix"
				replacements["HotFix"] = "Hotfix"
				replacements["Hotfix"] = "Fixes"
				replacements["Quest"] = "Quest API"
				replacements["Bots/Mercs"] = "Bots & Mercenaries"
				replacements["Bots & Mercs"] = "Bots & Mercenaries"

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

			if strings.Contains(lastReleaseNotes, message) {
				continue
			}

			if !hasEntry && !strings.Contains(category, "Release") {
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
						"* %v %v @%v %v\n",
						e.Message,
						pr,
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
		if val == item {
			return true
		}
	}
	return false
}

func (c *Changelog) GetStringInBetween(value string, a string, b string) string {
	firstSplit := strings.Split(value, a)
	if len(firstSplit) > 1 {
		secondSplit := strings.Split(firstSplit[1], b)
		if len(secondSplit) > 0 {
			return strings.TrimSpace(secondSplit[0])
		}
	}

	return ""
}
