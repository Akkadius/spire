package deploy

import (
	"github.com/Akkadius/spire/internal/discord"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type DeployController struct {
	logger *logrus.Logger
}

func NewDeployController(
	logger *logrus.Logger,
) *DeployController {
	return &DeployController{
		logger: logger,
	}
}

func (a *DeployController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodPost, "spire-deploy/:key", a.deploy, nil),
	}
}

func (a *DeployController) deploy(c echo.Context) error {
	if env.IsAppEnvProduction() &&
		env.Get("DEPLOY_KEY", "") == c.Param("key") {

		// run git pull
		cmd := exec.Command("git", "pull", "origin", "master")
		output, err := cmd.Output()
		if err != nil {
			log.Println(err)
		}

		if !strings.Contains(string(output), "Already up to date") {

			// webhook
			webhookUrl := env.Get("DISCORD_DEPLOY_WEBHOOK", "")
			if len(webhookUrl) > 0 {
				err := discord.SendDiscordWebhook(webhookUrl, "Hosted Spire is being deployed with latest master")
				if err != nil {
					return c.JSON(http.StatusInternalServerError, err.Error())
				}
			}

			// run git pull
			cmd := exec.Command("bash", "-c", "kill -9 1")
			_, err := cmd.Output()
			if err != nil {
				log.Println(err)
			}

			os.Exit(0)
		}

		return c.JSON(http.StatusOK, echo.Map{"data": "Ok"})
	}

	return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Key failed"})
}
