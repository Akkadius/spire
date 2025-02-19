package backup

import (
	"github.com/Akkadius/spire/internal/filepathcheck"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"path/filepath"
)

type Controller struct {
	mysql    *Mysql
	pathmgmt *pathmgmt.PathManagement
}

func NewController(
	mysql *Mysql,
	pathmgmt *pathmgmt.PathManagement,
) *Controller {
	return &Controller{
		mysql:    mysql,
		pathmgmt: pathmgmt,
	}
}

func (a *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodPost, "backup/mysql", a.backupMysql, nil),
		routes.RegisterRoute(http.MethodGet, "backup/mysql-dump-download/:file", a.mysqlBackupDownload, nil),
	}
}

func (a *Controller) backupMysql(c echo.Context) error {
	p := new(BackupRequest)
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	r := a.mysql.Backup(*p)

	return c.JSON(http.StatusOK, r)
}

func (a *Controller) mysqlBackupDownload(c echo.Context) error {
	file := c.Param("file")
	downloadPath := filepath.Join(a.pathmgmt.GetBackupsDir(), filepath.Base(file))

	err := filepathcheck.ValidateSafePath(a.pathmgmt.GetBackupsDir(), downloadPath)
	if err != nil {
		return err
	}

	return c.Inline(downloadPath, filepath.Base(file))
}
