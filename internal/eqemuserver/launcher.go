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
	"io"
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
	configLastModified   time.Time
	isRunning            bool // this is set by the server config
	runSharedMemory      bool
	runLoginserver       bool
	runQueryServ         bool
	minZoneProcesses     int
	currentZoneDynamics  int
	currentZoneStatics   int
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

	if env.IsAppModeWebserver() {
		l.serverProcessLauncherWatchdog()
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

	// start shared memory if needed
	// this needs to be started and completed before the server processes
	if l.runSharedMemory {
		l.logger.Info().Msg("Starting shared memory")
		err := l.startServerProcessSync(sharedMemoryProcessName)
		if err != nil {
			return err
		}
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

	list, err := l.serverApi.GetZoneList()
	if err != nil {
		l.logger.Debug().
			Any("error", err.Error()).
			Msg("Error getting zone list")
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
	zoneDynamicsToBoot := l.minZoneProcesses - zoneIdleDynamics
	if zoneDynamicsToBoot > 0 {
		l.logger.Debug().
			Any("zoneDynamicsToBoot", zoneDynamicsToBoot).
			Msg("Supervisor - Booting dynamic zone(s)")

		for i := 0; i < zoneDynamicsToBoot; i++ {
			err := l.startServerProcess(zoneProcessName)
			if err != nil {
				return err
			}
		}

		l.logger.Info().Msgf("Starting Dynamic Zones (%v)", zoneDynamicsToBoot)
	}

	l.logger.DebugVv().
		Any("took", time.Now().Sub(now).String()).
		Any("currentProcessCounts", l.currentProcessCounts).
		Any("currentZoneDynamics", l.currentZoneDynamics).
		Any("currentZoneStatics", l.currentZoneStatics).
		Any("zoneDynamicsToBoot", zoneDynamicsToBoot).
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

	if l.minZoneProcesses < 1 {
		l.minZoneProcesses = 5
	}

	if env.IsAppModeWebserver() && cfg.WebAdmin != nil {
		if cfg.WebAdmin.Discord != nil {
			l.watchCrashLogs = len(cfg.WebAdmin.Discord.CrashLogWebhook) > 0
			if l.watchCrashLogs {
				l.discordWebhook = cfg.WebAdmin.Discord.CrashLogWebhook
				l.StartCrashLogWatcher()
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

func (l *Launcher) serverProcessLauncherWatchdog() {
	go func() {
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

func (l *Launcher) StartCrashLogWatcher() {
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

	return ProcessDetails{
		Pid:             p.Pid,
		Cmdline:         cmdline,
		Cwd:             cwd,
		Exe:             exe,
		BaseProcessName: baseProcessName,
	}
}

// Copy copies the contents of the file at srcpath to a regular file
// at dstpath. If the file named by dstpath already exists, it is
// truncated. The function does not copy the file mode, file
// permission bits, or file attributes.
func copyFile(srcpath, dstpath string) (err error) {
	r, err := os.Open(srcpath)
	if err != nil {
		return err
	}
	defer r.Close() // ignore error: file was opened read-only.

	w, err := os.Create(dstpath)
	if err != nil {
		return err
	}

	defer func() {
		// Report the error, if any, from Close, but do so
		// only if there isn't already an outgoing error.
		if c := w.Close(); err == nil {
			err = c
		}
	}()

	_, err = io.Copy(w, r)
	return err
}
