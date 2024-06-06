package eqemuserver

import (
	"fmt"
	"github.com/Akkadius/spire/internal/discord"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/rfsnotify"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/fsnotify/fsnotify"
	"github.com/shirou/gopsutil/v3/process"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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
	logger       *logger.AppLogger
	serverconfig *eqemuserverconfig.Config
	settings     *spire.Settings
	pathmgmt     *pathmgmt.PathManagement
	serverApi    *Client
	watcher      *rfsnotify.RWatcher

	// properties
	configLastModified          time.Time
	isRunning                   bool // this is set by the server config
	runSharedMemory             bool
	runLoginserver              bool
	runQueryServ                bool
	updateOpcodesOnStart        bool
	deleteLogFilesOlderThanDays int
	patchesDirectory            string
	opcodesDirectory            string
	minZoneProcesses            int
	currentZoneDynamics         int
	currentZoneStatics          int
	lastRanLogTruncationTime    time.Time

	staticZonesToBoot    []string
	currentOnlineStatics []string
	currentProcessCounts map[string]int
	stopTimer            int // timer in seconds to stop the server

	// meta properties
	watchCrashLogs bool
	discordWebhook string
	serverLongName string
}

func NewLauncher(
	logger *logger.AppLogger,
	serverconfig *eqemuserverconfig.Config,
	settings *spire.Settings,
	pathmgmt *pathmgmt.PathManagement,
	serverApi *Client,
) *Launcher {
	l := &Launcher{
		logger:               logger,
		serverconfig:         serverconfig,
		settings:             settings,
		pathmgmt:             pathmgmt,
		currentProcessCounts: make(map[string]int),
		serverApi:            serverApi,
		stopTimer:            0,
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

		if l.watchCrashLogs && l.watcher == nil {
			l.startCrashLogWatcher()
		}

		if l.isRunning {
			err := l.Supervisor()
			if err != nil {
				fmt.Printf("Error running server launcher supervisor: %v\n", err)
			}
		}

		time.Sleep(processLoopTimer)
	}
}

// StartLauncherProcess starts the launcher process
func (l *Launcher) StartLauncherProcess() error {
	cfg := l.serverconfig.Get()
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
	l.logger.Info().Msg("Starting server launcher")

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
		l.logger.Info().Err(err).Msg("Error starting server launcher")
		return err
	}

	cfg := l.serverconfig.Get()
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

		_ = l.serverApi.MessageWorld(
			fmt.Sprintf(
				"[SERVER MESSAGE] The world will be coming down in [%v] minute(s), please log out before this time...",
				timeToStop.Sub(time.Now()).Round(time.Minute).Minutes(),
			),
		)

		for {
			if l.stopTimer == 0 {
				l.logger.Info().Msg("Stop - Timer cancelled")
				return nil
			}

			if time.Now().After(timeToStop) {
				l.logger.Info().Msg("Stopping server after timed restart")
				break
			}

			if time.Now().Second() == 0 {
				l.logger.Info().Msgf("Server will stop in %v minute(s)", timeToStop.Sub(time.Now()).Minutes())

				_ = l.serverApi.MessageWorld(
					fmt.Sprintf(
						"[SERVER MESSAGE] The world will be coming down in [%v] minute(s), please log out before this time...",
						timeToStop.Sub(time.Now()).Round(time.Minute).Minutes(),
					),
				)
			}

			time.Sleep(1 * time.Second)
		}
	}

	cfg := l.serverconfig.Get()
	cfg.WebAdmin.Launcher.IsRunning = false // shut off legacy launcher in-case it's running
	cfg.Spire.LauncherStart = false
	err := l.serverconfig.Save(cfg)
	if err != nil {
		return err
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

		for _, s := range l.GetServerProcessNames() {
			if s == proc.BaseProcessName {
				err := p.Kill()
				if err != nil {
					l.logger.Debug().
						Any("error", err.Error()).
						Any("pid", p.Pid).
						Msg("Error killing process")
				}
			}
		}

		l.logger.DebugVvv().
			Any("pid", p.Pid).
			Any("exe", proc.Exe).
			Any("cmdline", proc.Cmdline).
			Any("baseProcessName", proc.BaseProcessName).
			Msg("Stop - Checking process")

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
					// get arg
					// check if it's in the static zones
					// if it is, add it to the current online statics
					arg := strings.Split(proc.Cmdline, " ")
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
			Msg("Supervisor - Checking process")
	}

	// boot world if needed
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
	if l.currentProcessCounts[ucsProcessName] == 0 {
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

	zoneAssignedDynamics := 0
	zoneIdleDynamics := 0
	for _, z := range list.Data {
		if !z.IsStaticZone {
			if z.ZoneID == 0 {
				zoneIdleDynamics++
				continue
			}
			zoneAssignedDynamics++
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

		l.logger.Debug().
			Any("staticsToBoot", staticsToBoot).
			Msg("Supervisor - Booting static zone(s)")
	}

	// boot dynamics if needed
	for l.currentProcessCounts[zoneProcessName]-l.currentZoneStatics < (zoneAssignedDynamics + l.minZoneProcesses) {
		err := l.startServerProcess(zoneProcessName)
		if err != nil {
			return err
		}

		bootedTotalDynamics := l.currentProcessCounts[zoneProcessName] - l.currentZoneStatics
		targetDynamics := zoneAssignedDynamics + l.minZoneProcesses

		l.logger.Info().
			Any("bootedTotalDynamics", bootedTotalDynamics).
			Any("zoneAssignedDynamics", zoneAssignedDynamics).
			Any("minZoneProcesses", l.minZoneProcesses).
			Any("targetDynamics", targetDynamics).
			Msg("Starting Dynamic Zone")
	}

	l.logger.DebugVv().
		Any("took", time.Now().Sub(now).String()).
		Any("currentProcessCounts", l.currentProcessCounts).
		Any("currentZoneDynamics", l.currentZoneDynamics).
		Any("currentZoneStatics", l.currentZoneStatics).
		Any("zoneIdleDynamics", zoneIdleDynamics).
		Any("zoneAssignedDynamics", zoneAssignedDynamics).
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
	cfg := l.serverconfig.Get()
	l.isRunning = cfg.Spire.LauncherStart
	l.runSharedMemory = cfg.WebAdmin.Launcher.RunSharedMemory
	l.runLoginserver = cfg.WebAdmin.Launcher.RunLoginserver
	l.runQueryServ = cfg.WebAdmin.Launcher.RunQueryServ
	l.minZoneProcesses = cfg.WebAdmin.Launcher.MinZoneProcesses
	l.serverLongName = cfg.Server.World.Longname
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

	if env.IsAppModeWebserver() && cfg.WebAdmin != nil {
		if cfg.WebAdmin.Discord != nil {
			l.watchCrashLogs = len(cfg.WebAdmin.Discord.CrashLogWebhook) > 0
			if l.watchCrashLogs {
				l.discordWebhook = cfg.WebAdmin.Discord.CrashLogWebhook
			}
		}
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

// StopCancel stops the server cancel
func (l *Launcher) StopCancel() error {
	stopTimerMutex.Lock()
	l.stopTimer = 0
	stopTimerMutex.Unlock()

	l.logger.Info().Msg("Timed stop cancelled")

	_ = l.serverApi.MessageWorld("[SERVER MESSAGE] Server shutdown cancelled")

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
			}

			if stat.ModTime().After(l.configLastModified) {
				l.configLastModified = stat.ModTime()
				l.loadServerConfig()
				l.logger.Debug().Msg("Watchdog - Detected server config change")
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

		if strings.Contains(cmdline, "eqemu-server:launcher start") && p.Pid != int32(os.Getpid()) && !strings.Contains(cmdline, "go run") {
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

// startCrashLogWatcher starts the crash log watcher
func (l *Launcher) startCrashLogWatcher() {
	if !l.watchCrashLogs {
		return
	}

	if l.watcher != nil {
		return
	}

	var err error

	// Create new watcher.
	l.watcher, err = rfsnotify.NewWatcher()
	if err != nil {
		l.logger.Debug().
			Any("error", err.Error()).
			Msg("Error creating watcher")
		return
	}

	defer l.watcher.Close()

	done := make(chan bool)

	// Start listening for events.
	go func(l *Launcher) {
		for {
			if l.watcher == nil {
				return
			}

			select {
			case event, ok := <-l.watcher.Events:
				if !ok {
					l.logger.Info().Msg("Crash Log Watcher error - Closing watcher")
					if l.watcher != nil {
						l.watcher.Close()
					}
					l.watcher = nil
					return
				}

				// if event is create or write
				if event.Op == fsnotify.Create {
					l.logger.Info().Any("event.Name", event.Name).Msg("Crash Log Watcher - Detected crash log change")

					// ship to discord
					filename := filepath.Base(event.Name)
					contents, err := os.ReadFile(event.Name)
					if err != nil {
						l.logger.Debug().
							Any("error", err.Error()).
							Msg("Error reading crash log file")
					}

					discord.SendMessage(
						l.discordWebhook,
						fmt.Sprintf(
							"**Crash Report** | **Server** [%s] **File** [%s] ",
							l.serverLongName,
							filename,
						),
						string(contents),
					)

				}

			case err, ok := <-l.watcher.Errors:
				if !ok {
					l.logger.Info().Msg("Crash Log Watcher error - Closing watcher")
					if l.watcher != nil {
						l.watcher.Close()
					}
					l.watcher = nil
					return
				}
				log.Println("error:", err)
			}
		}
	}(l)

	// Add a path.
	path := filepath.Join(l.pathmgmt.GetLogsDirPath(), "crashes")

	l.logger.Info().Any("path", path).Msg("Crash Log Watcher - Watching for changes")

	// check if the path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		l.logger.Debug().Any("path", path).Msg("Crash path does not exist")
		return // path does not exist
	}

	err = l.watcher.AddRecursive(path)
	if err != nil {
		l.logger.Debug().
			Any("error", err.Error()).
			Msg("Error adding path to watcher")
	}

	<-done
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
