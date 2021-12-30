package controllers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type QuestFileApiController struct {
	logger *logrus.Logger
}

func NewQuestFileApiController(
	logger *logrus.Logger,
) *QuestFileApiController {
	return &QuestFileApiController{logger: logger}
}

func (h *QuestFileApiController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "quest-file-api/list", h.listFiles, nil),
		routes.RegisterRoute(http.MethodGet, "quest-file-api/:file", h.getFile, nil),
	}
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if strings.Contains(path, ".git") || strings.Contains(path, "./quests"){
			return nil
		}

		path = strings.ReplaceAll(path, "quests/", "")

		*files = append(*files, path)
		return nil
	}
}

func (h *QuestFileApiController) listFiles(c echo.Context) error {
	var files []string

	root := "./quests"
	err := filepath.Walk(root, visit(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}

	return c.JSON(http.StatusOK, echo.Map{"files": files})
}


func (h *QuestFileApiController) getFile(c echo.Context) error {
	path, err := url.QueryUnescape(c.Param("file"))
	if err != nil {
		panic(err)
	}

	if filepath.Ext(path) != ".pl" && filepath.Ext(path) != ".lua" {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "invalid file extension"})
	}
	if strings.Contains(path, "..") {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "invalid access pattern"})
	}

	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(fmt.Sprintf("./quests/%s", path))
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, echo.Map{"content": string(content)})
}
