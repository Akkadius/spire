package eqemuserver

import (
	"fmt"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/rfsnotify"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"time"
)

type QuestHotReloadWatcher struct {
	logger       *logger.DebugLogger
	serverconfig *eqemuserverconfig.Config
	pathmgmt     *pathmgmt.PathManagement
	serverApi    *Client

	// properties
	configLastModified time.Time
	isRunning          bool
	watcher            *rfsnotify.RWatcher
	fileSizes          map[string]int
}

func NewQuestHotReloadWatcher(
	logger *logger.DebugLogger,
	serverconfig *eqemuserverconfig.Config,
	pathmgmt *pathmgmt.PathManagement,
	serverApi *Client,
) *QuestHotReloadWatcher {
	l := &QuestHotReloadWatcher{
		logger:       logger,
		serverconfig: serverconfig,
		pathmgmt:     pathmgmt,
		serverApi:    serverApi,
	}

	if env.IsAppEnvLocal() {
		go func() {
			l.Process()
		}()
	}

	return l
}

// Process runs the process loop for the QuestHotReloadWatcher
// It only does anything if the server has the QuestHotReloadWatcher set to run
func (l *QuestHotReloadWatcher) Process() {
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
			Msg("Main QuestHotReloadWatcher process loop")

		if stat.ModTime().After(l.configLastModified) {
			l.configLastModified = stat.ModTime()
			l.loadServerConfig()
			l.logger.Debug().
				Any("configLastModified", l.configLastModified).
				Any("statTime", now.Sub(time.Now()).String()).
				Any("isRunning", l.isRunning).
				Msg("Detected server config change")
		}

		if l.isRunning && l.watcher == nil {
			go func(l *QuestHotReloadWatcher) {
				l.Start()
			}(l)
		}

		if !l.isRunning && l.watcher != nil {
			l.Stop()
		}

		time.Sleep(processLoopTimer)
	}
}

// loadServerConfig loads the server config
// this is called on startup and when the server config changes
func (l *QuestHotReloadWatcher) loadServerConfig() {
	cfg := l.serverconfig.Get()
	l.isRunning = cfg.WebAdmin.Quests.HotReload
}

func (l *QuestHotReloadWatcher) Stop() {
	fmt.Println("Spire > Stopping watcher")

	if l.watcher != nil {
		err := l.watcher.Close()
		if err != nil {
			l.logger.Debug().
				Any("error", err.Error()).
				Msg("Error closing watcher")
		} else {
			l.watcher = nil
		}
	}
}

func (l *QuestHotReloadWatcher) Start() {
	var err error

	// Create new watcher.
	l.watcher, err = rfsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer l.watcher.Close()

	var fileSizes map[string]int

	// Start listening for events.
	go func(l *QuestHotReloadWatcher) {
		for {
			if l.watcher == nil {
				return
			}

			select {
			case event, ok := <-l.watcher.Events:
				if !ok {
					break
				}

				// if event is create or write
				if event.Op == fsnotify.Create || event.Op == fsnotify.Write {
					if fileSizes == nil {
						fileSizes = make(map[string]int)
					}

					// initialize key if it doesn't exist
					if _, ok := fileSizes[event.Name]; !ok {
						fileSizes[event.Name] = 0
					}

					// get file size and store it
					fileInfo, err := os.Stat(event.Name)
					if err != nil {
						log.Println(err)
					}

					if fileInfo.Size() == 0 {
						l.logger.Debug().Any("file", event.Name).Msg("File size is 0, skipping")
						break
					}

					if _, ok := fileSizes[event.Name]; ok {
						l.logger.Debug().
							Any("current", fileSizes[event.Name]).
							Any("new", int(fileInfo.Size())).
							Any("file", event.Name).
							Msg("File size comparison")
						if fileSizes[event.Name] == int(fileInfo.Size()) {
							l.logger.Debug().Any("file", event.Name).Msg("File size is the same, skipping reload")
							break
						}
					}

					fileSizes[event.Name] = int(fileInfo.Size())

					// send reload request to server
					fmt.Printf("Spire > Hot Quest Reload > Reloading quests from [%s]\n", event.Name)
				}

			case err, ok := <-l.watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}(l)

	fmt.Println("Watching for changes in", l.pathmgmt.GetQuestsDir())

	// Add a path.
	err = l.watcher.AddRecursive(l.pathmgmt.GetQuestsDir())
	if err != nil {
		l.logger.Debug().
			Any("error", err.Error()).
			Msg("Error adding path to watcher")
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}
