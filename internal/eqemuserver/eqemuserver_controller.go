package eqemuserver

import (
	"bufio"
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/serverconfig"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type Controller struct {
	db             *database.DatabaseResolver
	logger         *logrus.Logger
	eqemuserverapi *Client
	pathmgmt       *pathmgmt.PathManagement
	settings       *spire.Settings
	serverconfig   *serverconfig.EQEmuServerConfig
	updater        *Updater
}

func NewController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	api *Client,
	serverconfig *serverconfig.EQEmuServerConfig,
	pathmgmt *pathmgmt.PathManagement,
	settings *spire.Settings,
	updater *Updater,
) *Controller {
	return &Controller{
		db:             db,
		logger:         logger,
		eqemuserverapi: api,
		serverconfig:   serverconfig,
		pathmgmt:       pathmgmt,
		updater:        updater,
		settings:       settings,
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
		routes.RegisterRoute(http.MethodPost, "eqemuserver/build-cancel", a.buildCancel, nil),
		routes.RegisterRoute(http.MethodGet, "eqemuserver/build/current-branch", a.getBuildCurrentBranch, nil),
		routes.RegisterRoute(http.MethodGet, "eqemuserver/build/branches", a.getBuildBranches, nil),
		routes.RegisterRoute(http.MethodPost, "eqemuserver/build/branch/:branch", a.setBuildBranch, nil),
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

type ZoneProcessInfo struct {
	Pid     int32   `json:"pid"`
	Name    string  `json:"name"`
	CmdLine string  `json:"cmd"`
	Cpu     float64 `json:"cpu"`
	Memory  uint64  `json:"memory"`
	Elapsed int64   `json:"elapsed"` // uptime
}

type ZoneListResponse struct {
	List        WorldZoneList     `json:"zone_list"`
	ProcessInfo []ZoneProcessInfo `json:"process_info"`
}

func (a *Controller) getZoneList(c echo.Context) error {
	zones, err := a.eqemuserverapi.GetZoneList()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Failed to connect to gameserver [%v]", err.Error())},
		)
	}

	var zoneProcessInfo []ZoneProcessInfo
	processes, _ := process.Processes()
	for _, p := range processes {
		for _, z := range zones.Data {
			if int(p.Pid) == z.ZoneOsPid {
				name, _ := p.Name()
				cmdLine, _ := p.Cmdline()
				cpuPercent, _ := p.CPUPercent()
				memory, _ := p.MemoryInfo()
				uptime, _ := p.CreateTime()
				now := time.Now().Unix()
				zoneProcessInfo = append(zoneProcessInfo, ZoneProcessInfo{
					Pid:     p.Pid,
					Name:    name,
					CmdLine: cmdLine,
					Cpu:     cpuPercent,
					Memory:  memory.RSS,
					Elapsed: now - (uptime / 1000),
				})
			}
		}
	}

	var r ZoneListResponse
	r.List = zones
	r.ProcessInfo = zoneProcessInfo

	return c.JSON(http.StatusOK, r)
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
		return c.String(http.StatusBadRequest, err.Error())
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

type CancelBuildContext struct {
	SourceDirectory string `json:"source_directory"`
	BuildTool       string `json:"build_tool"`
}

func (a *Controller) buildCancel(c echo.Context) error {
	r := new(CancelBuildContext)
	if err := c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	c.Response().WriteHeader(http.StatusOK)

	killProcs := []string{filepath.Base(r.BuildTool), "ccache"}
	for _, proc := range killProcs {
		cmd := exec.Command("pkill", "-9", proc)
		cmd.Env = os.Environ()
		cmd.Dir = r.SourceDirectory
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
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
	}

	c.Response().Flush()

	return nil
}

// list branches
// git fetch origin && git branch -a
// checkout branch
// cd %s && git fetch origin && git checkout -f %s && git pull
// current branch
// cd %s && git rev-parse --abbrev-ref HEAD

func (a *Controller) getBuildBranches(c echo.Context) error {
	var dirname string
	s := a.settings.GetSetting("BUILD_LOCATION")
	if s.ID != 0 {
		dirname = s.Value
	}

	cmd := exec.Command("git", "fetch", "origin")
	cmd.Dir = dirname
	cmd.Stderr = cmd.Stdout
	output, err := cmd.Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	cmd = exec.Command("git", "branch", "-a")
	cmd.Dir = dirname
	cmd.Stderr = cmd.Stdout
	output, err = cmd.Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	branchesRaw := strings.Split(string(output), "\n")
	branches := []string{}

	for _, s := range branchesRaw {
		branch := strings.TrimSpace(s)
		if len(branch) == 0 {
			continue
		}
		if strings.Contains(branch, " -> ") {
			continue
		}
		branch = strings.ReplaceAll(branch, "* ", "")

		branches = append(branches, branch)
	}

	return c.JSON(
		http.StatusOK,
		branches,
	)
}

func (a *Controller) setBuildBranch(c echo.Context) error {
	var dirname string
	s := a.settings.GetSetting("BUILD_LOCATION")
	if s.ID != 0 {
		dirname = s.Value
	}

	cmd := exec.Command("git", "fetch", "origin")
	cmd.Dir = dirname
	cmd.Stderr = cmd.Stdout
	output, err := cmd.Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	cmd = exec.Command("git", "checkout", "-f", c.Param("branch"))
	cmd.Dir = dirname
	cmd.Stderr = cmd.Stdout
	_, err = cmd.Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	cmd = exec.Command("git", "pull")
	cmd.Dir = dirname
	cmd.Stderr = cmd.Stdout
	output, err = cmd.Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{"branch": strings.TrimSpace(string(output))},
	)
}

func (a *Controller) getBuildCurrentBranch(c echo.Context) error {
	var dirname string
	s := a.settings.GetSetting("BUILD_LOCATION")
	if s.ID != 0 {
		dirname = s.Value
	}

	cmd := exec.Command("git", "fetch", "origin")
	cmd.Dir = dirname
	cmd.Stderr = cmd.Stdout
	output, err := cmd.Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	cmd = exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = dirname
	cmd.Stderr = cmd.Stdout
	output, err = cmd.Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(
		http.StatusOK,
		strings.TrimSpace(string(output)),
	)
}
