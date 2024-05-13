package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/http/request"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
	"net/http"
)

type Controller struct {
	pathmgmt *pathmgmt.PathManagement
	handler  *Handler
}

func NewController(
	pathmgmt *pathmgmt.PathManagement,
	handler *Handler,
) *Controller {
	return &Controller{
		pathmgmt: pathmgmt,
		handler:  handler,
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
	user := request.GetUser(c)

	websocket.Handler(func(ws *websocket.Conn) {

		// if no user authorized, kill connection
		if user.ID == 0 {
			_ = a.handler.HandleUnauthorized(ws)
			_ = ws.Close()
		}

		defer ws.Close()
		for {

			// Read from client
			msg := ""
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				// close if error
				if err := ws.Close(); err != nil {
					break
				}
			}

			// just close empty messages
			if len(msg) == 0 {
				// close if error
				if err := ws.Close(); err != nil {
					break
				}
			}

			// Write
			if len(msg) > 0 {
				var err error
				var m SpireWebsocketMessage
				err = json.Unmarshal([]byte(msg), &m)
				if err != nil {
					fmt.Println(err)
				}

				if m.Action == "hello" {
					err = a.handler.HandleHello(ws, msg)
				}
				if m.Action == "exec_server_bin" {
					err = a.handler.HandleExecServerBin(ws, msg)
				}
			}

			// close connection on error
			// improve this later but this works well enough for now
			if err != nil {
				if err := ws.Close(); err != nil {
					break
				}
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
