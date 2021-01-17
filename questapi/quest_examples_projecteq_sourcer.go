package questapi

import (
	"fmt"
	"github.com/Akkadius/spire/internal/unzip"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type QuestExamplesProjectEqSourcer struct {
	logger *logrus.Logger
	cache  *gocache.Cache
}

func NewQuestExamplesProjectEqSourcer(logger *logrus.Logger, cache *gocache.Cache) *QuestExamplesProjectEqSourcer {
	return &QuestExamplesProjectEqSourcer{logger: logger, cache: cache}
}

const (
	questRepoDir = "/tmp/quests/projecteqquests-master/"
	repoZipUrl   = "https://github.com/ProjectEQ/projecteqquests/archive/master.zip"
)

var projectEqQuests = map[string]string{}

const peqQuestReadLock = "peq-quest-read-lock"

// sources quest files into memory to be searched later
func (q *QuestExamplesProjectEqSourcer) Source() *QuestExamplesProjectEqSourcer {

	// if lock set, return
	_, found := q.cache.Get(peqQuestReadLock)
	if found || len(projectEqQuests) > 0 {
		return &QuestExamplesProjectEqSourcer{}
	}

	// set operation lock
	q.cache.Set(peqQuestReadLock, 1, time.Minute*10)

	unzipLoc := fmt.Sprintf("%v/quests/", os.TempDir())
	zipFileLocalLoc := fmt.Sprintf("%v/%v", os.TempDir(), "master-quests.zip")

	// if not exist, extract
	if _, err := os.Stat(unzipLoc); os.IsNotExist(err) {
		err := q.downloadFile(zipFileLocalLoc, repoZipUrl)
		if err != nil {
			q.logger.Error(err)
		}

		uz := unzip.New(zipFileLocalLoc, unzipLoc)
		err = uz.Extract()
		if err != nil {
			q.logger.Error(err)
		}
	}

	err := filepath.Walk(
		questRepoDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			fileName := strings.ReplaceAll(path, questRepoDir, "")
			if strings.Contains(path, ".pl") || strings.Contains(path, ".lua") {
				data, err := ioutil.ReadFile(path)
				if err != nil {
					q.logger.Fatal(err)
				}

				projectEqQuests[fileName] = string(data)
			}

			return nil
		},
	)
	if err != nil {
		q.logger.Error(err)
	}

	// delete lock
	q.cache.Delete(peqQuestReadLock)

	return q
}

// return files from memory
func (q *QuestExamplesProjectEqSourcer) Files() map[string]string {
	return projectEqQuests
}

const searchResultLinesBefore = 3
const searchResultLinesAfter = 6

type SearchResultSnippet struct {
	FileName        string `json:"file_name"`
	BeforeContent   string `json:"before_content"`
	LineMatch       string `json:"line_match"`
	LineNumberMatch int    `json:"line_number"`
	FullContents    string `json:"full_contents,omitempty"`
	AfterContent    string `json:"after_content"`
	SearchTerm      string `json:"search_term"`
}

// search files for results
func (q *QuestExamplesProjectEqSourcer) Search(searchStrings []string, language string) []SearchResultSnippet {
	q.Source()

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

				if resultCount > 50 {
					break
				}
			}
		}
	}

	return results
}

// downloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func (q *QuestExamplesProjectEqSourcer) downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
