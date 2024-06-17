package auditlog

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/discord"
	"github.com/Akkadius/spire/internal/http/request"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"os"
)

type UserEvent struct {
	db    *database.Resolver
	cache *cache.Cache
}

func NewUserEvent(
	db *database.Resolver,
	cache *cache.Cache,
) *UserEvent {
	return &UserEvent{
		cache: cache,
		db:    db,
	}
}

const (
	EventServerLock          = "SERVER_LOCK"
	EventServerUpdateRelease = "SERVER_UPDATE_RELEASE"
	EventServerStart         = "SERVER_START"
	EventServerStop          = "SERVER_STOP"
	EventServerRestart       = "SERVER_RESTART"
	EventServerCancelRestart = "SERVER_CANCEL_RESTART"
	EventServerHotReload     = "SERVER_HOT_RELOAD"
	EventServerUpdateConfig  = "SERVER_UPDATE_CONFIG"
)

func (e *UserEvent) LogUserEvent(c echo.Context, eventName string, description string) {
	// request user context
	user := request.GetUser(c)

	// create
	var event models.UserEventLog
	event.UserId = user.ID

	// key for the users database connection identifier
	connectionIdKey := fmt.Sprintf("active-connection-%v", user.ID)
	cachedConn, found := e.cache.Get(connectionIdKey)

	// found cached connection
	if found {
		connectionId := cachedConn.(uint)
		event.ServerDatabaseConnectionId = connectionId
	}

	if event.ServerDatabaseConnectionId > 0 {
		event.EventName = eventName
		event.Data = description
		_ = e.db.GetSpireDb().Create(&event)

		conn := e.db.GetUserConnection(user)
		if len(conn.ServerDatabaseConnection.DiscordWebhookUrl) > 0 {
			// connection webhook
			name := conn.ServerDatabaseConnection.Name
			_ = discord.SendDiscordWebhook(
				conn.ServerDatabaseConnection.DiscordWebhookUrl,
				fmt.Sprintf("[**Spire Audit Log**] [**%v**] [**%v**] %v", name, user.UserName, description),
			)
		}

		// monitor webhook
		if conn.ServerDatabaseConnection.ID > 0 {
			name := conn.ServerDatabaseConnection.Name
			monitor := os.Getenv("DISCORD_MONITOR_URL")
			if len(monitor) > 0 {
				_ = discord.SendDiscordWebhook(
					monitor,
					fmt.Sprintf("[**Spire Audit Log**] [**%v**] [**%v**] %v", name, user.UserName, description),
				)
			}
		}
	}
}
