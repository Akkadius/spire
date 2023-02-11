package eqemuserver

import (
	"bufio"
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/serverconfig"
	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Controller struct {
	db             *database.DatabaseResolver
	logger         *logrus.Logger
	eqemuserverapi *Client
	serverconfig   *serverconfig.EQEmuServerConfig
	updater        *Updater
}

func NewController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	api *Client,
	serverconfig *serverconfig.EQEmuServerConfig,
	updater *Updater,
) *Controller {
	return &Controller{
		db:             db,
		logger:         logger,
		eqemuserverapi: api,
		serverconfig:   serverconfig,
		updater:        updater,
	}
}

func (a *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "eqemuserver/zone-list", a.getZoneList, nil),
		routes.RegisterRoute(http.MethodGet, "eqemuserver/client-list", a.getClientList, nil),
		routes.RegisterRoute(http.MethodGet, "eqemuserver/server-stats", a.getServerStats, nil),
		routes.RegisterRoute(http.MethodGet, "eqemuserver/reload-types", a.getReloadTypes, nil),
		routes.RegisterRoute(http.MethodPost, "eqemuserver/reload/:type", a.reload, nil),
		routes.RegisterRoute(http.MethodGet, "eqemuserver/update-type", a.getUpdateType, nil),
		routes.RegisterRoute(http.MethodPost, "eqemuserver/update-type/:update-type", a.setUpdateType, nil),
		routes.RegisterRoute(http.MethodGet, "eqemuserver/version", a.serverVersion, nil),
		routes.RegisterRoute(http.MethodPost, "eqemuserver/install-release/:release", a.installRelease, nil),
		routes.RegisterRoute(http.MethodGet, "eqemuserver/get-build-info", a.getBuildInfo, nil),
		routes.RegisterRoute(http.MethodPost, "eqemuserver/build", a.build, nil),
		routes.RegisterRoute(http.MethodPost, "eqemuserver/build-clean", a.buildClean, nil),
	}
}

func (a *Controller) getReloadTypes(c echo.Context) error {
	types, err := a.eqemuserverapi.GetReloadTypes()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Failed to connect to gameserver [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, types)
}

func (a *Controller) reload(c echo.Context) error {
	reloadType := c.Param("type")
	r, err := a.eqemuserverapi.Reload(reloadType)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Failed to connect to gameserver [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, r)
}

type ServerStatsResponse struct {
	ServerName      string          `json:"server_name"`
	LauncherOnline  bool            `json:"launcher_online"`
	UcsOnline       bool            `json:"ucs_online"`
	LoginOnline     bool            `json:"login_online"`
	WorldOnline     bool            `json:"world_online"`
	QueryServOnline bool            `json:"query_serv_online"`
	ZoneList        WorldZoneList   `json:"zone_list"`
	PlayersOnline   WorldClientList `json:"client_list"`
}

func (a *Controller) getServerStats(c echo.Context) error {
	var r ServerStatsResponse

	processes, _ := process.Processes()
	for _, p := range processes {
		cmdline, err := p.Cmdline()
		if err != nil {
			return err
		}

		if strings.Contains(cmdline, "server-launcher") {
			r.LauncherOnline = true
		}
		if strings.Contains(cmdline, "world") {
			r.WorldOnline = true
		}
		if strings.Contains(cmdline, "ucs") {
			r.UcsOnline = true
		}
		if strings.Contains(cmdline, "loginserver") {
			r.LoginOnline = true
		}
		if strings.Contains(cmdline, "queryserv") {
			r.QueryServOnline = true
		}
	}

	zoneList, _ := a.eqemuserverapi.GetZoneList()
	if len(zoneList.Data) > 0 {
		r.ZoneList = zoneList
	}
	clientList, _ := a.eqemuserverapi.GetWorldClientList()
	if len(clientList.Data) > 0 {
		r.PlayersOnline = clientList
	}

	r.ServerName = a.serverconfig.Get().Server.World.Longname

	return c.JSON(http.StatusOK, r)
}

func (a *Controller) getClientList(c echo.Context) error {
	types, err := a.eqemuserverapi.GetWorldClientList()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Failed to connect to gameserver [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, types)
}

func (a *Controller) getZoneList(c echo.Context) error {
	types, err := a.eqemuserverapi.GetZoneList()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Failed to connect to gameserver [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, types)
}

// setUpdateType sets the update type for the server
// options are release or self-compiled
func (a *Controller) setUpdateType(c echo.Context) error {
	updateType := c.Param("update-type")
	updateTypes := []string{updateTypeRelease, updateTypeSelfCompiled}
	if !contains(updateTypes, updateType) {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Update must be of type(s) [%v]", strings.Join(updateTypes, ", "))},
		)
	}

	a.updater.SetUpdateType(updateType)

	return c.JSON(
		http.StatusOK,
		echo.Map{"message": fmt.Sprintf("Successfully set update type to [%v]", updateType)},
	)
}

func (a *Controller) getUpdateType(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		echo.Map{"updateType": a.updater.GetUpdateType()},
	)
}

func (a *Controller) serverVersion(c echo.Context) error {
	v, err := a.updater.GetVersionInfo()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": err.Error()},
		)
	}

	return c.JSON(
		http.StatusOK,
		v,
	)
}

func (a *Controller) installRelease(c echo.Context) error {
	release := c.Param("release")

	err := a.updater.InstallRelease(release)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": err.Error()},
		)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{"message": "Installed successfully"},
	)
}

func (a *Controller) getBuildInfo(c echo.Context) error {
	build, err := a.updater.GetBuildInfo()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": err.Error()},
		)
	}

	return c.JSON(
		http.StatusOK,
		build,
	)
}

type BuildContext struct {
	SourceDirectory string `json:"source_directory"`
	BuildTool       string `json:"build_tool"`
	BuildCores      int    `json:"cores"`
}

func (a *Controller) build(c echo.Context) error {
	r := new(BuildContext)
	if err := c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	c.Response().WriteHeader(http.StatusOK)

	cmd := exec.Command(r.BuildTool, fmt.Sprintf("-j%v", r.BuildCores))
	cmd.Env = os.Environ()
	if runtime.GOOS == "linux" {
		cmd.Env = append(cmd.Env, "TERM=xterm")
	}
	cmd.Dir = r.SourceDirectory
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		a.logger.Fatalf("could not get stdout pipe: %v", err)
	}

	cmd.Stderr = cmd.Stdout

	err = cmd.Start()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	merged := io.MultiReader(stdout)
	scanner := bufio.NewScanner(merged)
	for scanner.Scan() {
		c.String(http.StatusOK, scanner.Text())
		c.Response().Flush()
	}

	err = cmd.Wait()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	c.Response().Flush()

	return nil
}

type CleanBuildContext struct {
	SourceDirectory string `json:"source_directory"`
	BuildTool       string `json:"build_tool"`
}

func (a *Controller) buildClean(c echo.Context) error {
	r := new(CleanBuildContext)
	if err := c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	c.Response().WriteHeader(http.StatusOK)

	cmd := exec.Command(r.BuildTool, "clean")
	cmd.Env = os.Environ()
	if runtime.GOOS == "linux" {
		cmd.Env = append(cmd.Env, "TERM=xterm")
	}
	cmd.Dir = r.SourceDirectory
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		a.logger.Fatalf("could not get stdout pipe: %v", err)
	}

	cmd.Stderr = cmd.Stdout

	err = cmd.Start()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	merged := io.MultiReader(stdout)
	scanner := bufio.NewScanner(merged)
	for scanner.Scan() {
		c.String(http.StatusOK, scanner.Text())
		c.Response().Flush()
	}

	err = cmd.Wait()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	c.Response().Flush()

	return nil
}
