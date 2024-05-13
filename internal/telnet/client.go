package telnet

import (
	"fmt"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/ziutek/telnet"
	"io"
	"net"
	"strings"
	"sync"
	"time"
)

type Client struct {
	debugging bool
	t         *telnet.Conn
	logger    *logger.AppLogger
	mu        sync.Mutex
}

func NewClient(logger *logger.AppLogger) *Client {
	return &Client{
		debugging: env.GetInt("TELNET_DEBUG", "0") >= 3,
		logger:    logger,
	}
}

const (
	linebreak = "\n\r> "
)

func expect(t *telnet.Conn, d ...string) bool {
	err := t.SkipUntil(d...)
	if err != nil {
		return false
	}

	return true
}

func sendln(t *telnet.Conn, s string) error {
	defer func() {
		if r := recover(); r != nil {
			t.Close()
		}
	}()

	buf := make([]byte, len(s)+1)
	copy(buf, s)
	buf[len(s)] = '\n'
	_, err := t.Write(buf)
	return err
}

func (c *Client) Connect() error {
	var err error

	// connection check
	if c.t != nil {
		one := make([]byte, 1)
		_ = c.t.SetReadDeadline(time.Now())
		if _, err := c.t.Read(one); err == io.EOF {
			c.Close()
		}
		if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
			c.Close()
		}
	}

	if c.t != nil {
		return nil
	}

	d := 300 * time.Millisecond
	c.t, err = telnet.DialTimeout("tcp", "127.0.0.1:9000", d)
	if err != nil {
		return err
	}

	err = c.t.SetReadDeadline(time.Now().Add(d))
	if err != nil {
		return err
	}
	err = c.t.SetWriteDeadline(time.Now().Add(d))
	if err != nil {
		return err
	}

	err = c.t.SetEcho(false)
	if err != nil {
		return err
	}

	// what the console expects when connecting locally
	if expect(c.t, "assuming admin") {
		c.debug("\n###################################\n# Logging into World\n###################################")

		expect(c.t, ">")
		_ = sendln(c.t, "echo off")
		expect(c.t, ">")
		_ = sendln(c.t, "acceptmessages off")
		expect(c.t, ">")
	}

	return nil
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
		c.Close()
		return "", err
	}

	err = c.t.SetReadDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		return "", err
	}
	err = c.t.SetWriteDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		return "", err
	}

	if c.debugging {
		c.logger.Debug().Any("command", cmd.Command).Msg("Sending command")
	}

	sendln(c.t, cmd.Command)

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
			c.logger.Warn().Err(err).Msg("Failed to read from telnet")
			c.Close()
			return "", err
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
					c.Close()
					return "", fmt.Errorf("response was not json: %v", output)
				}
			}

			return output, nil
		}
	}
}

func (c *Client) Close() {
	if c.t != nil {
		err := c.t.Close()
		if err != nil {
			c.logger.Error().Err(err).Msg("Failed to close telnet connection")
		}
		c.t = nil
	}
}

func (c *Client) debug(msg string, a ...interface{}) {
	if c.debugging {
		c.logger.Debug().Msgf(msg, a...)
	}
}
