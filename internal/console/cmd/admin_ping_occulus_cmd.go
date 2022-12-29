package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/serverconfig"
	"github.com/k0kubun/pp/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"io"
	"net/http"
)

type AdminPingOcculus struct {
	db           *gorm.DB
	logger       *logrus.Logger
	command      *cobra.Command
	serverconfig *serverconfig.EQEmuServerConfig
}

func (c *AdminPingOcculus) Command() *cobra.Command {
	return c.command
}

func NewAdminPingOcculus(db *gorm.DB, logger *logrus.Logger, serverconfig *serverconfig.EQEmuServerConfig) *AdminPingOcculus {
	i := &AdminPingOcculus{
		db:           db,
		logger:       logger,
		serverconfig: serverconfig,
		command: &cobra.Command{
			Use:   "admin:ping-occulus",
			Short: "Pings Occulus",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *AdminPingOcculus) Handle(cmd *cobra.Command, args []string) {
	config := c.serverconfig.Get()
	//pp.Println(config.WebAdmin)

	type PostBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	postBody := PostBody{
		Username: "admin",
		Password: config.WebAdmin.Application.Admin.Password,
	}

	payload, err := json.Marshal(postBody)
	if err != nil {
		c.logger.Error(err)
	}

	resp, err := http.Post(
		"http://192.168.65.115:3000/api/v1/auth/login", "application/json",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		c.logger.Error(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error(err)
	}

	type LoginResponse struct {
		Success     string `json:"success"`
		AccessToken string `json:"access_token"`
	}

	var loginResponse LoginResponse
	err = json.Unmarshal(b, &loginResponse)
	if err != nil {
		c.logger.Error(err)
	}

	pp.Println(loginResponse)

	req, err := http.NewRequest("GET", "http://192.168.65.115:3000/api/v1/server/process_counts", nil)
	if err != nil {
		c.logger.Error(err)
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", loginResponse.AccessToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.logger.Error(err)
	}

	defer res.Body.Close()

	b, err = io.ReadAll(res.Body)
	if err != nil {
		c.logger.Error(err)
	}

	pp.Println(string(b))
}

// Validate implementation of the Command interface
func (c *AdminPingOcculus) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
