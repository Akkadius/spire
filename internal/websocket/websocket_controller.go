package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
	"net/http"
)

type Controller struct {
	pathmgmt *pathmgmt.PathManagement
	handler  *Handler
	manager  *ClientManager
	logger   *logger.AppLogger
}

func NewController(
	pathmgmt *pathmgmt.PathManagement,
	handler *Handler,
	manager *ClientManager,
	logger *logger.AppLogger,
) *Controller {
	return &Controller{
		pathmgmt: pathmgmt,
		handler:  handler,
		manager:  manager,
		logger:   logger,
	}
}

func (a *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "websocket", a.websocketHandler, nil),
	}
}

type SpireWebsocketMessage struct {
	Action string `json:"action"`
}

func (a *Controller) websocketHandler(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Register client
			remoteAddress := ws.Request().RemoteAddr
			uniqueID := fmt.Sprintf("%s", remoteAddress)
			client := &Client{WS: ws, ID: uniqueID} // Assign a real unique ID
			a.manager.register <- client

			// Read from client
			msg := ""
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				// close if error
				if err := ws.Close(); err != nil {
					a.manager.unregister <- client
					break
				}
			}

			// just close empty messages
			if len(msg) == 0 {
				// close if error
				if err := ws.Close(); err != nil {
					a.manager.unregister <- client
					break
				}
			}

			// Write
			if len(msg) > 0 {
				//a.manager.broadcast <- msg

				a.logger.Debug().Any("msg", msg).Msg("Received message")

				var err error
				var m SpireWebsocketMessage
				err = json.Unmarshal([]byte(msg), &m)
				if err != nil {
					fmt.Println(err)
				}

				if m.Action == "hello" {
					err = a.handler.HandleHello(ws, msg)
				}
				// this was a POC anyhow
				//if m.Action == "exec_server_bin" {
				//	err = a.handler.HandleExecServerBin(ws, msg)
				//}
			}

			// close connection on error
			// improve this later but this works well enough for now
			if err != nil {
				if err := ws.Close(); err != nil {
					a.manager.unregister <- client
					break
				}
			}

		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
