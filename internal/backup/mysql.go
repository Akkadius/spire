package backup

import (
	"bufio"
	"fmt"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"os/exec"
	"path/filepath"
	"strings"
)

type Mysql struct {
	pathmanager *pathmgmt.PathManagement
}

func NewMysql(pathmanager *pathmgmt.PathManagement) *Mysql {
	return &Mysql{pathmanager: pathmanager}
}

type BackupRequest struct {
	DumpAllTables   bool `json:"dump_all_tables"`
	ContentTables   bool `json:"content_tables"`
	PlayerTables    bool `json:"player_tables"`
	BotTables       bool `json:"bot_tables"`
	StateTables     bool `json:"state_tables"`
	SystemTables    bool `json:"system_tables"`
	QueryServTables bool `json:"query_serv_tables"`
	LoginTables     bool `json:"login_tables"`
	Compress        bool `json:"compress"`
}

type MysqlBackupResponse struct {
	Command  string `json:"command"`
	StdOut   string `json:"stdout"`
	FilePath string `json:"file_path"`
}

func (m *Mysql) Backup(r BackupRequest) MysqlBackupResponse {
	var args = []string{"database:dump"}

	args = append(args, "--table-lock=false")

	if r.DumpAllTables {
		args = append(args, "--all")
	} else {
		if r.ContentTables {
			args = append(args, "--content-tables")
		}
		if r.PlayerTables {
			args = append(args, "--player-tables")
		}
		if r.LoginTables {
			args = append(args, "--login-tables")
		}
		if r.SystemTables {
			args = append(args, "--system-tables")
		}
		if r.StateTables {
			args = append(args, "--state-tables")
		}
		if r.BotTables {
			args = append(args, "--bot-tables")
		}
	}

	if r.Compress {
		args = append(args, "--compress")
	}

	runPath := filepath.Join(m.pathmanager.GetEQEmuServerPath(), "bin", "world")
	cmd := exec.Command(runPath, args...)
	stdout, _ := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	err := cmd.Start()
	if err != nil {
		fmt.Printf("failed to start command: %v", err)
	}

	output := ""
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Printf("[Backup] %v\n", scanner.Text())
		output += fmt.Sprintf("%v\n", scanner.Text())
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Printf("failed to run command: %v", err)
	}

	finalPath := ""
	for _, s := range strings.Split(output, "\n") {
		if strings.Contains(s, "dump created at ") {
			split := strings.Split(s, "dump created at")
			if len(split) > 0 {
				path := strings.ReplaceAll(split[1], "]", "")
				path = strings.ReplaceAll(path, "[", "")
				path = strings.TrimSpace(path)
				path = filepath.Join(m.pathmanager.GetEQEmuServerPath(), path)
				finalPath = path
			}
		}
	}

	return MysqlBackupResponse{
		Command:  fmt.Sprintf("%v %v", runPath, strings.Join(args, " ")),
		StdOut:   output,
		FilePath: finalPath,
	}
}
