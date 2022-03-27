package controllers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/clientfiles"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
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
	db       *database.DatabaseResolver
}

func NewClientFilesController(
	logger *logrus.Logger,
	exporter *clientfiles.Exporter,
	importer *clientfiles.Importer,
	db *database.DatabaseResolver,
) *ClientFilesController {
	return &ClientFilesController{
		logger:   logger,
		exporter: exporter,
		importer: importer,
		db:       db,
	}
}

func (f *ClientFilesController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "client-files/export/spells", f.exportSpells, nil),
		routes.RegisterRoute(http.MethodGet, "client-files/export/dbstr", f.exportDbStr, nil),
		routes.RegisterRoute(http.MethodPost, "client-files/import/file", f.importFile, nil),
	}
}

func (f *ClientFilesController) exportSpells(c echo.Context) error {
	// todo, bubble error handling
	// quick and dirty for now
	db := f.db.Get(&models.SpellsNew{}, c)
	contents := f.exporter.ExportSpells(db)
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

func (f *ClientFilesController) exportDbStr(c echo.Context) error {
	// todo, bubble error handling
	// quick and dirty for now
	db := f.db.Get(&models.DbStr{}, c)
	contents := f.exporter.ExportDbStr(db)
	folderPath := filepath.Join(os.TempDir(), "dbstr-export", randomString(10))
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error creating temp path [%v]", err.Error())},
		)
	}
	filePath := filepath.Join(folderPath, "dbstr_us.txt")
	err = os.WriteFile(filePath, []byte(contents), 0755)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error writing file [%v]", err.Error())},
		)
	}

	return c.Attachment(filePath, "dbstr_us.txt")
}

func (f *ClientFilesController) importFile(c echo.Context) error {
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
	if !strings.Contains(fileName, "spells_us") && !strings.Contains(fileName, "dbstr_us") {
		return c.HTML(
			http.StatusInternalServerError,
			fmt.Sprintf("File not valid"),
		)
	}

	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return c.HTML(
			http.StatusInternalServerError,
			fmt.Sprintf("Error [%v]", err.Error()),
		)
	}

	fileContents := string(fileBytes)
	r := clientfiles.ImportResult{}
	if strings.Contains(fileName, "spells_us") {
		r, err = f.importer.ImportSpells(fileContents)
		if err != nil {
			return c.HTML(
				http.StatusInternalServerError,
				fmt.Sprintf("Error [%v]", err.Error()),
			)
		}
	}
	if strings.Contains(fileName, "dbstr_us") {
		r, err = f.importer.ImportDbStr(fileContents)
		if err != nil {
			return c.HTML(
				http.StatusInternalServerError,
				fmt.Sprintf("Error [%v]", err.Error()),
			)
		}
	}

	// for commify
	p := message.NewPrinter(language.English)

	return c.HTML(
		http.StatusOK,
		p.Sprintf("Success! Table [%v] Deleted [%d] Imported [%d]", r.Table, r.DroppedRows, r.ImportedRows),
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
