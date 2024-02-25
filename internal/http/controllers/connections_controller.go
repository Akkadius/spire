package controllers

import (
	"errors"
	"fmt"
	"github.com/Akkadius/spire/internal/connection"
	"github.com/Akkadius/spire/internal/connection/contexts"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/request"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/Akkadius/spire/internal/permissions"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/Akkadius/spire/internal/user"
	"github.com/labstack/echo/v4"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type ConnectionsController struct {
	db                        *database.Resolver
	logger                    *logrus.Logger
	cache                     *gocache.Cache
	dbConnectionCreateService *connection.Create
	dbConnectionCheckService  *connection.Check
	permissions               *permissions.Service
	spireInit                 *spire.Init
	spireuser                 *user.User
}

func NewConnectionsController(
	db *database.Resolver,
	logger *logrus.Logger,
	cache *gocache.Cache,
	dbConnectionCreateService *connection.Create,
	dbConnectionCheckService *connection.Check,
	permissions *permissions.Service,
	spireInit *spire.Init,
	spireuser *user.User,
) *ConnectionsController {
	return &ConnectionsController{
		db:                        db,
		logger:                    logger,
		cache:                     cache,
		permissions:               permissions,
		dbConnectionCreateService: dbConnectionCreateService,
		dbConnectionCheckService:  dbConnectionCheckService,
		spireInit:                 spireInit,
		spireuser:                 spireuser,
	}
}

func (cc *ConnectionsController) Routes() []*routes.Route {
	return []*routes.Route{

		// connection
		routes.RegisterRoute(http.MethodPost, "connection", cc.create, nil),
		routes.RegisterRoute(http.MethodGet, "connections", cc.list, nil),
		routes.RegisterRoute(http.MethodGet, "connection-check/:connection", cc.check, nil),
		routes.RegisterRoute(http.MethodPost, "connection/:connection/set-active", cc.setActive, nil),
		routes.RegisterRoute(http.MethodDelete, "connection/:connection", cc.delete, nil),
		routes.RegisterRoute(http.MethodGet, "connection/:connection_id/audit-log", cc.auditLog, nil),

		// user add / remove
		routes.RegisterRoute(http.MethodPost, "connection/:connection_id/add-user/:user_id", cc.addUser, nil),
		routes.RegisterRoute(http.MethodDelete, "connection/:connection_id/add-user/:user_id", cc.deleteUser, nil),

		// permission
		routes.RegisterRoute(http.MethodPost, "connection-permissions/:connection_id/user/:user_id", cc.savePermissions, nil),
		routes.RegisterRoute(http.MethodGet, "connection-permissions/:connection_id/user/:user_id", cc.getPermissions, nil),

		// discord
		routes.RegisterRoute(http.MethodPost, "connection/:connection_id/discord-webhook", cc.setDiscordWebhook, nil),
		routes.RegisterRoute(http.MethodGet, "connection/:connection_id/discord-webhook", cc.getDiscordWebhook, nil),

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
	LogsDbName        string `json:"logs_db_name"`
	LogsDbHost        string `json:"logs_db_host"`
	LogsDbPort        string `json:"logs_db_port"`
	LogsDbUsername    string `json:"logs_db_username"`
	LogsDbPassword    string `json:"logs_db_password"`
}

func (cc *ConnectionsController) clearConnection(c echo.Context) {
	u := request.GetUser(c)
	cc.spireuser.PurgeUserCache(u.ID)
}

// ConnectionCreateRequest handle
func (cc *ConnectionsController) create(c echo.Context) error {
	u := request.GetUser(c)

	r := new(ConnectionCreateRequest)
	if err := c.Bind(r); err != nil {
		return err
	}

	// context
	ctx, err := contexts.NewConnectionCreateContext(
		u.ID,
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

	// If any logs params are set
	if len(r.LogsDbName) > 0 {
		ctx.SetLogsDbName(r.LogsDbName)
	}
	if len(r.LogsDbHost) > 0 {
		ctx.SetLogsDbHost(r.LogsDbHost)
	}
	if len(r.LogsDbUsername) > 0 {
		ctx.SetLogsDbUsername(r.LogsDbUsername)
	}
	if len(r.LogsDbPassword) > 0 {
		ctx.SetLogsDbPassword(r.LogsDbPassword)
	}
	if len(r.LogsDbPort) > 0 {
		ctx.SetLogsDbPort(r.LogsDbPort)
	}

	// created address
	remoteAddr := strings.Split(c.Request().RemoteAddr, ":")
	ipAddress := ""
	if len(remoteAddr) > 0 {
		ipAddress = remoteAddr[0]
	}
	ctx.SetCreatedFromIp(ipAddress)
	ctx.SetCheckSecondaryDbConnection(true)
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
	if cc.db.GetSpireDb() != nil {
		query := cc.db.GetSpireDb().Model(&models.UserServerDatabaseConnection{})
		for _, relationship := range relationships {
			query = query.Preload(relationship)
		}

		err := query.Or("user_id = ?", request.GetUser(c).ID).Find(&results).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
		}
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
	q = db.Model(models.UserServerDatabaseConnection{})
	q.Where("user_id = ? and id = ?", user.ID, connectionId).Update("active", 1)

	cc.clearConnection(c)

	return c.JSON(http.StatusOK, echo.Map{"data": echo.Map{"message": "Success"}})
}

func (cc *ConnectionsController) delete(c echo.Context) error {
	user := request.GetUser(c)
	db := cc.db.GetSpireDb()

	connectionId, err := strconv.ParseUint(c.Param("connection"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	// local setups - prevent from deleting main connection
	if cc.spireInit.IsAuthEnabled() && connectionId == 1 {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": "Cannot delete this connection, it is tied to your server configuration"},
		)
	}

	// Validate: Invoking user is owner of connection
	var conn models.ServerDatabaseConnection
	_ = cc.db.GetSpireDb().Where("created_by = ? and id = ?", user.ID, connectionId).First(&conn).Error
	if uint64(conn.ID) != connectionId {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invoking user does not own this connection"})
	}

	// delete users associated to connection
	db.Where("server_database_connection_id = ?", conn.ID).Delete(models.UserServerDatabaseConnection{})
	// delete event logs
	db.Where("server_database_connection_id = ?", conn.ID).Delete(models.UserEventLog{})
	// delete permissions
	db.Where("server_database_connection_id = ?", conn.ID).Delete(models.UserServerResourcePermission{})
	// delete connection itself
	db.Where("id = ?", conn.ID).Delete(models.ServerDatabaseConnection{})

	// clear cache
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

func (cc *ConnectionsController) addUser(c echo.Context) error {

	// request user context
	ctx := request.GetUser(c)

	// params
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}
	connectionId, err := strconv.ParseUint(c.Param("connection_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	// Validate: Invoking user is owner of connection
	var conn models.ServerDatabaseConnection
	_ = cc.db.GetSpireDb().Where("created_by = ? and id = ?", ctx.ID, connectionId).First(&conn).Error
	if uint64(conn.ID) != connectionId {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invoking user does not own this connection"})
	}

	// Validate: User inquired is valid
	var user models.User
	err = cc.db.GetSpireDb().Where("id = ?", userId).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}
	if uint64(user.ID) != userId {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "User does not exist"})
	}

	// Validate: User doesn't already exist on connection
	var connUser models.UserServerDatabaseConnection
	_ = cc.db.GetSpireDb().Where("user_id = ? and server_database_connection_id = ?", userId, connectionId).First(&connUser).Error
	if connUser.UserId > 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "User already exists on connection!"})
	}

	// associate user with connection
	var userConnection models.UserServerDatabaseConnection
	userConnection.UserId = uint(userId)
	userConnection.Active = 0
	userConnection.ServerDatabaseConnectionId = uint(connectionId)
	userConnection.CreatedBy = ctx.ID
	err = cc.db.GetSpireDb().Create(&userConnection).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"data": echo.Map{"message": "Success"}})
}

func (cc *ConnectionsController) deleteUser(c echo.Context) error {

	// request user context
	ctx := request.GetUser(c)

	// TODO: Consolidate validation for add/delete as they are similar

	// params
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}
	connectionId, err := strconv.ParseUint(c.Param("connection_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	// Validate: Invoking user is owner of connection
	var conn models.ServerDatabaseConnection
	_ = cc.db.GetSpireDb().Where("created_by = ? and id = ?", ctx.ID, connectionId).First(&conn).Error
	if uint64(conn.ID) != connectionId {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invoking user does not own this connection"})
	}

	// Validate: User inquired is valid
	var user models.User
	err = cc.db.GetSpireDb().Where("id = ?", userId).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}
	if uint64(user.ID) != userId {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "User does not exist"})
	}

	// Validate: User doesn't already exist on connection
	var connUser models.UserServerDatabaseConnection
	_ = cc.db.GetSpireDb().Where("user_id = ? and server_database_connection_id = ?", userId, connectionId).First(&connUser).Error
	if connUser.UserId == 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "User does not exists on connection!"})
	}

	// Validate: user to be removed is not the owner
	var serverDatabaseConn models.ServerDatabaseConnection
	_ = cc.db.GetSpireDb().Where("created_by = ? and id = ?", userId, connectionId).First(&serverDatabaseConn).Error
	if serverDatabaseConn.ID > 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot delete owner of connection"})
	}

	// Delete
	err = cc.db.GetSpireDb().
		Where("user_id = ? and server_database_connection_id = ?", userId, connectionId).
		Delete(&models.UserServerDatabaseConnection{}).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	// Delete entries first
	err = cc.db.GetSpireDb().
		Where("user_id = ? and server_database_connection_id = ?", userId, connectionId).
		Delete(&models.UserServerResourcePermission{}).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"data": echo.Map{"message": "Success"}})
}

func (cc *ConnectionsController) savePermissions(c echo.Context) error {

	// request user context
	ctx := request.GetUser(c)

	// params
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}
	connectionId, err := strconv.ParseUint(c.Param("connection_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	// Validate: Invoking user is owner of connection
	var conn models.ServerDatabaseConnection
	_ = cc.db.GetSpireDb().Where("created_by = ? and id = ?", ctx.ID, connectionId).First(&conn).Error
	if uint64(conn.ID) != connectionId {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invoking user does not own this connection"})
	}

	// Validate: User inquired is valid
	var user models.User
	err = cc.db.GetSpireDb().Where("id = ?", userId).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}
	if uint64(user.ID) != userId {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "User does not exist"})
	}

	type Permissions []struct {
		Permission string `json:"permission"`
		Read       bool   `json:"read"`
		Write      bool   `json:"write"`
	}

	var permissions Permissions
	if err := c.Bind(&permissions); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var addPermissions []models.UserServerResourcePermission
	for _, permission := range permissions {
		var p models.UserServerResourcePermission
		p.ServerDatabaseConnectionId = uint(connectionId)
		p.ResourceName = permission.Permission

		canRead := 0
		if permission.Read {
			canRead = 1
		}

		canWrite := 0
		if permission.Write {
			canWrite = 1
		}

		p.UserId = uint(userId)
		p.CanRead = uint8(canRead)
		p.CanWrite = uint8(canWrite)

		addPermissions = append(addPermissions, p)
	}

	// Delete entries first
	err = cc.db.GetSpireDb().
		Where("user_id = ? and server_database_connection_id = ?", userId, connectionId).
		Delete(&models.UserServerResourcePermission{}).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	// Bulk Insert
	err = cc.db.GetSpireDb().
		CreateInBatches(addPermissions, 1000).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	// clear cache
	cc.permissions.ClearUserPermissionsCache(userId)

	return c.JSON(http.StatusOK, echo.Map{"data": echo.Map{"message": "Success"}})
}

func (cc *ConnectionsController) getPermissions(c echo.Context) error {
	// request user context
	ctx := request.GetUser(c)

	// params
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}
	connectionId, err := strconv.ParseUint(c.Param("connection_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	// Validate: Invoking user is member of connection
	var userConn models.UserServerDatabaseConnection
	_ = cc.db.GetSpireDb().Where("user_id = ? and server_database_connection_id = ?", ctx.ID, connectionId).First(&userConn).Error

	isOwnerOfConnection := userConn.CreatedBy == ctx.ID

	// Validate: Connection inquired is valid
	var conn models.ServerDatabaseConnection
	_ = cc.db.GetSpireDb().Where("id = ?", connectionId).First(&conn).Error
	if uint64(conn.ID) != connectionId && uint64(userConn.ID) != connectionId {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Connection requested is invalid"})
	}

	// Validate: User inquired is valid
	var user models.User
	_ = cc.db.GetSpireDb().Where("id = ?", userId).First(&user).Error
	if uint64(user.ID) != userId {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "User does not exist"})
	}

	if !isOwnerOfConnection && userId != uint64(ctx.ID) {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "User does not own connection or is not user itself"})
	}

	var results []models.UserServerResourcePermission
	if cc.db.GetSpireDb() != nil {
		query := cc.db.GetSpireDb().Model(&models.UserServerResourcePermission{})

		err := query.Where("user_id = ? and server_database_connection_id = ?", userId, connectionId).Find(&results).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
		}

		return c.JSON(http.StatusOK, results)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": echo.Map{"message": "Success"}})
}

func (cc *ConnectionsController) auditLog(c echo.Context) error {
	// request user context
	ctx := request.GetUser(c)

	// param
	connectionId, err := strconv.ParseUint(c.Param("connection_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	// Validate: Invoking user is owner of connection
	var conn models.ServerDatabaseConnection
	_ = cc.db.GetSpireDb().Where("created_by = ? and id = ?", ctx.ID, connectionId).First(&conn).Error
	if uint64(conn.ID) != connectionId {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invoking user does not own this connection"})
	}

	// paging
	pageSize := 20
	queryOffset := 0
	page, err := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}
	if page > 0 {
		queryOffset = int(page - 1)
	}

	// query
	var results []models.UserEventLog
	query := cc.db.GetSpireDb().Model(&models.UserEventLog{}).
		Preload("User").
		Where("server_database_connection_id = ?", connectionId).
		Limit(pageSize).
		Offset(queryOffset * pageSize).
		Order("id DESC")
	_ = query.Find(&results).Error

	// count
	var count int64
	cc.db.GetSpireDb().Model(&models.UserEventLog{}).
		Where("server_database_connection_id = ?", connectionId).
		Distinct("id").Count(&count)

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"limit":      pageSize,
			"total_rows": count,
			"data":       results,
		},
	)
}

type DiscordWebhookRequest struct {
	WebhookUrl string `json:"webhook_url"`
}

func (cc *ConnectionsController) setDiscordWebhook(c echo.Context) error {
	// request user context
	ctx := request.GetUser(c)

	// param
	connectionId, err := strconv.ParseUint(c.Param("connection_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	// Validate: Invoking user is owner of connection
	var conn models.ServerDatabaseConnection
	_ = cc.db.GetSpireDb().Where("created_by = ? and id = ?", ctx.ID, connectionId).First(&conn).Error
	if uint64(conn.ID) != connectionId {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invoking user does not own this connection"})
	}

	webhookUrl := new(DiscordWebhookRequest)
	if err := c.Bind(webhookUrl); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	if len(webhookUrl.WebhookUrl) > 0 {
		conn.DiscordWebhookUrl = webhookUrl.WebhookUrl
		_ = cc.db.GetSpireDb().
			Model(&conn).
			Where("created_by = ? and id = ?", ctx.ID, connectionId).
			Update("discord_webhook_url", webhookUrl.WebhookUrl)

		return c.JSON(
			http.StatusOK,
			"Webhook updated successfully!",
		)
	}

	return c.JSON(
		http.StatusInternalServerError,
		echo.Map{"error": "Failed to update discord webhook"},
	)
}

func (cc *ConnectionsController) getDiscordWebhook(c echo.Context) error {
	// request user context
	ctx := request.GetUser(c)

	// param
	connectionId, err := strconv.ParseUint(c.Param("connection_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	// Validate: Invoking user is owner of connection
	var conn models.ServerDatabaseConnection
	_ = cc.db.GetSpireDb().Where("created_by = ? and id = ?", ctx.ID, connectionId).First(&conn).Error
	if uint64(conn.ID) != connectionId {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invoking user does not own this connection"})
	}

	return c.JSON(
		http.StatusOK,
		conn.DiscordWebhookUrl,
	)
}
