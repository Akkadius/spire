package eqemuserver

import (
	"errors"
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type PublicController struct {
	db             *database.Resolver
	eqemuserverapi *Client
	pathmgmt       *pathmgmt.PathManagement
	settings       *spire.Settings
	serverconfig   *eqemuserverconfig.Config
	updater        *Updater
}

func NewPublicController(
	db *database.Resolver,
	api *Client,
	serverconfig *eqemuserverconfig.Config,
	pathmgmt *pathmgmt.PathManagement,
	settings *spire.Settings,
	updater *Updater,
) *PublicController {
	return &PublicController{
		db:             db,
		eqemuserverapi: api,
		serverconfig:   serverconfig,
		pathmgmt:       pathmgmt,
		updater:        updater,
		settings:       settings,
	}
}

func (a *PublicController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "eqemuserver/export-client-file/:type", a.exportClientFile, nil),
	}
}

type ExportType struct {
	arg  string
	file string
}

func (a *PublicController) exportClientFile(c echo.Context) error {
	bin := a.pathmgmt.GetExportClientFilesBinPath()
	if _, err := os.Stat(bin); errors.Is(err, os.ErrNotExist) {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Failed to find [%v] binary", a.pathmgmt.GetExportClientFilesBinPath())},
		)
	}

	t := []ExportType{
		{arg: "spells", file: "spells_us.txt"},
		{arg: "skills", file: "SkillCaps.txt"},
		{arg: "basedata", file: "BaseData.txt"},
		{arg: "dbstring", file: "dbstr_us.txt"},
	}

	exportType := c.Param("type")
	foundExport := false
	var export ExportType
	for _, e := range t {
		if e.arg == exportType {
			foundExport = true
			export = e
		}
	}

	if !foundExport {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invalid export type"})
	}

	err := a.pathmgmt.MakeExportDirIfNotExists()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	cmd := exec.Command(bin, exportType)
	cmd.Dir = a.pathmgmt.GetEQEmuServerPath()
	_, err = cmd.Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	c.Response().Header().Add("Access-Control-Expose-Headers", "Content-Disposition")

	downloadFile := filepath.Join(a.pathmgmt.GetExportDir(), export.file)
	if _, err := os.Stat(downloadFile); !errors.Is(err, os.ErrNotExist) {
		return c.Attachment(downloadFile, filepath.Base(downloadFile))
	}

	return nil
}
