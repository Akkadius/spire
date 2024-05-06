package eqemuserver

import (
	"fmt"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/shirou/gopsutil/v3/process"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

type Launcher struct {
	logger       *logger.DebugLogger
	serverconfig *eqemuserverconfig.Config
	settings     *spire.Settings
	pathmgmt     *pathmgmt.PathManagement
	serverApi    *Client

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
}

func NewLauncher(
	logger *logger.DebugLogger,
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

	l.serverProcessLauncherWatchdog()

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
			Any("configStatTime", now.Sub(time.Now()).String()).
			Any("configLastModified", l.configLastModified).
			Any("isRunning", l.isRunning).
			Msg("Main launcher process loop")

		if stat.ModTime().After(l.configLastModified) {
			l.configLastModified = stat.ModTime()
			l.loadServerConfig()
			l.logger.Debug().
				Any("configLastModified", l.configLastModified).
				Any("statTime", now.Sub(time.Now()).String()).
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

	fmt.Println("Spire > Starting server launcher")

	// start shared memory if needed
	// this needs to be started and completed before the server processes
	if l.runSharedMemory {
		fmt.Println("Spire > Starting shared memory")
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

	fmt.Println("Spire > Restarting server launcher")
	return l.StartLauncherProcess()
}

// Stop stops the launcher
func (l *Launcher) Stop() error {
	fmt.Println("Spire > Stopping server launcher")

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

		exe, err := p.Exe()
		if err != nil {
			l.logger.Debug().
				Any("error", err.Error()).
				Any("pid", p.Pid).
				Msg("Error getting process exe")
		}

		baseProcessName := filepath.Base(exe)
		for _, s := range l.GetServerProcessNames() {
			if s == baseProcessName {
				l.logger.Debug().
					Any("pid", p.Pid).
					Any("baseProcessName", baseProcessName).
					Any("cmdline", cmdline).
					Msg("Stop - Killing server process")

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
			Any("exe", exe).
			Any("cmdline", cmdline).
			Any("baseProcessName", baseProcessName).
			Msg("Stop - Checking process")

		// kill any instances of launchers
		isServerLauncher := strings.Contains(cmdline, "server-launcher") || strings.Contains(cmdline, "eqemu-server:launcher")

		if isServerLauncher && p.Pid != int32(os.Getpid()) {
			if err := p.Terminate(); err != nil {
				l.logger.Debug().
					Any("error", err.Error()).
					Any("pid", p.Pid).
					Msg("Error killing process")
			}
		}
	}

	fmt.Println("Spire > Stopped server launcher")

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

		cwd, err := p.Cwd()
		if err != nil {
			if strings.Contains(err.Error(), "no such file or directory") {
				continue
			}
			if strings.Contains(err.Error(), "permission denied") {
				continue
			}
			l.logger.Debug().
				Any("error", err.Error()).
				Any("pid", p.Pid).
				Msg("Error getting process cwd")
		}

		exe, err := p.Exe()
		if err != nil {
			l.logger.Debug().
				Any("error", err.Error()).
				Any("pid", p.Pid).
				Msg("Error getting process exe")
		}

		baseProcessName := filepath.Base(exe)
		for _, s := range l.GetServerProcessNames() {
			if s == baseProcessName {
				l.logger.DebugVv().
					Any("pid", p.Pid).
					Any("baseProcessName", baseProcessName).
					Any("cmdline", cmdline).
					Msg("Supervisor - Found server process")

				if _, ok := l.currentProcessCounts[s]; !ok {
					l.currentProcessCounts[s] = 0
				}

				if s == zoneProcessName {
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
			Any("cwd", cwd).
			Any("exe", exe).
			Any("cmdline", cmdline).
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
		fmt.Println("Spire > Starting World")
		err := l.startServerProcess(worldProcessName)
		if err != nil {
			return err
		}
	}

	// boot loginserver if needed
	if l.runLoginserver && l.currentProcessCounts[loginServerProcessName] == 0 {
		fmt.Println("Spire > Starting LoginServer")
		err := l.startServerProcess(loginServerProcessName)
		if err != nil {
			return err
		}
	}

	// boot queryserv if needed
	if l.runQueryServ && l.currentProcessCounts[queryServProcessName] == 0 {
		fmt.Println("Spire > Starting QueryServ")
		err := l.startServerProcess(queryServProcessName)
		if err != nil {
			return err
		}
	}

	// boot ucs if needed
	if l.currentProcessCounts[ucsProcessName] == 0 {
		fmt.Println("Spire > Starting UCS")
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
			fmt.Printf("Spire > Started Static Zones (%v) [%+v]\n", len(staticsToBoot), staticsToBoot)
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

		fmt.Printf("Spire > Started Dynamic Zones (%v)\n", zoneDynamicsToBoot)
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
			staticZonesToBoot = append(staticZonesToBoot, z)
		}
	}

	l.staticZonesToBoot = staticZonesToBoot
}

func (l *Launcher) SetStopTimer(timer int) {
	l.stopTimer = timer
}

// StopCancel stops the server cancel
func (l *Launcher) StopCancel() error {
	// handle

	return nil
}

func (l *Launcher) serverProcessLauncherWatchdog() {
	go func() {
		for {
			l.loadServerConfig()
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

		if strings.Contains(cmdline, "eqemu-server:launcher start") && p.Pid != int32(os.Getpid()) {
			l.logger.Debug().
				Any("pid", p.Pid).
				Any("cmdline", cmdline).
				Msg("Launcher process already running")

			fmt.Println("Launcher process already running")

			return fmt.Errorf("launcher process already running")
		}
	}
	return nil
}
