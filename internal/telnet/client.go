package telnet

import (
	"github.com/sirupsen/logrus"
	"github.com/ziutek/telnet"
	"io"
	"net"
	"strings"
	"syscall"
	"time"
)

type Client struct {
	t      *telnet.Conn
	logger *logrus.Logger
}

func NewClient(logger *logrus.Logger) *Client {
	return &Client{
		logger: logger,
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
	buf := make([]byte, len(s)+1)
	copy(buf, s)
	buf[len(s)] = '\n'
	_, err := t.Write(buf)
	return err
}

func connCheck(conn net.Conn) error {
	var sysErr error = nil
	rc, err := conn.(syscall.Conn).SyscallConn()
	if err != nil {
		return err
	}
	err = rc.Read(func(fd uintptr) bool {
		var buf []byte = []byte{0}
		n, _, err := syscall.Recvfrom(int(fd), buf, syscall.MSG_PEEK|syscall.MSG_DONTWAIT)
		switch {
		case n == 0 && err == nil:
			sysErr = io.EOF
		case err == syscall.EAGAIN || err == syscall.EWOULDBLOCK:
			sysErr = nil
		default:
			sysErr = err
		}
		return true
	})
	if err != nil {
		return err
	}

	return sysErr
}

func (c *Client) Connect() error {
	var err error

	if c.t != nil {
		err := connCheck(c.t.Conn)
		if err != nil {
			c.Close()
			c.t = nil
		}
	}

	if c.t != nil {
		return nil
	}

	d := 1000 * time.Second
	c.t, err = telnet.DialTimeout("tcp", "localhost:9000", d)
	if err != nil {
		return err
	}

	err = c.t.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return err
	}
	err = c.t.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return err
	}

	err = c.t.SetEcho(false)
	if err != nil {
		return err
	}

	// what the console expects when connecting locally
	if expect(c.t, "assuming admin") {
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

	err = c.Connect()
	if err != nil {
		return "", err
	}

	err = c.t.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return "", err
	}
	err = c.t.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return "", err
	}

	sendln(c.t, cmd)

	buf := make([]byte, 512)
	output := ""
	i := 0
	for {
		n, _ := c.t.Read(buf) // Use raw read to find issue #15.
		output += string(buf[:n])

		// sanity
		if len(output) == 0 {
			i++
		}
		if i > 10 {
			break
		}

		if strings.Contains(output, linebreak) {
			break
		}
	}

	return strings.Replace(output, linebreak, "", 1), nil
}

func (c *Client) Close() {
	err := c.t.Close()
	if err != nil {
		c.logger.Error(err)
	}
}
