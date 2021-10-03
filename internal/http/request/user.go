package request

import (
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) models.User {
	user, ok := c.Get("user").(models.User)
	if !ok {
		return models.User{}
	}

	return user
}
