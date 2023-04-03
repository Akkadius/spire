package desktop

import (
	"fmt"
	"github.com/Akkadius/spire/internal/http"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	gohttp "net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strconv"
	"time"
)

type WebBoot struct {
	logger *logrus.Logger
	server *http.Server
}

func NewWebBoot(logger *logrus.Logger, server *http.Server) *WebBoot {
	return &WebBoot{logger: logger, server: server}
}

func (c *WebBoot) Boot() {
	port := 0

	// get free network port from OS
	for i := 8090; i <= 8099; i++ {
		found, err := checkIfPortAvailable(i)
		if found && err == nil {
			port = i
			break
		}
	}

	if port == 0 {
		fmt.Println("Failed to find free port, exiting...")
		os.Exit(1)
	}

	// start web server
	go func() {
		if err := c.server.Serve(uint(port)); err != nil {
			c.logger.WithError(err).Fatal(err.Error())
		}
	}()

	// open browser window
	web := fmt.Sprintf("http://localhost:%v", port)
	err := waitForSiteToBeAvailable(web, time.Minute*15)
	if err != nil {
		c.logger.Fatal(err)
	}
	openBrowser(web)

	// wait for signal to kill
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	for _ = range ch {
		// sig is a ^C, handle it
		os.Exit(0)
	}
}

func checkIfPortAvailable(port int) (status bool, err error) {
	// Concatenate a colon and the port
	host := ":" + strconv.Itoa(port)

	// Try to create a server with the port
	server, err := net.Listen("tcp", host)

	// if it fails then the port is likely taken
	if err != nil {
		return false, err
	}

	// close the server
	server.Close()

	// we successfully used and closed the port
	// so it's now available to be used again
	return true, nil

}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		// only try to open a browser window if there is a desktop environment present
		if len(os.Getenv("XDG_CURRENT_DESKTOP")) > 0 {
			err = exec.Command("xdg-open", url).Start()
		}
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func waitForSiteToBeAvailable(URL string, timeout time.Duration) error {
	ch := make(chan bool)
	go func() {
		for {
			_, err := gohttp.Get(URL)
			if err == nil {
				ch <- true
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()

	select {
	case <-ch:
		return nil
	case <-time.After(timeout):
		return fmt.Errorf("server did not reply after %v", timeout)
	}
}
