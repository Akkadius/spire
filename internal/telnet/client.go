package telnet

import (
	"fmt"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/ziutek/telnet"
	"strings"
	"sync"
	"time"
)

// Client is a telnet client
type Client struct {
	debugging bool
	t         *telnet.Conn
	logger    *logger.AppLogger
	mu        sync.Mutex
}

// NewClient creates a new telnet client
func NewClient(logger *logger.AppLogger) *Client {
	return &Client{
		debugging: env.GetInt("TELNET_DEBUG", "0") >= 3,
		logger:    logger,
	}
}

const (
	linebreak = "\n\r> "
)

// Connect connects to the telnet server
func (c *Client) Connect() error {
	var err error

	// If the connection is already alive, return early
	if c.t != nil {
		if c.isConnectionAlive() {
			return nil
		}
		c.Close()
	}

	d := 2 * time.Second // Increased timeout for stability
	c.t, err = telnet.DialTimeout("tcp", "127.0.0.1:9000", d)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			c.fail(fmt.Errorf("panic during Connect: %v", r))
		}
	}()

	if err = c.t.SetReadDeadline(time.Now().Add(d)); err != nil {
		return c.fail(err)
	}
	if err = c.t.SetWriteDeadline(time.Now().Add(d)); err != nil {
		return c.fail(err)
	}
	if err = c.t.SetEcho(false); err != nil {
		return c.fail(err)
	}

	if err = c.expect("assuming admin"); err != nil {
		return c.fail(err)
	}

	c.debug("\n###################################\n# Logging into World\n###################################")

	if err = c.expect(">"); err != nil {
		return c.fail(err)
	}
	if err = c.sendln("echo off"); err != nil {
		return c.fail(err)
	}
	if err = c.expect(">"); err != nil {
		return c.fail(err)
	}
	if err = c.sendln("acceptmessages off"); err != nil {
		return c.fail(err)
	}
	if err = c.expect(">"); err != nil {
		return c.fail(err)
	}

	return nil
}

// fail closes the connection and returns the error
func (c *Client) fail(err error) error {
	c.Close()
	return err
}

// expect reads the telnet connection until it finds the expected string
func (c *Client) expect(d ...string) error {
	err := c.t.SkipUntil(d...)
	if err != nil {
		return c.fail(err)
	}
	return nil
}

// sendln sends a string to the telnet connection
func (c *Client) sendln(s string) error {
	defer func() {
		if r := recover(); r != nil {
			c.fail(fmt.Errorf("panic in sendln: %v", r))
		}
	}()

	buf := make([]byte, len(s)+1)
	copy(buf, s)
	buf[len(s)] = '\n'
	_, err := c.t.Write(buf)

	if err != nil {
		return c.fail(fmt.Errorf("failed to send command: %w", err))
	}

	return err
}

// CommandConfig is a configuration for a command
type CommandConfig struct {
	Command     string // the command to send
	EnforceJson bool   // error the connection if the response is not a json
}

// Command sends a command to the telnet server and returns the output
func (c *Client) Command(cmd CommandConfig) (string, error) {
	var err error

	c.mu.Lock()
	defer c.mu.Unlock()

	err = c.Connect()
	if err != nil {
		return "", c.fail(err)
	}

	err = c.t.SetReadDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		return "", c.fail(err)
	}
	err = c.t.SetWriteDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		return "", c.fail(err)
	}

	if c.debugging {
		c.logger.Debug().Any("command", cmd.Command).Msg("Sending command")
	}

	err = c.sendln(cmd.Command)
	if err != nil {
		return "", c.fail(err)
	}

	defer func() {
		if r := recover(); r != nil {
			c.debug("Panic in read, close connection")
			c.Close()
		}
	}()

	var data []byte
	var output string
	for {
		start := time.Now()
		data, err = c.t.ReadUntil(linebreak)
		c.debug("Read operation took %v", time.Since(start))
		if err != nil {
			c.logger.Debug().Err(err).Msg("Warning - Failed to read from telnet, this may mean World API is down and not accepting connections.")
			return "", c.fail(err)
		}

		output += string(data)

		if strings.Contains(output, linebreak) {
			output = strings.Replace(output, linebreak, "", 1)
			if c.debugging {
				c.logger.DebugVvv().
					Any("response", output).
					Any("command", cmd.Command).
					Any("enforce_json", cmd.EnforceJson).
					Any("took", time.Since(start).String()).
					Msg("Telnet response")
			}

			// if we are enforcing json, make sure the output is json
			if cmd.EnforceJson {
				if !strings.HasPrefix(output, "{") && !strings.HasSuffix(output, "}") {
					return "", c.fail(fmt.Errorf("response was not json: %v", output))
				}
			}

			return output, nil
		}
	}
}

// Close closes the telnet connection
func (c *Client) Close() {
	if c.t != nil {
		c.debug("Closing telnet connection")
		err := c.t.Close()
		if err != nil {
			c.logger.Error().Err(err).Msg("Failed to close telnet connection")
		}
		c.t = nil
	}
}

// debug logs a debug message
func (c *Client) debug(msg string, a ...interface{}) {
	if c.debugging {
		c.logger.Debug().Msgf(msg, a...)
	}
}

// isConnectionAlive checks if the connection is alive
func (c *Client) isConnectionAlive() bool {
	if c.t == nil {
		return false
	}

	// Perform a zero-byte write to check if the connection is alive
	if err := c.t.SetWriteDeadline(time.Now().Add(1 * time.Millisecond)); err != nil {
		return false
	}

	_, err := c.t.Write([]byte{})
	if err != nil {
		return false // Connection is dead
	}

	return true
}
