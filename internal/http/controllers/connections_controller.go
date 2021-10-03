package controllers

import (
	"errors"
	"fmt"
	"github.com/Akkadius/spire/internal/connection"
	"github.com/Akkadius/spire/internal/connection/contexts"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/request"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/models"
	"github.com/labstack/echo/v4"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
	"time"
)

type ConnectionsController struct {
	db                        *database.DatabaseResolver
	logger                    *logrus.Logger
	cache                     *gocache.Cache
	dbConnectionCreateService *connection.DbConnectionCreateService
	dbConnectionCheckService  *connection.DbConnectionCheckService
}

func NewConnectionsController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	cache *gocache.Cache,
	dbConnectionCreateService *connection.DbConnectionCreateService,
	dbConnectionCheckService *connection.DbConnectionCheckService,
) *ConnectionsController {
	return &ConnectionsController{
		db:                        db,
		logger:                    logger,
		cache:                     cache,
		dbConnectionCreateService: dbConnectionCreateService,
		dbConnectionCheckService:  dbConnectionCheckService,
	}
}

func (cc *ConnectionsController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodPost, "connection", cc.create, nil),
		routes.RegisterRoute(http.MethodGet, "connections", cc.list, nil),
		routes.RegisterRoute(http.MethodGet, "connection-check/:connection", cc.check, nil),
		routes.RegisterRoute(http.MethodPost, "connection/:connection/set-active", cc.setActive, nil),
		routes.RegisterRoute(http.MethodDelete, "connection/:connection", cc.delete, nil),

		// default connection(s)
		routes.RegisterRoute(http.MethodGet, "connection-default", cc.getDefault, nil),
		routes.RegisterRoute(http.MethodPost, "connection-default/set-active", cc.defaultSetActive, nil),
	}
}

// ConnectionCreateRequest
type ConnectionCreateRequest struct {
	ConnectionName    string `json:"connection_name"`
	DbName            string `json:"db_name"`
	DbHost            string `json:"db_host"`
	DbPort            string `json:"db_port"`
	DbUsername        string `json:"db_username"`
	DbPassword        string `json:"db_password"`
	ContentDbName     string `json:"content_db_name"`
	ContentDbHost     string `json:"content_db_host"`
	ContentDbPort     string `json:"content_db_port"`
	ContentDbUsername string `json:"content_db_username"`
	ContentDbPassword string `json:"content_db_password"`
}

func (cc *ConnectionsController) clearConnection(c echo.Context) {
	user := request.GetUser(c)
	cc.cache.Delete(fmt.Sprintf("active-connection-%v-default", user.ID))
	cc.cache.Delete(fmt.Sprintf("active-connection-%v-eqemu_content", user.ID))
}

// ConnectionCreateRequest handle
func (cc *ConnectionsController) create(c echo.Context) error {
	user := request.GetUser(c)

	r := new(ConnectionCreateRequest)
	if err := c.Bind(r); err != nil {
		return err
	}

	// context
	ctx, err := contexts.NewConnectionCreateContext(
		user.ID,
		r.ConnectionName,
		r.DbName,
		r.DbHost,
		r.DbPort,
		r.DbUsername,
		r.DbPassword,
	)
	if err != nil {
		return err
	}

	// If any content params are set
	if len(r.ContentDbName) > 0 {
		ctx.SetContentDbName(r.ContentDbName)
	}
	if len(r.ContentDbHost) > 0 {
		ctx.SetContentDbHost(r.ContentDbHost)
	}
	if len(r.ContentDbUsername) > 0 {
		ctx.SetContentDbUsername(r.ContentDbUsername)
	}
	if len(r.ContentDbPassword) > 0 {
		ctx.SetContentDbPassword(r.ContentDbPassword)
	}
	if len(r.ContentDbPort) > 0 {
		ctx.SetContentDbPort(r.ContentDbPort)
	}

	// created address
	remoteAddr := strings.Split(c.Request().RemoteAddr, ":")
	ipAddress := ""
	if len(remoteAddr) > 0 {
		ipAddress = remoteAddr[0]
	}
	ctx.SetCreatedFromIp(ipAddress)

	err = cc.dbConnectionCreateService.Handle(ctx)
	if err != nil {
		return err
	}

	cc.clearConnection(c)

	return c.JSON(http.StatusOK, echo.Map{"data": "Connection created successfully"})
}

func (cc *ConnectionsController) list(c echo.Context) error {
	var results []models.UserServerDatabaseConnection
	relationships := models.UserServerDatabaseConnection{}.Relationships()
	query := cc.db.GetSpireDb().Model(&models.UserServerDatabaseConnection{})
	for _, relationship := range relationships {
		query = query.Preload(relationship)
	}

	err := query.Where("user_id = ?", request.GetUser(c).ID).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"data": results})
}

func (cc *ConnectionsController) check(c echo.Context) error {
	err := cc.dbConnectionCheckService.Handle(request.GetUser(c).ID, c.Param("connection"))
	if err != nil {
		return c.JSON(
			http.StatusOK,
			echo.Map{"data": echo.Map{"message": err.Error()}},
		)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": echo.Map{"message": "Server online"}})
}

func (cc *ConnectionsController) setActive(c echo.Context) error {
	connectionId := c.Param("connection")
	user := request.GetUser(c)
	db := cc.db.GetSpireDb()

	var conn models.UserServerDatabaseConnection
	err := db.Where("user_id = ?", user.ID).First(&conn, connectionId).Error
	if err != nil {
		return err
	}

	if conn.ID == 0 {
		return errors.New("cannot find connection")
	}

	q := db.Model(models.UserServerDatabaseConnection{})
	q.Where("user_id = ?", user.ID).Update("active", "0")
	q.Where("user_id = ? and id = ?", user.ID, connectionId).Update("active", "1")

	cc.clearConnection(c)

	return c.JSON(http.StatusOK, echo.Map{"data": echo.Map{"message": "Success"}})
}

func (cc *ConnectionsController) delete(c echo.Context) error {
	connectionId := c.Param("connection")
	user := request.GetUser(c)
	db := cc.db.GetSpireDb()

	var conn models.UserServerDatabaseConnection
	err := cc.db.GetSpireDb().Where("user_id = ? and id = ?", user.ID, connectionId).First(&conn).Error
	if err != nil {
		return err
	}

	db.Delete(models.UserServerDatabaseConnection{}, connectionId)
	db.Where("id = ?", conn.ServerDatabaseConnectionId).Delete(models.ServerDatabaseConnection{})

	cc.clearConnection(c)

	return c.JSON(http.StatusOK, echo.Map{"data": echo.Map{"message": "Success"}})
}

// gets default connection
func (cc *ConnectionsController) getDefault(c echo.Context) error {
	result := models.UserServerDatabaseConnection{
		ServerDatabaseConnection: models.ServerDatabaseConnection{
			ID:         0,
			Name:       "ProjectEQ Server Database (Read Only) (Local)",
			DbHost:     os.Getenv("MYSQL_EQEMU_HOST"),
			DbPort:     os.Getenv("MYSQL_EQEMU_PORT"),
			DbName:     os.Getenv("MYSQL_EQEMU_DATABASE"),
			DbUsername: os.Getenv("MYSQL_EQEMU_USERNAME"),
			DbPassword: os.Getenv("MYSQL_EQEMU_PASSWORD"),
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
		},
	}

	return c.JSON(http.StatusOK, echo.Map{"data": result})
}

// set default connection active
func (cc *ConnectionsController) defaultSetActive(c echo.Context) error {
	user := request.GetUser(c)
	db := cc.db.GetSpireDb()

	q := db.Model(models.UserServerDatabaseConnection{})
	q.Where("user_id = ?", user.ID).Update("active", "0")

	cc.clearConnection(c)

	return c.JSON(http.StatusOK, echo.Map{"data": echo.Map{"message": "Success"}})
}
