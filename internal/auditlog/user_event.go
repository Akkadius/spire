package auditlog

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/request"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

type UserEvent struct {
	db     *database.DatabaseResolver
	cache  *cache.Cache
	logger *logrus.Logger
}

func NewUserEvent(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	cache *cache.Cache,
) *UserEvent {
	return &UserEvent{
		cache:  cache,
		db:     db,
		logger: logger,
	}
}

func (e UserEvent) LogUserEvent(c echo.Context, eventName string, description string) {
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
	}
}
