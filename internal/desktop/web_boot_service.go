package desktop

import (
	"fmt"
	"github.com/Akkadius/spire/internal/http"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
)

type WebBoot struct {
	logger *logrus.Logger
	router *routes.Router
}

func NewWebBoot(logger *logrus.Logger, router *routes.Router) *WebBoot {
	return &WebBoot{logger: logger, router: router}
}

func (c *WebBoot) Boot() {
	// get free network port from OS
	port, err := getFreePort()
	if err != nil {
		c.logger.Fatal(err)
	}

	// start web server
	go func() {
		if err := http.Serve(uint(port), c.logger, c.router); err != nil {
			c.logger.WithError(err).Fatal(err.Error())
		}
	}()

	// open browser window
	openBrowser(fmt.Sprintf("http://localhost:%v", port))

	// wait for signal to kill
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	for _ = range ch {
		// sig is a ^C, handle it
		os.Exit(0)
	}

}

func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
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
