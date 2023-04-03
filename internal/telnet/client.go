package telnet

import (
	"fmt"
	"github.com/Akkadius/spire/internal/env"
	"github.com/k0kubun/pp/v3"
	"github.com/sirupsen/logrus"
	"github.com/ziutek/telnet"
	"io"
	"net"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Client struct {
	debugging bool
	t         *telnet.Conn
	logger    *logrus.Logger
	mu        sync.Mutex
}

func NewClient(logger *logrus.Logger) *Client {
	return &Client{
		debugging: env.GetInt("DEBUG", "0") >= 3,
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
	c.t, err = telnet.DialTimeout("tcp", "localhost:9000", d)
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
		_ = sendln(c.t, "acceptmessages on")
		expect(c.t, ">")
	}

	return nil
}

func (c *Client) Command(cmd string) (string, error) {
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

	sendln(c.t, cmd)

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
			c.logger.Warnf("[telnet] read failed: %s", err)
			c.Close()
			return "", err
		}

		output += string(data)

		if strings.Contains(output, linebreak) {
			output = strings.Replace(output, linebreak, "", 1)
			c.debug("[Output] %v", output)
			return output, nil
		}
	}
}

func (c *Client) Close() {
	if c.t != nil {
		err := c.t.Close()
		if err != nil {
			c.logger.Error(err)
		}
		c.t = nil
	}
}

func (c *Client) debug(msg string, a ...interface{}) {
	if c.debugging {
		_, file, _, ok := runtime.Caller(1)
		if ok {
			file = filepath.Base(file)
			if len(a) > 0 {
				pp.Printf(fmt.Sprintf("[%v] ", file) + fmt.Sprintf(msg, a...) + "\n")
				return
			}
			pp.Printf(fmt.Sprintf("[%v] ", file) + fmt.Sprintf(msg) + "\n")
		}
	}
}
