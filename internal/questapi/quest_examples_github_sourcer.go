package questapi

import (
	"github.com/Akkadius/spire/internal/github"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"strings"
)

type QuestExamplesGithubSourcer struct {
	logger     *logrus.Logger
	cache      *gocache.Cache
	downloader *github.GithubSourceDownloader
	files      map[string]string
}

func NewQuestExamplesGithubSourcer(
	logger *logrus.Logger,
	cache *gocache.Cache,
	downloader *github.GithubSourceDownloader,
) *QuestExamplesGithubSourcer {
	return &QuestExamplesGithubSourcer{logger: logger, cache: cache, downloader: downloader}
}

const (
	searchResultLinesBefore = 3
	searchResultLinesAfter  = 6
	maxResultsPerSearchTerm = 25
)

type SearchResultSnippet struct {
	FileName        string `json:"file_name"`
	BeforeContent   string `json:"before_content"`
	LineMatch       string `json:"line_match"`
	LineNumberMatch int    `json:"line_number"`
	FullContents    string `json:"full_contents,omitempty"`
	AfterContent    string `json:"after_content"`
	SearchTerm      string `json:"search_term"`
}

// return files from memory
func (q *QuestExamplesGithubSourcer) Files() map[string]string {
	return q.files
}

// source files
func (q *QuestExamplesGithubSourcer) Source(
	org string,
	repo string,
	branch string,
	forceRefresh bool,
) map[string]string {

	// return cache if not refresh
	if !forceRefresh && len(q.files) > 0 {
		return q.files
	}

	// set to memory
	q.files = q.downloader.Source(org, repo, branch, forceRefresh).Files

	// return local reference
	return q.files
}

// search files for results
func (q *QuestExamplesGithubSourcer) Search(
	org string,
	repo string,
	branch string,
	searchStrings []string,
	language string,
	forceRefresh bool,
) []SearchResultSnippet {
	q.Source(org, repo, branch, forceRefresh)

	fileExt := ""
	if language == "perl" {
		fileExt = ".pl"
	}
	if language == "lua" {
		fileExt = ".lua"
	}

	var results []SearchResultSnippet

	for _, searchString := range searchStrings {
		resultCount := 0
		for file, contents := range q.Files() {

			// match file extension and contents
			if strings.Contains(file, fileExt) && strings.Contains(contents, searchString) {

				// match at the line level
				splitContents := strings.Split(contents, "\n")
				for i, line := range splitContents {

					// match at line level
					if strings.Contains(line, searchString) {
						lineNumber := i + 1

						// before content
						beforeContent := ""
						for n := searchResultLinesBefore; n > 0; n-- {
							readIndex := i - n
							if readIndex > 0 {
								beforeContent = beforeContent + splitContents[readIndex] + "\n"
							}
						}

						// after content
						afterContent := ""
						for n := 0; n < searchResultLinesAfter; n++ {
							readIndex := i + n + 1
							if readIndex > 0 && len(splitContents) > readIndex {
								afterContent = afterContent + splitContents[readIndex] + "\n"
							}
						}

						results = append(
							results, SearchResultSnippet{
								FileName:        file,
								SearchTerm:      searchString,
								BeforeContent:   beforeContent,
								AfterContent:    afterContent,
								LineMatch:       line + "\n",
								LineNumberMatch: lineNumber,
								FullContents:    contents,
							},
						)

						//fmt.Printf("[%v]\n", file)
						//fmt.Printf("Before Content\n%v\n", beforeContent)
						//fmt.Printf("Line\n[%v] %v\n", lineNumber, line)
						//fmt.Printf("After Content\n%v\n\n\n", afterContent)

						break
					}

				}

				resultCount++

				if resultCount >= maxResultsPerSearchTerm {
					break
				}
			}
		}
	}

	return results
}
