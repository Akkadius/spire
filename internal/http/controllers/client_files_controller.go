package controllers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/clientfiles"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ClientFilesController struct {
	logger   *logrus.Logger
	exporter *clientfiles.Exporter
	importer *clientfiles.Importer
}

func NewClientFilesController(
	logger *logrus.Logger,
	exporter *clientfiles.Exporter,
	importer *clientfiles.Importer,
) *ClientFilesController {
	return &ClientFilesController{
		logger:   logger,
		exporter: exporter,
		importer: importer,
	}
}

func (f *ClientFilesController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "client-files/export/spells", f.exportSpells, nil),
		routes.RegisterRoute(http.MethodPost, "client-files/import/spells", f.importSpells, nil),
	}
}

func (f *ClientFilesController) exportSpells(c echo.Context) error {
	// todo, bubble error handling
	// quick and dirty for now
	contents := f.exporter.ExportSpells()
	folderPath := filepath.Join(os.TempDir(), "spell-export", randomString(10))
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error creating temp path [%v]", err.Error())},
		)
	}
	filePath := filepath.Join(folderPath, "spells_us.txt")
	err = os.WriteFile(filePath, []byte(contents), 0755)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error writing file [%v]", err.Error())},
		)
	}

	return c.Attachment(filePath, "spells_us.txt")
}

func (f *ClientFilesController) importSpells(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	fileName := file.Filename
	if !strings.Contains(fileName, "spells_us") {
		return c.HTML(
			http.StatusInternalServerError,
			fmt.Sprintf("File not valid"),
		)
	}

	//fmt.Println(fileName)

	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return c.HTML(
			http.StatusInternalServerError,
			fmt.Sprintf("Error [%v]", err.Error()),
		)
	}

	fileContents := string(fileBytes)
	err = f.importer.ImportSpells(fileContents)
	if err != nil {
		return c.HTML(
			http.StatusInternalServerError,
			fmt.Sprintf("Error [%v]", err.Error()),
		)
	}

	//fmt.Println(fileContents)

	return c.HTML(
		http.StatusOK,
		fmt.Sprintf("Success!"),
	)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}
