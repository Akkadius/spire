package eqemuserver

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/labstack/echo/v4"
	"github.com/mholt/archiver/v4"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Controller struct {
	db             *database.DatabaseResolver
	logger         *logrus.Logger
	eqemuserverapi *Client
	pathmgmt       *pathmgmt.PathManagement
	settings       *spire.Settings
	serverconfig   *eqemuserverconfig.Config
	updater        *Updater
}

func NewController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	api *Client,
	serverconfig *eqemuserverconfig.Config,
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
		routes.RegisterRoute(http.MethodGet, "eqemuserver/dashboard-stats", a.getDashboardStats, nil),
		routes.RegisterRoute(http.MethodGet, "eqemuserver/manual-backup/:type", a.getManualBackup, nil),
		routes.RegisterRoute(http.MethodGet, "eqemuserver/logs", a.listLogFiles, nil),
		routes.RegisterRoute(http.MethodGet, "eqemuserver/log/:file", a.getFileLog, nil),
		routes.RegisterRoute(http.MethodDelete, "eqemuserver/log/:file", a.deleteFileLog, nil),
		routes.RegisterRoute(http.MethodGet, "eqemuserver/log-search/:search", a.logSearch, nil),
		routes.RegisterRoute(http.MethodGet, "eqemuserver/pre-flight/:process", a.preflight, nil),
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
	// validation: check if server is online
	online := false
	processes, _ := process.Processes()
	for _, p := range processes {
		cmdline, err := p.Cmdline()
		if err != nil {
			return err
		}
		if strings.Contains(cmdline, "world") {
			online = true
		}
	}

	if online {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": "You cannot issue an update while the server is online, please shut the server off before trying to update"},
		)
	}

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
		c.String(http.StatusOK, scanner.Text()+"\n")
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

type DashboardStatsResponse struct {
	Accounts   int64  `json:"accounts"`
	Characters int64  `json:"characters"`
	Guilds     int64  `json:"guilds"`
	Items      int64  `json:"items"`
	Npcs       int64  `json:"npcs"`
	Uptime     string `json:"uptime"`
}

func (a *Controller) getDashboardStats(c echo.Context) error {
	tables := []string{"account", "character_data", "guilds", "items", "npc_types"}
	counts := make(map[string]int64)
	for _, table := range tables {
		var count int64
		a.db.GetEqemuDb().Raw(fmt.Sprintf("select count(*) as count from %v", table)).Scan(&count)
		counts[table] = count
	}

	uptime, err := a.eqemuserverapi.GetWorldUptime()
	if err != nil {
		uptime = "Server offline"
	}

	r := DashboardStatsResponse{
		Accounts:   counts["account"],
		Characters: counts["character_data"],
		Guilds:     counts["guilds"],
		Items:      counts["items"],
		Npcs:       counts["npc_types"],
		Uptime:     uptime,
	}

	return c.JSON(http.StatusOK, r)
}

type DownloadType struct {
	path string
	name string
}

func (a *Controller) getManualBackup(c echo.Context) error {
	t := []DownloadType{
		{path: a.pathmgmt.GetQuestsDir(), name: "quests"},
		{path: a.pathmgmt.GetMapsDir(), name: "maps"},
	}

	requestedType := c.Param("type")
	foundExport := false
	var downloadType DownloadType
	for _, e := range t {
		if e.name == requestedType {
			foundExport = true
			downloadType = e
		}
	}

	if !foundExport {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invalid download type"})
	}

	f := make(map[string]string, 0)

	err := filepath.Walk(
		downloadType.path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if len(path) == 0 {
				return nil
			}

			if path == downloadType.path {
				return nil
			}

			// stat file
			fi, err := os.Stat(path)
			if err != nil {
				return nil
			}

			if fi.Mode().IsRegular() {
				replacePath := fmt.Sprintf("%v%v", downloadType.path, string(filepath.Separator))
				zipPath := strings.ReplaceAll(path, replacePath, "")
				zipPath = strings.ReplaceAll(zipPath, "\\", "/")
				f[path] = zipPath
			}

			return nil
		},
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// map files on disk to their paths in the archive
	files, err := archiver.FilesFromDisk(nil, f)
	if err != nil {
		return err
	}

	file := fmt.Sprintf("%v-%v.zip", downloadType.name, time.Now().Format("2006-01-02"))
	downloadFile := filepath.Join(os.TempDir(), file)

	_ = os.Remove(downloadFile)

	// create the output file we'll write to
	out, err := os.Create(downloadFile)
	if err != nil {
		return err
	}
	defer out.Close()

	// we can use the CompressedArchive type to gzip a tarball
	// (compression is not required; you could use Tar directly)
	format := archiver.Zip{}

	// create the archive
	err = format.Archive(context.Background(), out, files)
	if err != nil {
		return err
	}

	c.Response().Header().Add("Access-Control-Expose-Headers", "Content-Disposition")

	if _, err := os.Stat(downloadFile); !errors.Is(err, os.ErrNotExist) {
		return c.Attachment(downloadFile, filepath.Base(downloadFile))
	}

	return nil
}

type FileInfo struct {
	Name         string `json:"name"`          // base name of the file
	Path         string `json:"path"`          // relative path to file anchored from walk path
	Size         int64  `json:"size"`          // length in bytes for regular files; system-dependent for others
	Mode         int    `json:"mode"`          // file mode bits
	ModifiedTime int64  `json:"modified_time"` // modification time
	IsDirectory  bool   `json:"is_directory"`  // abbreviation for Mode().IsDir()
}

func (a *Controller) listLogFiles(c echo.Context) error {
	var files []FileInfo
	err := filepath.Walk(a.pathmgmt.GetLogsDirPath(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(info.Name()) == ".log" {
			newPath := strings.ReplaceAll(path, a.pathmgmt.GetLogsDirPath()+string(filepath.Separator), "")
			files = append(files, FileInfo{
				Name:         info.Name(),
				Path:         newPath,
				Size:         info.Size(),
				Mode:         int(info.Mode()),
				ModifiedTime: info.ModTime().Unix(),
				IsDirectory:  info.IsDir(),
			})
			return nil
		}
		return nil
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].ModifiedTime > files[j].ModifiedTime
	})

	return c.JSON(http.StatusOK, files)
}

type FileReadResponse struct {
	Contents string `json:"contents"`
	Cursor   int    `json:"cursor"`
}

func (a *Controller) getFileLog(c echo.Context) error {
	logFile := filepath.Join(a.pathmgmt.GetLogsDirPath(), c.Param("file"))
	if !strings.Contains(logFile, a.pathmgmt.GetLogsDirPath()) {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invalid access!"})
	}

	stat, err := os.Stat(logFile)
	if os.IsNotExist(err) {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "File does not exist!"})
	}

	f, err := os.Open(logFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// read from cursor
	var contents []byte
	bufferSize := stat.Size()
	if len(c.QueryParam("cursor")) > 0 {
		cursorIn, err := strconv.ParseInt(c.QueryParam("cursor"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		bufferSize = stat.Size() - cursorIn
		bytes := make([]byte, bufferSize)
		cursor, err := f.ReadAt(bytes, cursorIn)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		contents = bytes[:cursor]

		return c.JSON(http.StatusOK, FileReadResponse{
			Contents: string(contents),
			Cursor:   cursor,
		})
	}

	// read whole file, initial read
	bytes := make([]byte, bufferSize)
	cursor, err := f.Read(bytes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	contents = bytes[:cursor]

	return c.JSON(http.StatusOK, FileReadResponse{
		Contents: string(contents),
		Cursor:   cursor,
	})
}

func (a *Controller) deleteFileLog(c echo.Context) error {
	logFile := filepath.Join(a.pathmgmt.GetLogsDirPath(), c.Param("file"))
	if !strings.Contains(logFile, a.pathmgmt.GetLogsDirPath()) {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invalid access!"})
	}

	_, err := os.Stat(logFile)
	if os.IsNotExist(err) {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "File does not exist!"})
	}

	if err := os.Remove(logFile); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Deleted successfully!"})
}

type Line struct {
	Line   string `json:"line"`
	Number int    `json:"line_number"`
}

type LogSearchResult struct {
	File  string `json:"file"`
	Lines []Line `json:"lines"`
}

const lineMatchLimit = 100

func (a *Controller) logSearch(c echo.Context) error {
	search := strings.ToLower(c.Param("search"))

	var results []LogSearchResult
	err := filepath.Walk(a.pathmgmt.GetLogsDirPath(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(info.Name()) == ".log" {
			b, err := os.ReadFile(path) // just pass the file name
			if err != nil {
				fmt.Print(err)
			}

			var lines []Line
			content := string(b)

			// check the file first before attempting to search line by line
			if !strings.Contains(strings.ToLower(content), search) {
				return nil
			}

			line := 1
			for _, s := range strings.Split(content, "\n") {
				if strings.Contains(strings.ToLower(s), search) {
					lines = append(lines, Line{
						Line:   s,
						Number: line,
					})

					if len(lines) > lineMatchLimit {
						break
					}
				}
				line++
			}

			newPath := strings.ReplaceAll(path, a.pathmgmt.GetLogsDirPath()+string(filepath.Separator), "")

			results = append(results, LogSearchResult{
				File:  newPath,
				Lines: lines,
			})

			return nil
		}
		return nil
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

type PreflightCheck struct {
	process   string
	assertion string
	asserted  bool
	timeout   int
	args      []string
}

func (a *Controller) preflight(c echo.Context) error {
	checks := []PreflightCheck{
		{process: "world", assertion: "listener started on port", timeout: 120},
		{process: "zone", assertion: "Zone bootup type", timeout: 5, args: []string{"soldungb"}},
		{process: "shared_memory", assertion: "Loading base data", timeout: 5},
		{process: "ucs", assertion: "LoadChatChannels", timeout: 5},
		{process: "loginserver", assertion: "Server Started", timeout: 5},
	}

	p := strings.ReplaceAll(strings.ToLower(c.Param("process")), " ", "_")

	var check = PreflightCheck{}
	for _, ch := range checks {
		if ch.process == p {
			check = ch
		}
	}

	if len(check.assertion) == 0 {
		return c.String(http.StatusBadRequest, "Invalid check type!")
	}

	c.Response().WriteHeader(http.StatusOK)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(check.timeout)*time.Second)
	defer cancel()

	bin := filepath.Join(a.pathmgmt.GetEQEmuServerBinPath(), check.process)
	cmd := exec.CommandContext(ctx, bin, check.args...)
	cmd.Env = os.Environ()
	if runtime.GOOS == "linux" {
		cmd.Env = append(cmd.Env, "IS_TTY=true")
	}
	cmd.Dir = a.pathmgmt.GetEQEmuServerPath()
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
		if strings.Contains(stripAnsi(scanner.Text()), check.assertion) && !check.asserted {
			check.asserted = true
			cmd.Process.Kill()
		}

		c.String(http.StatusOK, scanner.Text()+"\n")
		c.Response().Flush()
	}

	err = cmd.Wait()
	if err != nil {
		if strings.Contains(err.Error(), "signal: killed") {
			return c.String(http.StatusOK, "Completed pre-flight checks successfully\n")
		}

		return c.String(http.StatusBadRequest, err.Error())
	}

	c.Response().Flush()

	return nil
}
