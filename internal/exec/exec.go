package exec

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/Akkadius/spire/internal/console"
	"os"
	"os/exec"
	"runtime/debug"
	"strings"
	"time"
)

// Cfg is the configuration for the Exec function.
type Cfg struct {
	Command           string
	Args              []string
	ShadowPrintBefore bool      // shadow prints the command before execution
	ShadowPrint       bool      // shadow prints the command after execution
	Interactive       bool      // interactive mode
	Background        bool      // start in background
	NoEscapeArgsPrint bool      // don't escape args when printing
	Timeout           int       // timeout in seconds
	Dir               string    // directory to execute the command in
	start             time.Time // start time of execution
}

// Result is a command result
type Result struct {
	Stdout   string
	Stderr   string
	Pid      int
	Process  *os.Process
	ExitCode int
	Config   Cfg
}

// Command executes a command and returns the output.
func Command(cfg Cfg) (Result, error) {
	ctx := context.Background()
	if cfg.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
		defer cancel()
	}

	// execute the command
	cmd := exec.CommandContext(ctx, cfg.Command, cfg.Args...)
	var stdout, stderr bytes.Buffer

	// set the directory if provided
	if len(cfg.Dir) > 0 {
		cmd.Dir = cfg.Dir
	}

	exitCode := 0

	// set the environment
	cmd.Env = os.Environ()

	// tie the buffers
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// interactive mode
	if cfg.Interactive {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		// let the command handler catch stderr
		cmd.Stderr = os.Stderr
	}

	if cfg.ShadowPrintBefore {
		shadowPrintCmd(cfg)
	}

	// measure execution time
	cfg.start = time.Now()

	// run the command
	err := cmd.Start()
	var eerr *exec.ExitError
	if errors.As(err, &eerr) {
		exitCode = eerr.ExitCode()
	}

	r := Result{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		Pid:      cmd.Process.Pid,
		Process:  cmd.Process,
		ExitCode: exitCode,
	}

	if err != nil {
		return resultCallback(cfg, r), errors.New(stdout.String() + stderr.String())
	}

	if !cfg.Background {
		if err := cmd.Wait(); err != nil {
			if errors.As(err, &eerr) {
				r.ExitCode = eerr.ExitCode()
			}

			r.Stdout = stdout.String()
			r.Stderr = stderr.String()

			return resultCallback(cfg, r), errors.New(stdout.String() + stderr.String())
		}
	}

	if cfg.Background {
		err := cmd.Process.Release()
		if err != nil {
			return resultCallback(cfg, r), err
		}
	}

	r.Stdout = stdout.String()
	r.Stderr = stderr.String()

	return resultCallback(cfg, r), nil
}

// resultCallback is a callback for the execution result
func resultCallback(cfg Cfg, result Result) Result {
	result.Config = cfg

	// print the command
	if cfg.ShadowPrint {
		shadowPrintCmd(cfg)
	}

	return result
}

// shadowPrintCmd prints the command to the console.
func shadowPrintCmd(cfg Cfg) {
	// copy the args
	var args []string
	for i := range cfg.Args {
		args = append(args, cfg.Args[i])
	}

	// escape args
	if !cfg.NoEscapeArgsPrint {
		for i := range args {
			args[i] = Quote(args[i])
		}
	}

	isAsync := ""
	if strings.Contains(string(debug.Stack()), "in goroutine") {
		isAsync = "(async)"
	}

	timing := ""
	if !cfg.start.IsZero() {
		timing = fmt.Sprintf(" (%s)", time.Since(cfg.start).Round(time.Millisecond).String())
	}

	_, _ = fmt.Fprintf(
		os.Stderr,
		"%s > %s%s%s %s%s%s %s %s\n",
		console.BoldHighIntensityBlack,
		console.Reset,
		console.HighIntensityBlack,
		cfg.Command,
		strings.Join(args, " "),
		console.FadedGray,
		timing,
		isAsync,
		console.Reset,
	)
}
