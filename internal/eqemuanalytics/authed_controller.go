package eqemuanalytics

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/discord"
	"github.com/Akkadius/spire/internal/http/request"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"net/http"
	"os"
	"time"
)

type AuthedController struct {
	db *database.Resolver
}

func NewAuthedController(
	db *database.Resolver,
) *AuthedController {
	return &AuthedController{
		db: db,
	}
}

func (a *AuthedController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodPost, "analytics/server-crash-reports/:crash-id/mark-resolved", a.markCrashResolved, nil),
		routes.RegisterRoute(http.MethodPost, "analytics/server-crash-reports/:crash-id/mark-unresolved", a.markCrashUnresolved, nil),
	}
}

func (a *AuthedController) markCrashResolved(c echo.Context) error {
	crashId := c.Param("crash-id")
	if len(crashId) == 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invalid crash report request"})
	}

	// request user context
	user := request.GetUser(c)
	if user.ID == 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invalid user"})
	}

	if !user.IsServerDeveloper {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Not server developer"})
	}

	var r models.CrashReport
	a.db.GetSpireDb().Find(&r, crashId)
	if r.ID == 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invalid crash report"})
	}

	// update all crash reports with this fingerprint to resolved
	result := a.db.GetSpireDb().
		Model(&models.CrashReport{}).
		Where("fingerprint = ?", r.Fingerprint).
		Updates(models.CrashReport{Resolved: true, ResolvedBy: user.ID, ResolvedAt: null.TimeFrom(time.Now())})

	// send discord webhook
	go func() {
		link := fmt.Sprintf(
			"%s/dev/release/%s?id=%d",
			os.Getenv("VUE_APP_FRONTEND_BASE_URL"),
			r.ServerVersion,
			r.ID,
		)

		err := discord.SendDiscordWebhook(
			os.Getenv("DISCORD_CRASH_REPORT_WEBHOOK_URL"),
			fmt.Sprintf(
				"✅ Crash fingerprint **%v** (%v total crashes) marked as resolved by **%v** can be viewed at %v",
				r.Fingerprint,
				result.RowsAffected,
				user.UserName,
				link,
			),
		)
		if err != nil {
			fmt.Println(err)
		}
	}()

	return c.JSON(http.StatusOK, echo.Map{"data": "Crash fingerprint marked as resolved"})
}

func (a *AuthedController) markCrashUnresolved(c echo.Context) error {
	crashId := c.Param("crash-id")
	if len(crashId) == 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invalid crash report request"})
	}

	// request user context
	user := request.GetUser(c)
	if user.ID == 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invalid user"})
	}

	if !user.IsServerDeveloper {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Not server developer"})
	}

	var r models.CrashReport
	a.db.GetSpireDb().Find(&r, crashId)
	if r.ID == 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Invalid crash report"})
	}

	// update all crash reports with this fingerprint to resolved
	result := a.db.GetSpireDb().
		Model(&models.CrashReport{}).
		Where("fingerprint = ?", r.Fingerprint).
		Updates(map[string]interface{}{"resolved": 0, "resolved_by": 0, "resolved_at": nil})

	// send discord webhook
	go func() {
		link := fmt.Sprintf(
			"%s/dev/release/%s?id=%d",
			os.Getenv("VUE_APP_FRONTEND_BASE_URL"),
			r.ServerVersion,
			r.ID,
		)

		err := discord.SendDiscordWebhook(
			os.Getenv("DISCORD_CRASH_REPORT_WEBHOOK_URL"),
			fmt.Sprintf(
				"❌ Crash fingerprint **%v** (%v total crashes) marked as resolved by **%v** can be viewed at %v",
				r.Fingerprint,
				result.RowsAffected,
				user.UserName,
				link,
			),
		)
		if err != nil {
			fmt.Println(err)
		}
	}()

	return c.JSON(http.StatusOK, echo.Map{"data": "Crash fingerprint marked as unresolved"})
}
