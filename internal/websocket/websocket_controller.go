package websocket

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type Controller struct {
	logger   *logrus.Logger
	pathmgmt *pathmgmt.PathManagement
}

func NewController(
	logger *logrus.Logger,
	pathmgmt *pathmgmt.PathManagement,
) *Controller {
	return &Controller{
		logger:   logger,
		pathmgmt: pathmgmt,
	}
}

func (a *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "websocket", a.websocketHandler, nil),
	}
}

type SpireWebsocketMessage struct {
	Action    string   `json:"action"`
	Command   string   `json:"command"`
	Arguments []string `json:"args"`
}

func (a *Controller) websocketHandler(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {

		defer ws.Close()
		for {

			// Read
			msg := ""
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}

			// just close empty messages
			if len(msg) == 0 {
				// close if error
				if err := ws.Close(); err != nil {
					break
				}
			}

			// Write
			if len(msg) > 0 {
				var m SpireWebsocketMessage
				err := json.Unmarshal([]byte(msg), &m)
				if err != nil {
					log.Println(err)
				}

				if m.Action == "hello" {
					err = websocket.Message.Send(ws, "Hello, Client!")
					if err != nil {
						// close if error
						if err := ws.Close(); err != nil {
							break
						}
						c.Logger().Error(err)
					}
				}

				if m.Action == "exec_server_bin" {
					// execute and get a pipe

					basePath := filepath.Join(a.pathmgmt.GetEQEmuServerPath(), "bin")
					startCmd := ""
					if _, err := os.Stat(filepath.Join(basePath, m.Command)); err == nil {
						startCmd = filepath.Join(basePath, m.Command)
					} else if _, err := os.Stat(filepath.Join(startCmd, fmt.Sprintf("%v.exe", m.Command))); err == nil {
						startCmd = filepath.Join(basePath, fmt.Sprintf("%v.exe", m.Command))
					}

					if len(startCmd) == 0 {
						err = websocket.Message.Send(ws, fmt.Sprintf("Error: Unable to find exec binary [%v]", startCmd))
						if err != nil {
							// close if error
							if err := ws.Close(); err != nil {
								break
							}
							c.Logger().Error(err)
						}
					}

					go func() {
						cmd := exec.Command(startCmd, m.Arguments...)
						cmd.Dir = a.pathmgmt.GetEQEmuServerPath()
						cmd.Env = os.Environ()
						if runtime.GOOS == "linux" {
							cmd.Env = append(cmd.Env, "IS_TTY=true")
						}
						//pp.Println(cmd.Env)
						stdout, err := cmd.StdoutPipe()
						if err != nil {
							log.Println(err)
							return
						}
						stderr, err := cmd.StderrPipe()
						if err != nil {
							log.Println(err)
							return
						}

						if err := cmd.Start(); err != nil {
							log.Println(err)
							return
						}

						s := bufio.NewScanner(io.MultiReader(stdout, stderr))
						for s.Scan() {
							err = websocket.Message.Send(ws, s.Text())
							if err != nil {
								// close if error
								if err := ws.Close(); err != nil {
									break
								}
							}
						}

						if err := cmd.Wait(); err != nil {
							log.Println(err)
							return
						}
					}()
				}
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
