package eqemuserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/rfsnotify"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/Akkadius/spire/internal/websocket"
	"github.com/shirou/gopsutil/v3/process"
	"golang.org/x/sync/errgroup"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	zoneProcessName         = "zone"
	worldProcessName        = "world"
	ucsProcessName          = "ucs"
	queryServProcessName    = "queryserv"
	loginServerProcessName  = "loginserver"
	sharedMemoryProcessName = "shared_memory"
	processLoopTimer        = 1 * time.Second
)

// stopTimerMutex is used to lock the stop timer
// web requests are in goroutines and have the potential
// to write to the stop timer at the same time
var stopTimerMutex = sync.Mutex{}

type Launcher struct {
	logger          *logger.AppLogger
	serverconfig    *eqemuserverconfig.Config
	settings        *spire.Settings
	pathmgmt        *pathmgmt.PathManagement
	serverApi       *Client
	watcher         *rfsnotify.RWatcher
	websocketMgr    *websocket.ClientManager
	crashLogWatcher *CrashLogWatcher

	// properties
	configLastModified          time.Time
	isRunning                   bool // this is set by the server config
	runSharedMemory             bool
	runLoginserver              bool
	runQueryServ                bool
	runUcs                      bool
	updateOpcodesOnStart        bool
	deleteLogFilesOlderThanDays int
	patchesDirectory            string
	opcodesDirectory            string
	minZoneProcesses            int
	currentZoneDynamics         int
	currentZoneStatics          int
	lastRanLogTruncationTime    time.Time

	// distributed zone server mode
	isDistributedRoot    bool
	isLeafNode           bool
	nodes                []LauncherDistributedNode
	leafNodeDisconnected bool

	// counter properties
	staticZonesToBoot    []string
	currentOnlineStatics []string
	currentProcessCounts map[string]int
	stopTimer            int // timer in seconds to stop the server
	stopType             string
	bootedTotalDynamics  int
	targetDynamics       int

	processSleepTime     time.Duration
	zoneAssignedDynamics int

	// mutexes
	pollProcessMutex sync.Mutex
	nodesMutex       sync.Mutex
	configMutex      sync.Mutex
}

func NewLauncher(
	logger *logger.AppLogger,
	serverconfig *eqemuserverconfig.Config,
	settings *spire.Settings,
	pathmgmt *pathmgmt.PathManagement,
	serverApi *Client,
	websocketMgr *websocket.ClientManager,
	crashLogWatcher *CrashLogWatcher,

) *Launcher {
	l := &Launcher{
		logger:               logger,
		serverconfig:         serverconfig,
		settings:             settings,
		pathmgmt:             pathmgmt,
		currentProcessCounts: make(map[string]int),
		serverApi:            serverApi,
		stopTimer:            0,
		websocketMgr:         websocketMgr,
		processSleepTime:     1 * time.Second,
		crashLogWatcher:      crashLogWatcher,
	}

	return l
}

// Process runs the process loop for the launcher
// It only does anything if the server has the launcher set to run
func (l *Launcher) Process() {
	for {
		now := time.Now()
		stat, err := os.Stat(l.pathmgmt.GetEQEmuServerConfigFilePath())
		if err != nil {
			l.logger.Debug().
				Any("error", err.Error()).
				Any("isRunning", l.isRunning).
				Msg("Error getting server config file stat")
		}

		timeToRunTruncation := l.lastRanLogTruncationTime.IsZero() || time.Now().Sub(l.lastRanLogTruncationTime) > time.Hour
		if l.deleteLogFilesOlderThanDays > 0 && timeToRunTruncation {
			l.lastRanLogTruncationTime = time.Now()
			l.logger.Info().Any("days", l.deleteLogFilesOlderThanDays).Msg("Truncating logs older than configured days")
			l.truncateLogs()
		}

		l.logger.DebugVvv().
			Any("configStatTime", time.Now().Sub(now).String()).
			Any("configLastModified", l.configLastModified).
			Any("isRunning", l.isRunning).
			Msg("Main launcher process loop")

		if stat.ModTime().After(l.configLastModified) {
			l.configLastModified = stat.ModTime()
			l.loadServerConfig()
			l.logger.Debug().
				Any("configLastModified", l.configLastModified).
				Any("statTime", time.Now().Sub(now).String()).
				Any("isRunning", l.isRunning).
				Any("runSharedMemory", l.runSharedMemory).
				Any("runLoginserver", l.runLoginserver).
				Any("runQueryServ", l.runQueryServ).
				Any("minZoneProcesses", l.minZoneProcesses).
				Any("staticZones", l.staticZonesToBoot).
				Msg("Detected server config change")
		}

		if l.isRunning {
			err := l.Supervisor()
			if err != nil {
				fmt.Printf("Error running server launcher supervisor: %v\n", err)
			}
		}

		// variable sleep time based on the number of processes running
		if l.currentProcessCounts[zoneProcessName] > 0 {
			cnt := l.currentProcessCounts[zoneProcessName]
			if cnt > 1000 {
				l.processSleepTime = 10 * time.Second
			} else if cnt > 500 {
				l.processSleepTime = 5 * time.Second
			} else {
				l.processSleepTime = 1 * time.Second
			}
		}

		if l.isLeafNode {
			l.processSleepTime = time.Second * 10

			// if the leaf node is disconnected, we want to attempt registration more frequently
			if l.leafNodeDisconnected {
				l.processSleepTime = time.Second * 1
			}
		}

		time.Sleep(l.processSleepTime)
	}
}

// StartLauncherProcess starts the launcher process
func (l *Launcher) StartLauncherProcess() error {
	cfg, _ := l.serverconfig.Get()
	cfg.WebAdmin.Launcher.IsRunning = false // shut off legacy launcher in-case it's running
	cfg.Spire.LauncherStart = true
	err := l.serverconfig.Save(cfg)
	if err != nil {
		return err

	}

	err = l.checkIfLauncherIsRunning()
	if err != nil {
		return err
	}

	err = l.startLauncherProcess()
	if err != nil {
		return err
	}
	return nil
}

// Start starts the launcher
// Only call this from the launcher process itself
func (l *Launcher) Start() error {
	err := l.checkIfLauncherIsRunning()
	if err != nil {
		return err
	}

	l.loadServerConfig()

	cfg, _ := l.serverconfig.Get()

	l.logger.Info().
		Any("World name", cfg.Server.World.Longname).
		Msg("Starting server launcher")

	if env.IsAppModeCli() {
		l.logger.Info().Msg("Starting crash log watcher")
		go l.crashLogWatcher.Process()
	}

	if cfg.WebAdmin != nil && cfg.WebAdmin.Launcher != nil {
		if l.isLauncherDistributedModeLeaf() {
			l.isLeafNode = true
			l.logger.Info().Msg("Server detected is a leaf node")
			for {
				err := l.rpcClientRegister()
				if err != nil {
					l.logger.Error().Err(err).Msg("Error registering leaf node")
				} else {
					break
				}

				l.logger.Info().Msg("Retrying registration in 1 second(s)")

				time.Sleep(1 * time.Second)
			}
		}

		if l.isLauncherDistributedModeRoot() {
			l.isDistributedRoot = true
			l.logger.Info().Msg("Server detected is a distributed zone root")
		}

		if l.isDistributedRoot {
			hostname, _ := os.Hostname()
			cfg, _ := l.serverconfig.Get()

			l.nodes = append(
				l.nodes,
				LauncherDistributedNode{
					Address:          cfg.Server.World.Address,
					ConnectedAddress: "127.0.0.1",
					Hostname:         hostname,
					LastSeen:         time.Now(),
					NodeType:         LauncherNodeTypeRoot,
				},
			)
		}

		if l.isLeafNode || l.isDistributedRoot {
			// goroutine to start the rpc server
			go func() {
				err := l.StartRpcServer(3005)
				if err != nil {
					return
				}
			}()
		}
	}

	serverOnline := false
	uptime, err := l.serverApi.GetWorldUptime()
	if err == nil && len(uptime) > 0 {
		serverOnline = true
	}

	var g errgroup.Group

	// start shared memory if needed
	// this needs to be started and completed before the server processes
	if l.runSharedMemory {
		if !serverOnline {
			g.Go(func() error {
				l.logger.Info().Msg("Starting shared memory")
				err := l.startServerProcessSync(sharedMemoryProcessName)
				if err != nil {
					return err
				}

				return nil
			})
		} else {
			l.logger.Info().Msg("Server already online, skipping running shared memory")
		}
	}

	if l.updateOpcodesOnStart {
		g.Go(func() error {
			l.updatePatchFiles()

			return nil
		})
	}

	// run pre-boot operations in parallel
	if err := g.Wait(); err != nil {
		l.logger.Info().
			Err(err).
			Msg("Error starting server launcher")
		return err
	}

	cfg.Spire.LauncherStart = true
	err = l.serverconfig.Save(cfg)
	if err != nil {
		return err
	}

	return nil
}

// Restart stops and starts the launcher
func (l *Launcher) Restart() error {
	err := l.Stop()
	if err != nil {
		return err
	}

	time.Sleep(1 * time.Second)

	l.logger.Info().Msg("Restarting server launcher")

	return l.StartLauncherProcess()
}

// Stop stops the launcher
func (l *Launcher) Stop() error {
	l.logger.Info().Msg("Stopping server launcher")

	if l.stopTimer > 0 {
		l.logger.Info().Msgf("Stopping server in %v minute(s)", l.stopTimer/60)

		// message every minute
		timeToStop := time.Now().Add(time.Duration(l.stopTimer) * time.Second)

		l.MessageWorld(
			fmt.Sprintf(
				"[SERVER MESSAGE] The world will be coming down in [%v] minute(s), please log out before this time...",
				timeToStop.Sub(time.Now()).Round(time.Minute).Minutes(),
			),
		)

		stopTimer := l.stopTimer

		for {
			if l.stopTimer == 0 {
				l.logger.Info().Msg("Stop - Timer cancelled")
				return nil
			}

			if time.Now().After(timeToStop) {
				l.logger.Info().Msg("Stopping server after timed restart")
				l.MessageWorld("[SERVER MESSAGE] Server is shutting down now")
				l.emitStopTimer(time.Time{})
				break
			}

			if l.stopTimer != stopTimer {
				l.logger.Info().Msg("Stop - Timer changed")
				return nil
			}

			// message top of every minute
			diff := timeToStop.Sub(time.Now())
			if int(diff.Round(time.Second).Seconds())%60 == 0 {
				l.MessageWorld(
					fmt.Sprintf(
						"[SERVER MESSAGE] The world will be coming down in [%v] minute(s), please log out before this time...",
						timeToStop.Sub(time.Now()).Round(time.Minute).Minutes(),
					),
				)
			}

			l.emitStopTimer(timeToStop)

			time.Sleep(1 * time.Second)
		}
	}

	cfg, _ := l.serverconfig.Get()
	cfg.WebAdmin.Launcher.IsRunning = false // shut off legacy launcher in-case it's running
	cfg.Spire.LauncherStart = false
	err := l.serverconfig.Save(cfg)
	if err != nil {
		return err
	}

	// Usually invoked from the web interface
	// If the launcher is running, we need to send a shutdown action to the root node
	// This will stop all the servers in the distributed zone server setup
	if l.isLauncherDistributedModeRoot() {
		err := l.rpcClientRootNodeServerShutdown()
		if err != nil {
			l.logger.Error().Err(err).Msg("Error sending shutdown action to root node")
		}
	}

	// kill all server processes
	processes, _ := process.Processes()
	for _, p := range processes {
		proc := l.getProcessDetails(p)
		for _, s := range l.GetServerProcessNames() {
			if s == proc.BaseProcessName {
				l.logger.Debug().
					Any("pid", p.Pid).
					Any("baseProcessName", proc.BaseProcessName).
					Any("cmdline", proc.Cmdline).
					Msg("Stop - Killing server process")

				l.logger.Info().Any("pid", p.Pid).Any("process", proc.BaseProcessName).Msg("Killing server process")

				if err := p.Terminate(); err != nil {
					l.logger.Debug().
						Any("error", err.Error()).
						Any("pid", p.Pid).
						Msg("Error terminating process")
				}
			}
		}
	}

	// give the processes a second to terminate gracefully before killing
	time.Sleep(1 * time.Second)

	for _, p := range processes {
		for _, s := range l.GetServerProcessNames() {
			proc := l.getProcessDetails(p)
			if s == proc.BaseProcessName {
				err := p.Kill()
				if err != nil {
					l.logger.Debug().
						Any("error", err.Error()).
						Any("pid", p.Pid).
						Msg("Error killing process")
				}
			}

			// kill any instances of launchers
			isServerLauncher := strings.Contains(proc.Cmdline, "server-launcher") || strings.Contains(proc.Cmdline, "eqemu-server:launcher")
			if isServerLauncher && p.Pid != int32(os.Getpid()) {
				if err := p.Terminate(); err != nil {
					l.logger.Debug().
						Any("error", err.Error()).
						Any("pid", p.Pid).
						Msg("Error killing process")
				}
			}
		}
	}

	l.stopTimer = 0

	l.logger.Info().Msg("Stopped server launcher")

	return nil
}

// GetServerProcessNames returns a list of server process names
func (l *Launcher) GetServerProcessNames() []string {
	return []string{
		zoneProcessName,
		worldProcessName,
		ucsProcessName,
		queryServProcessName,
		loginServerProcessName,
	}
}

// Supervisor is the main process loop for the server launcher
// It will monitor the server processes and restart them if they crash
// this only runs when the server is set to run the launcher
func (l *Launcher) Supervisor() error {
	now := time.Now()

	l.pollProcessCounts()

	if l.isLeafNode {
		l.processDistributed()
		return nil
	}

	// boot world process if needed
	if l.currentProcessCounts[worldProcessName] == 0 {
		l.logger.Info().Msg("Starting World")
		err := l.startServerProcess(worldProcessName)
		if err != nil {
			return err
		}
	}

	// boot loginserver if needed
	if l.runLoginserver && l.currentProcessCounts[loginServerProcessName] == 0 {
		l.logger.Info().Msg("Starting Loginserver")
		err := l.startServerProcess(loginServerProcessName)
		if err != nil {
			return err
		}
	}

	// boot queryserv if needed
	if l.runQueryServ && l.currentProcessCounts[queryServProcessName] == 0 {
		l.logger.Info().Msg("Starting QueryServ")
		err := l.startServerProcess(queryServProcessName)
		if err != nil {
			return err
		}
	}

	// boot ucs if needed
	if l.runUcs && l.currentProcessCounts[ucsProcessName] == 0 {
		l.logger.Info().Msg("Starting UCS")
		err := l.startServerProcess(ucsProcessName)
		if err != nil {
			return err
		}
	}

	list, err := l.serverApi.GetZoneList()
	if err != nil {
		l.logger.Debug().Any("error", err.Error()).Msg("Error getting zone list")
		return nil
	}

	l.zoneAssignedDynamics = 0
	for _, z := range list.Data {
		if z.ZoneID == 0 {
			continue
		}
		if !z.IsStaticZone {
			l.zoneAssignedDynamics++
		}
	}

	// boot statics if needed
	if len(l.staticZonesToBoot) > 0 {
		var staticsToBoot []string
		for _, z := range l.staticZonesToBoot {
			isInList := false
			for _, cz := range l.currentOnlineStatics {
				if cz == z {
					isInList = true
					break
				}
			}

			// check if the zone is in staticsToBoot
			for _, cz := range staticsToBoot {
				if cz == z {
					isInList = true
					break
				}
			}

			if !isInList {
				err := l.startServerProcess(zoneProcessName, z)
				if err != nil {
					return err
				}
				staticsToBoot = append(staticsToBoot, z)
			}
		}

		if len(staticsToBoot) > 0 {
			l.logger.Info().Msgf("Starting Static Zones (%v) [%+v]", len(staticsToBoot), staticsToBoot)
		}

		l.logger.DebugVvv().
			Any("staticsToBoot", staticsToBoot).
			Msg("Supervisor - Booting static zone(s)")
	}

	l.bootedTotalDynamics = l.currentProcessCounts[zoneProcessName] - l.currentZoneStatics
	l.targetDynamics = l.zoneAssignedDynamics + l.minZoneProcesses

	// boot dynamics if needed
	if !l.isDistributedRoot {
		for l.currentProcessCounts[zoneProcessName]-l.currentZoneStatics < (l.zoneAssignedDynamics + l.minZoneProcesses) {
			// we don't want to start dynamic zone processes normally when distributed mode is enabled
			err := l.startServerProcess(zoneProcessName)
			if err != nil {
				return err
			}

			l.bootedTotalDynamics = l.currentProcessCounts[zoneProcessName] - l.currentZoneStatics
			l.targetDynamics = l.zoneAssignedDynamics + l.minZoneProcesses

			l.logger.Info().
				Any("bootedTotalDynamics", l.bootedTotalDynamics).
				Any("zoneAssignedDynamics", l.zoneAssignedDynamics).
				Any("minZoneProcesses", l.minZoneProcesses).
				Any("targetDynamics", l.targetDynamics).
				Msg("Starting Dynamic Zone")
		}
	} else {
		// distributed root
		l.processDistributed()
	}

	l.logger.DebugVv().
		Any("took", time.Now().Sub(now).String()).
		Any("currentProcessCounts", l.currentProcessCounts).
		Any("currentZoneDynamics", l.currentZoneDynamics).
		Any("currentZoneStatics", l.currentZoneStatics).
		Any("zoneAssignedDynamics", l.zoneAssignedDynamics).
		Any("statics - staticZonesToBoot", l.staticZonesToBoot).
		Any("statics - staticZonesToBoot (count)", len(l.staticZonesToBoot)).
		Any("statics - currentOnlineStatics (count)", len(l.currentOnlineStatics)).
		Any("statics - currentOnlineStatics", l.currentOnlineStatics).
		Msg("Supervisor - Loop complete")

	return nil
}

// startServerProcessSync starts a server process synchronously
func (l *Launcher) startServerProcessSync(name string, args ...string) error {
	bin, err := exec.LookPath(filepath.Join(l.pathmgmt.GetEQEmuServerPath(), "bin", name))
	if err != nil {
		return err
	}

	cmd := exec.Command(bin, args...)
	cmd.Dir = l.pathmgmt.GetEQEmuServerPath()
	if err := cmd.Start(); err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

// loadServerConfig loads the server config
// this is called on startup and when the server config changes
func (l *Launcher) loadServerConfig() {
	l.configMutex.Lock()
	defer l.configMutex.Unlock()

	cfg, _ := l.serverconfig.Get()
	l.isRunning = cfg.Spire.LauncherStart
	l.runSharedMemory = cfg.WebAdmin.Launcher.RunSharedMemory
	l.runLoginserver = cfg.WebAdmin.Launcher.RunLoginserver
	l.runQueryServ = cfg.WebAdmin.Launcher.RunQueryServ
	l.runUcs = cfg.WebAdmin.Launcher.RunUcs
	l.minZoneProcesses = cfg.WebAdmin.Launcher.MinZoneProcesses
	l.updateOpcodesOnStart = cfg.WebAdmin.Launcher.UpdateOpcodesOnStart
	l.deleteLogFilesOlderThanDays = cfg.WebAdmin.Launcher.DeleteLogFilesOlderThanDays

	if l.minZoneProcesses < 1 {
		l.minZoneProcesses = 5
	}

	if cfg.Server.Directories.Patches != "" {
		l.patchesDirectory = filepath.Join(
			l.pathmgmt.GetEQEmuServerPath(),
			cfg.Server.Directories.Patches,
		)
	}

	if cfg.Server.Directories.Opcodes != "" {
		l.opcodesDirectory = filepath.Join(
			l.pathmgmt.GetEQEmuServerPath(),
			cfg.Server.Directories.Opcodes,
		)
	}

	if cfg.WebAdmin.Launcher.DeleteLogFilesOlderThanDays == 0 {
		cfg.WebAdmin.Launcher.DeleteLogFilesOlderThanDays = 7
		l.deleteLogFilesOlderThanDays = 7
		// save
		_ = l.serverconfig.Save(cfg)
	}

	var staticZonesToBoot []string

	// filter duplicate definitions
	for _, z := range strings.Split(cfg.WebAdmin.Launcher.StaticZones, ",") {
		isInList := false
		for _, cz := range staticZonesToBoot {
			if cz == z {
				isInList = true
				break
			}
		}

		if !isInList {
			if len(z) > 0 {
				staticZonesToBoot = append(staticZonesToBoot, z)
			}
		}
	}

	l.staticZonesToBoot = staticZonesToBoot
}

// SetStopTimer sets the stop timer
func (l *Launcher) SetStopTimer(timer int) {
	stopTimerMutex.Lock()
	l.stopTimer = timer
	stopTimerMutex.Unlock()
}

func (l *Launcher) GetStopTimer() int {
	stopTimerMutex.Lock()
	defer stopTimerMutex.Unlock()

	return l.stopTimer
}

func (l *Launcher) SetStopTypeStopping() {
	l.stopType = "stopping"
}

func (l *Launcher) SetStopTypeRestarting() {
	l.stopType = "restarting"
}

// StopCancel stops the server cancel
func (l *Launcher) StopCancel() error {
	stopTimerMutex.Lock()
	l.stopTimer = 0
	stopTimerMutex.Unlock()

	l.emitStopTimer(time.Time{})

	l.MessageWorld("[SERVER MESSAGE] Server shutdown cancelled")

	return nil
}

// ServerProcessLauncherWatchdog is a watchdog that will monitor the server process launcher
func (l *Launcher) ServerProcessLauncherWatchdog() {
	go func() {
		l.logger.Info().Msg("Starting server process launcher watchdog")
		for {
			stat, err := os.Stat(l.pathmgmt.GetEQEmuServerConfigFilePath())
			if err != nil {
				l.logger.Debug().
					Any("error", err.Error()).
					Any("isRunning", l.isRunning).
					Msg("Error getting server config file stat")
			} else {
				if stat.ModTime().After(l.configLastModified) {
					l.configLastModified = stat.ModTime()
					l.loadServerConfig()
					l.logger.Debug().Msg("Watchdog - Detected server config change")
				}
			}

			if l.isRunning {
				isLauncherRunning := false
				processes, _ := process.Processes()
				for _, p := range processes {
					cmdline, err := p.Cmdline()
					if err != nil {
						if strings.Contains(err.Error(), "no such file or directory") {
							continue
						}

						l.logger.Debug().
							Any("error", err.Error()).
							Any("pid", p.Pid).
							Msg("Error getting process command line")
					}

					if strings.Contains(cmdline, "eqemu-server:launcher start") {
						isLauncherRunning = true
						break
					}
				}

				if !isLauncherRunning {
					l.logger.Debug().
						Msg("Launcher process not running - starting")

					err := l.startLauncherProcess()
					if err != nil {
						l.logger.Debug().
							Any("error", err.Error()).
							Msg("Error starting launcher process")
					}
				}
			}

			time.Sleep(10 * time.Second)
		}
	}()
}

// checkIfLauncherIsRunning checks if the launcher is already running
func (l *Launcher) checkIfLauncherIsRunning() error {
	// check if the launcher is already running
	processes, _ := process.Processes()
	for _, p := range processes {
		cmdline, err := p.Cmdline()
		if err != nil {
			if strings.Contains(err.Error(), "no such file or directory") {
				continue
			}

			l.logger.Debug().
				Any("error", err.Error()).
				Any("pid", p.Pid).
				Msg("Error getting process command line")
		}

		isRanFromGoRun := strings.Contains(cmdline, "go run")
		if strings.Contains(cmdline, "eqemu-server:launcher start") && p.Pid != int32(os.Getpid()) && !isRanFromGoRun {
			isRanFromAirWatcher := strings.Contains(cmdline, "air")
			isRanFromAir := strings.Contains(cmdline, "/tmp/bin/app")
			if isRanFromAirWatcher || isRanFromAir {
				continue
			}

			l.logger.Debug().
				Any("os.Getpid()", os.Getpid()).
				Any("pid", p.Pid).
				Any("cmdline", cmdline).
				Msg("Launcher process already running")

			fmt.Println("Launcher process already running")

			return fmt.Errorf("launcher process already running")
		}
	}
	return nil
}

type ProcessDetails struct {
	Pid             int32
	Cmdline         string
	Cwd             string
	Exe             string
	BaseProcessName string
}

// getProcessDetails returns the process details
func (l *Launcher) getProcessDetails(p *process.Process) ProcessDetails {
	cmdline, err := p.Cmdline()
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return ProcessDetails{}
		}

		l.logger.Debug().
			Any("error", err.Error()).
			Any("pid", p.Pid).
			Msg("Error getting process command line")
	}
	if runtime.GOOS == "windows" {
		cmdline = strings.ReplaceAll(cmdline, ".exe", "")
	}

	cwd, err := p.Cwd()
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return ProcessDetails{}
		}
		if strings.Contains(err.Error(), "permission denied") {
			return ProcessDetails{}
		}
		l.logger.Debug().
			Any("error", err.Error()).
			Any("pid", p.Pid).
			Msg("Error getting process cwd")
	}
	if runtime.GOOS == "windows" {
		cwd = strings.ReplaceAll(cwd, ".exe", "")
	}

	exe, err := p.Exe()
	if err != nil {
		l.logger.Debug().
			Any("error", err.Error()).
			Any("pid", p.Pid).
			Msg("Error getting process exe")
	}
	if runtime.GOOS == "windows" {
		exe = strings.ReplaceAll(exe, ".exe", "")
	}

	baseProcessName := filepath.Base(exe)
	if runtime.GOOS == "windows" {
		baseProcessName = strings.ReplaceAll(baseProcessName, ".exe", "")
	}

	// remove (deleted) from the process name
	// not sure why this shows up in this format
	baseProcessName = strings.ReplaceAll(baseProcessName, " (deleted)", "")

	return ProcessDetails{
		Pid:             p.Pid,
		Cmdline:         cmdline,
		Cwd:             cwd,
		Exe:             exe,
		BaseProcessName: baseProcessName,
	}
}

// updatePatchFiles will download the patch files from github
func (l *Launcher) updatePatchFiles() {
	l.logger.Info().Msg("Updating patches (opcodes)")

	// get the patch files
	patchFiles := []string{
		"patch_RoF.conf",
		"patch_RoF2.conf",
		"patch_SoD.conf",
		"patch_SoF.conf",
		"patch_Titanium.conf",
		"patch_UF.conf",
	}

	// get the opcode files
	opcodeFiles := []string{
		"opcodes.conf",
		"mail_opcodes.conf",
	}

	now := time.Now()

	// download the patch files in errgroup
	var g errgroup.Group
	for _, p := range patchFiles {
		url := fmt.Sprintf("https://raw.githubusercontent.com/EQEmu/Server/master/utils/patches/%s", p)
		path := filepath.Join(l.patchesDirectory, p)
		relative := strings.ReplaceAll(path, l.pathmgmt.GetEQEmuServerPath()+string(filepath.Separator), "")

		g.Go(func() error {
			if err := downloadFile(url, path); err != nil {
				l.logger.Debug().Msg(fmt.Sprintf("Failed to download %s: %v", url, err))
				return err
			}
			l.logger.Debug().Msg(fmt.Sprintf("Successfully downloaded %s", url))
			l.logger.Info().Any("file", relative).Msg("Saved opcodes")
			return nil
		})
	}

	for _, o := range opcodeFiles {
		url := fmt.Sprintf("https://raw.githubusercontent.com/EQEmu/Server/master/utils/patches/%s", o)
		path := filepath.Join(l.opcodesDirectory, o)
		relative := strings.ReplaceAll(path, l.pathmgmt.GetEQEmuServerPath()+string(filepath.Separator), "")

		g.Go(func() error {
			if err := downloadFile(url, path); err != nil {
				l.logger.Debug().Msg(fmt.Sprintf("Failed to download %s: %v", url, err))
				return err
			}
			l.logger.Debug().Msg(fmt.Sprintf("Successfully downloaded %s", url))
			l.logger.Info().Any("file", relative).Msg("Saved opcodes")
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		l.logger.Warn().Err(err).Msg("Failed to download patch files from github")
	} else {
		l.logger.Info().Any("took", time.Since(now).String()).Msg("Updated patch files")
	}
}

// truncateLogs will truncate log files older than the specified days
func (l *Launcher) truncateLogs() {
	logsDir := l.pathmgmt.GetLogsDirPath()
	_ = filepath.Walk(logsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			l.logger.Debug().Any("error", err.Error()).Any("path", path).Msg("Error walking logs")
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, ".log") {
			return nil
		}

		if time.Since(info.ModTime()).Hours() > float64(l.deleteLogFilesOlderThanDays*24) {
			l.logger.Info().Any("path", path).Msg("Truncating log file")
			err := os.Remove(path)
			if err != nil {
				l.logger.Debug().Any("error", err.Error()).
					Any("path", path).
					Any("days", l.deleteLogFilesOlderThanDays).
					Msg("Error truncating log file")
			}
		}

		return nil
	})
}

// MessageWorld sends a message to the world
func (l *Launcher) MessageWorld(message string) {
	l.logger.Info().Any("message", message).Msg("Sent message to world")
	l.websocketMgr.Broadcast("notification", message)
	_ = l.serverApi.MessageWorld(message)
}

// emitStopTimer emits the stop timer to the websocket
func (l *Launcher) emitStopTimer(stopTime time.Time) {
	type StopTimer struct {
		Time int64  `json:"time"`
		Type string `json:"type"`
	}

	data := StopTimer{
		Time: stopTime.Unix(),
		Type: l.stopType,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		l.logger.Debug().Any("error", err.Error()).Msg("Error marshalling stop timer data")
		return
	}

	l.websocketMgr.Broadcast("stopTimer", string(jsonData))
}

type ZoneServerClient struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Race  int    `json:"race"`
	Class int    `json:"class"`
	Guild int    `json:"guild"`
}

type ZoneServer struct {
	ID                int                `json:"id"`
	ZoneID            int                `json:"zone_id"`
	ZoneName          string             `json:"zone_name"`
	ZoneLongName      string             `json:"zone_long_name"`
	ZoneOsPid         int                `json:"zone_os_pid"`
	IsStaticZone      bool               `json:"is_static_zone"`
	ClientPort        int                `json:"client_port"`
	NumberPlayers     int                `json:"number_players"`
	InstanceID        int                `json:"instance_id,omitempty"`
	Clients           []ZoneServerClient `json:"clients"`
	ZoneServerAddress string             `json:"zone_server_address"`

	// set by the launcher
	ConnectedAddress  string `json:"connected_address"`
	ConfiguredAddress string `json:"configured_address"`

	// process info
	Pid     int32   `json:"pid"`
	Name    string  `json:"name"`
	CmdLine string  `json:"cmd"`
	Cpu     float64 `json:"cpu"`
	Memory  uint64  `json:"memory"`
	Elapsed int64   `json:"elapsed"` // uptime
}

func (l *Launcher) GetZoneserverList() ([]ZoneServer, error) {
	zones, err := l.serverApi.GetZoneList()
	if err != nil {
		return []ZoneServer{}, err
	}

	// Fetch client list
	clients, err := l.serverApi.GetWorldClientList()
	if err != nil {
		return []ZoneServer{}, err
	}

	// Create a map of zone ID to clients for quick lookup
	clientMap := make(map[int][]ZoneServerClient, 0)
	for _, client := range clients.Data {
		zoneID := client.Server.ID
		clientInfo := ZoneServerClient{
			ID:    client.ID,
			Name:  client.Name,
			Race:  client.Race,
			Class: client.Class,
			Guild: client.GuildID,
		}
		clientMap[zoneID] = append(clientMap[zoneID], clientInfo)
	}

	processes, _ := process.Processes()

	cfg, _ := l.serverconfig.Get()

	// Combine zone data with associated clients
	var combinedData []ZoneServer
	for _, zone := range zones.Data {
		clients := clientMap[zone.ID]
		if clients == nil {
			clients = []ZoneServerClient{}
		}

		zoneInfo := ZoneServer{
			ID:                zone.ID,
			ZoneID:            zone.ZoneID,
			ZoneName:          zone.ZoneName,
			ZoneLongName:      zone.ZoneLongName,
			ZoneOsPid:         zone.ZoneOsPid,
			IsStaticZone:      zone.IsStaticZone,
			ClientPort:        zone.ClientPort,
			NumberPlayers:     zone.NumberPlayers,
			InstanceID:        zone.InstanceID,
			Clients:           clients,
			ZoneServerAddress: zone.ClientAddress,
			ConnectedAddress:  zone.IP,
		}

		combinedData = append(combinedData, zoneInfo)
	}

	for _, p := range processes {
		for i, zone := range combinedData {
			if zone.ZoneServerAddress == cfg.Server.World.Address && int(p.Pid) == zone.ZoneOsPid {
				name, _ := p.Name()
				cmdLine, _ := p.Cmdline()
				cpuPercent, _ := p.CPUPercent()
				uptime, _ := p.CreateTime()
				now := time.Now().Unix()
				combinedData[i].Pid = p.Pid
				combinedData[i].Name = name
				combinedData[i].CmdLine = cmdLine
				combinedData[i].Cpu = cpuPercent
				combinedData[i].Memory = getMemoryUsage(p)
				combinedData[i].Elapsed = now - (uptime / 1000)
			}
		}
	}

	return combinedData, nil
}

// processDistributed processes launcher for both root and leaf nodes
func (l *Launcher) processDistributed() {
	//
	//defer l.nodesMutex.Unlock()

	l.logger.Info().
		//Any("isDistributedRoot", l.isDistributedRoot).
		//Any("isLeafNode", l.isLeafNode).
		Any("bootedTotalDynamics", l.bootedTotalDynamics).
		Msg("Processing distributed")

	if l.isDistributedRoot {
		totalZoneProcesses := 0
		var indicesToRemove []int
		l.nodesMutex.Lock()

		if len(l.nodes) == 1 {
			l.logger.Info().Msg("Waiting for leaf nodes to connect before starting zones")
			l.nodesMutex.Unlock()
			return
		}

		for i, node := range l.nodes {
			r, err := l.rpcClientGetZoneCount(node)
			if err != nil {
				l.logger.Error().Err(err).Msg("Error getting zone count from node, removing node from list")
				indicesToRemove = append(indicesToRemove, i)
				continue
			}

			l.nodes[i].CurrentZoneCount = r.ZoneCount
			l.nodes[i].TargetZoneCount = r.ZoneCount
			l.nodes[i].MaxZoneCount = r.MaxZoneCount
			totalZoneProcesses += r.ZoneCount

			l.logger.Info().
				Any("node", node.Hostname).
				Any("address", node.Address).
				Any("ConnectedAddress", node.ConnectedAddress).
				Any("zoneCount", r.ZoneCount).
				Any("maxZoneCount", r.MaxZoneCount).
				Msg("Processing node, got zone count")
		}

		// Remove nodes in reverse order to avoid index shifting
		for i := len(indicesToRemove) - 1; i >= 0; i-- {
			index := indicesToRemove[i]
			l.logger.Info().
				Any("node", l.nodes[index].Hostname).
				Any("address", l.nodes[index].Address).
				Msg("Removing node from list")
			l.nodes = append(l.nodes[:index], l.nodes[index+1:]...)
		}
		l.nodesMutex.Unlock()

		delta := totalZoneProcesses - l.zoneAssignedDynamics
		l.logger.DebugVvv().
			Any("delta", delta).
			Any("minZoneProcesses", l.minZoneProcesses).
			Msg("Processing distributed")
		if delta < l.minZoneProcesses {
			zonesToBoot := l.minZoneProcesses - delta

			l.logger.Info().
				Any("totalZones", totalZoneProcesses).
				Any("zoneAssignedDynamics", l.zoneAssignedDynamics).
				Any("bootedTotalDynamics", l.bootedTotalDynamics).
				Any("targetDynamics", l.targetDynamics).
				Any("minZoneProcesses", l.minZoneProcesses).
				Any("zonesToBoot", zonesToBoot).
				Msg("Processing distributed")

			// distribute the zones to the nodes
			// sort the nodes by the number of zones they are running, least to most
			l.nodesMutex.Lock()
			sort.Slice(l.nodes, func(i, j int) bool {
				return l.nodes[i].CurrentZoneCount < l.nodes[j].CurrentZoneCount
			})

			// reset
			for i, _ := range l.nodes {
				l.nodes[i].AtMaxZoneCount = false
			}

			// loop through the nodes and increment the target zone count
			// we need to make sure we don't exceed the max zones per node
			iterations := 0
			for zonesToBoot > 0 {
				for i, _ := range l.nodes {

					// sanity check, should never happen
					if iterations > 10 {
						l.logger.Error().Msg("Failed to distribute zones, too many iterations")
						zonesToBoot = 0
						break
					}

					// if the node is at max zone count, skip
					if l.nodes[i].AtMaxZoneCount {
						// if the node is at max zone count, skip
						iterations++
						continue
					}

					// if the node is at max zone count, skip
					atMaxZoneCount := l.nodes[i].CurrentZoneCount >= l.nodes[i].MaxZoneCount || l.nodes[i].TargetZoneCount >= l.nodes[i].MaxZoneCount
					if l.nodes[i].MaxZoneCount > 0 && atMaxZoneCount {
						l.logger.Info().
							Any("node", l.nodes[i].Hostname).
							Any("address", l.nodes[i].Address).
							Any("CurrentZoneCount", l.nodes[i].CurrentZoneCount).
							Any("MaxZoneCount", l.nodes[i].MaxZoneCount).
							Msg("Node at max zone count")

						l.nodes[i].AtMaxZoneCount = true

						// if we only have one node, we can't boot any more zones if it's at max
						if len(l.nodes) == 1 {
							l.logger.Error().Msg("Only one node and it's at max zone count, can't distribute zones")
							zonesToBoot = 0
							break
						}

						// keep track if each node is at max zone count
						iterations++

						continue
					}

					l.nodes[i].TargetZoneCount++
					zonesToBoot--

					// if we have no more zones to boot, break
					if zonesToBoot == 0 {
						break
					}
				}
			}
			l.nodesMutex.Unlock()

			// check to see if all nodes are at max zone count
			allAtMaxZoneCount := true
			for _, node := range l.nodes {
				if !node.AtMaxZoneCount {
					allAtMaxZoneCount = false
					break
				}
			}

			// if all nodes are at max zone count, we can't distribute zones
			if allAtMaxZoneCount {
				l.logger.Error().Err(errors.New("all nodes at max zone count")).Msg("Can't distribute zones")
				return
			}

			// send the target zone count to each node
			for _, node := range l.nodes {
				if node.TargetZoneCount == 0 {
					continue
				}

				err := l.rpcClientSetTargetZoneCount(node, node.TargetZoneCount)
				if err != nil {
					l.logger.Error().
						Err(err).
						Any("node", node.Hostname).
						Any("address", node.Address).
						Msg("Error launching zones on node")
				}

				l.logger.Info().
					Any("node", node.Hostname).
					Any("CurrentZoneCount", node.CurrentZoneCount).
					Any("targetZoneCount", node.TargetZoneCount).
					Msg("Processing node")
			}
		}
	}

	if l.isLeafNode {
		err := l.rpcClientRegister()
		if err != nil {
			l.logger.Error().Err(err).Msg("Error registering with root node")
		}
	}
}

func (l *Launcher) pollProcessCounts() {
	l.pollProcessMutex.Lock()
	defer l.pollProcessMutex.Unlock()

	// reset
	l.currentProcessCounts = make(map[string]int)
	l.currentOnlineStatics = []string{}
	l.currentZoneDynamics = 0
	l.currentZoneStatics = 0

	processes, _ := process.Processes()
	for _, p := range processes {
		proc := l.getProcessDetails(p)
		for _, s := range l.GetServerProcessNames() {
			if s == proc.BaseProcessName {
				l.logger.DebugVv().
					Any("pid", p.Pid).
					Any("baseProcessName", proc.BaseProcessName).
					Any("cmdline", proc.Cmdline).
					Msg("Supervisor - Found server process")

				if _, ok := l.currentProcessCounts[s]; !ok {
					l.currentProcessCounts[s] = 0
				}

				if s == zoneProcessName {
					// cmdline path can contain spaces which will break the split
					cmdline := proc.Cmdline
					cmdline = strings.ReplaceAll(cmdline, proc.Cwd, "")
					cmdline = strings.TrimSpace(cmdline)

					// get arg
					// check if it's in the static zones
					// if it is, add it to the current online statics
					arg := strings.Split(cmdline, " ")
					if len(arg) > 1 {
						for _, z := range l.staticZonesToBoot {
							if z == arg[1] {
								// make sure it's not already in the list
								isInList := false
								for _, cz := range l.currentOnlineStatics {
									if cz == z {
										isInList = true
										break
									}
								}
								if !isInList {
									l.currentOnlineStatics = append(l.currentOnlineStatics, z)
								}
							}
						}
						l.currentZoneStatics++
					} else {
						l.currentZoneDynamics++
					}
				}

				l.currentProcessCounts[s]++
			}
		}

		l.logger.DebugVvv().
			Any("pid", p.Pid).
			Any("cwd", proc.Cwd).
			Any("exe", proc.Exe).
			Any("cmdline", proc.Cmdline).
			Msg("pollProcessCounts")
	}

	l.bootedTotalDynamics = l.currentProcessCounts[zoneProcessName] - l.currentZoneStatics
}

// isLauncherDistributedModeRoot checks if the launcher is in distributed mode and is the root node
// loads from the server config
func (l *Launcher) isLauncherDistributedModeRoot() bool {
	cfg, _ := l.serverconfig.Get()

	if cfg.WebAdmin != nil && cfg.WebAdmin.Launcher != nil && cfg.WebAdmin.Launcher.DistributedNodeType == LauncherNodeTypeRoot {
		return true
	}

	return false
}

// isLauncherDistributedModeLeaf checks if the launcher is in distributed mode leaf
// loads from the server config
func (l *Launcher) isLauncherDistributedModeLeaf() bool {
	cfg, _ := l.serverconfig.Get()

	if cfg.WebAdmin != nil && cfg.WebAdmin.Launcher != nil && cfg.WebAdmin.Launcher.DistributedNodeType == LauncherNodeTypeLeaf {
		return true
	}

	return false
}
