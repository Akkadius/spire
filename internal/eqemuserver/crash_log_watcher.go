package eqemuserver

import (
	"fmt"
	"github.com/Akkadius/spire/internal/discord"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/rfsnotify"
	"github.com/Akkadius/spire/internal/websocket"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
	"time"
)

const crashLogWatcherLoopTimer = 10 * time.Second

type CrashLogWatcher struct {
	logger       *logger.AppLogger
	serverconfig *eqemuserverconfig.Config
	pathmgmt     *pathmgmt.PathManagement
	websocketMgr *websocket.ClientManager

	// properties
	configLastModified time.Time
	watcher            *rfsnotify.RWatcher

	// meta properties
	watchCrashLogs bool
	discordWebhook string
	serverLongName string
}

func NewCrashLogWatcher(
	logger *logger.AppLogger,
	serverconfig *eqemuserverconfig.Config,
	pathmgmt *pathmgmt.PathManagement,
	websocketMgr *websocket.ClientManager,
) *CrashLogWatcher {
	return &CrashLogWatcher{
		logger:       logger,
		serverconfig: serverconfig,
		pathmgmt:     pathmgmt,
		websocketMgr: websocketMgr,
	}
}

func (l *CrashLogWatcher) Process() {
	for {
		now := time.Now()
		stat, err := os.Stat(l.pathmgmt.GetEQEmuServerConfigFilePath())
		if err != nil {
			l.logger.Debug().
				Any("error", err.Error()).
				Any("watchCrashLogs", l.watchCrashLogs).
				Msg("Error getting server config file stat")
		}

		l.logger.DebugVvv().
			Any("configStatTime", time.Now().Sub(now).String()).
			Any("configLastModified", l.configLastModified).
			Any("watchCrashLogs", l.watchCrashLogs).
			Msg("Main crash log watcher process loop")

		if stat.ModTime().After(l.configLastModified) {
			l.configLastModified = stat.ModTime()
			l.loadServerConfig()
			l.logger.Debug().
				Any("configLastModified", l.configLastModified).
				Any("statTime", time.Now().Sub(now).String()).
				Any("watchCrashLogs", l.watchCrashLogs).
				Msg("Detected server config change")
		}

		if l.watchCrashLogs && l.watcher == nil {
			go l.startCrashLogWatcher()
		}

		l.logger.DebugVvv().Any("took", time.Now().Sub(now).String()).Msg("Crash Log Watcher loop")

		time.Sleep(crashLogWatcherLoopTimer)
	}
}

// loadServerConfig loads the server config
// this is called on startup and when the server config changes
func (l *CrashLogWatcher) loadServerConfig() {
	cfg, _ := l.serverconfig.Get()
	l.serverLongName = cfg.Server.World.Longname

	// kill the watcher if it exists
	if l.watcher != nil {
		l.watcher.Close()
		l.watcher = nil
	}

	if cfg.WebAdmin != nil {
		if cfg.WebAdmin.Discord != nil {
			l.watchCrashLogs = len(cfg.WebAdmin.Discord.CrashLogWebhook) > 0
			if l.watchCrashLogs {
				l.discordWebhook = cfg.WebAdmin.Discord.CrashLogWebhook
			}
		}
	}
}

// startCrashLogWatcher starts the crash log watcher
func (l *CrashLogWatcher) startCrashLogWatcher() {
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
	go func(l *CrashLogWatcher) {
		for {
			if l.watcher == nil {
				return
			}

			select {
			case event, ok := <-l.watcher.Events:
				if !ok {
					l.logger.Info().Any("event", event).Msg("Watcher error - Closing watcher")
					if l.watcher != nil {
						l.watcher.Close()
					}
					l.watcher = nil
					return
				}

				// if event is create or write
				if event.Op == fsnotify.Create {
					l.logger.Info().
						Any("event.Name", event.Name).
						Msg("Detected crash log change")

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
					l.logger.Info().Msg("Closing watcher")
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

	l.logger.Info().Any("path", path).Msg("Watching for changes")

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
