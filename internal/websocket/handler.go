package websocket

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type Handler struct {
	pathmgmt *pathmgmt.PathManagement
}

func NewHandler(
	pathmgmt *pathmgmt.PathManagement,
) *Handler {
	return &Handler{
		pathmgmt: pathmgmt,
	}
}

func (h *Handler) HandleHello(ws *websocket.Conn, msg string) error {
	err := websocket.Message.Send(ws, "Hello, Client!")
	if err != nil {
		return err
	}
	return nil
}

type SpireExecCommand struct {
	Action    string   `json:"action"`
	Command   string   `json:"command"`
	Arguments []string `json:"args"`
}

func (h *Handler) HandleExecServerBin(ws *websocket.Conn, msg string) error {
	var m SpireExecCommand
	err := json.Unmarshal([]byte(msg), &m)
	if err != nil {
		return err
	}

	// execute and get a pipe
	basePath := filepath.Join(h.pathmgmt.GetEQEmuServerPath(), "bin")
	startCmd := ""
	if _, err := os.Stat(filepath.Join(basePath, m.Command)); err == nil {
		startCmd = filepath.Join(basePath, m.Command)
	} else if _, err := os.Stat(filepath.Join(startCmd, fmt.Sprintf("%v.exe", m.Command))); err == nil {
		startCmd = filepath.Join(basePath, fmt.Sprintf("%v.exe", m.Command))
	}

	if len(startCmd) == 0 {
		err = websocket.Message.Send(ws, fmt.Sprintf("Error: Unable to find exec binary [%v]", startCmd))
		if err != nil {
			return err
		}
	}

	go func() {
		cmd := exec.Command(startCmd, m.Arguments...)
		cmd.Dir = h.pathmgmt.GetEQEmuServerPath()
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

	return nil
}

func (h *Handler) HandleUnauthorized(ws *websocket.Conn) error {
	err := websocket.Message.Send(ws, "Unauthorized")
	if err != nil {
		return err
	}

	return errors.New("Unauthorized")
}
