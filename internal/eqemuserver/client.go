package eqemuserver

import (
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/telnet"
	"strings"
)

type Client struct {
	telnet *telnet.Client
	logger *logger.AppLogger

	pool                []*telnet.Client
	lastUsedClientIndex int
}

func NewClient(t *telnet.Client, logger *logger.AppLogger) *Client {
	poolSize := 50
	var pool []*telnet.Client
	for i := 0; i < poolSize; i++ {
		pool = append(pool, telnet.NewClient(logger))
	}

	return &Client{
		telnet: t,
		logger: logger,
		pool:   pool,
	}
}

// GetTelnetClient returns a telnet client from the pool
func (c *Client) GetTelnetClient() *telnet.Client {
	c.lastUsedClientIndex++
	if c.lastUsedClientIndex >= len(c.pool) {
		c.lastUsedClientIndex = 0
	}

	c.logger.DebugVvv().Any("lastUsedClientIndex", c.lastUsedClientIndex).Msg("GetTelnetClient")

	return c.pool[c.lastUsedClientIndex]
}

type WorldZoneList struct {
	Data []struct {
		BootingUp          bool   `json:"booting_up,omitempty"`
		ClientAddress      string `json:"client_address,omitempty"`
		ClientLocalAddress string `json:"client_local_address,omitempty"`
		ClientPort         int    `json:"client_port,omitempty"`
		CompileDate        string `json:"compile_date,omitempty"`
		CompileTime        string `json:"compile_time,omitempty"`
		CompileVersion     string `json:"compile_version,omitempty"`
		ID                 int    `json:"id,omitempty"`
		InstanceID         int    `json:"instance_id,omitempty"`
		IP                 string `json:"ip,omitempty"`
		IsStaticZone       bool   `json:"is_static_zone,omitempty"`
		LaunchName         string `json:"launch_name,omitempty"`
		LaunchedName       string `json:"launched_name,omitempty"`
		NumberPlayers      int    `json:"number_players,omitempty"`
		Port               int    `json:"port,omitempty"`
		PreviousZoneID     int    `json:"previous_zone_id,omitempty"`
		UUID               string `json:"uuid,omitempty"`
		ZoneID             int    `json:"zone_id,omitempty"`
		ZoneLongName       string `json:"zone_long_name,omitempty"`
		ZoneName           string `json:"zone_name,omitempty"`
		ZoneOsPid          int    `json:"zone_os_pid,omitempty"`
	} `json:"data"`
	ExecutionTime string `json:"execution_time"`
	Method        string `json:"method"`
}

func (c *Client) GetZoneList() (WorldZoneList, error) {
	o, err := c.GetTelnetClient().Command(
		telnet.CommandConfig{Command: "api get_zone_list", EnforceJson: true},
	)
	if err != nil {
		return WorldZoneList{}, err
	}

	var zoneList WorldZoneList
	err = json.Unmarshal([]byte(o), &zoneList)
	if err != nil {
		return WorldZoneList{}, err
	}

	return zoneList, nil
}

type ServerCountsResponse struct {
	Data struct {
		ClientCount *int `json:"client_count,omitempty"`
		ZoneCount   *int `json:"zone_count,omitempty"`
	} `json:"data"`
	ExecutionTime string `json:"execution_time"`
	Method        string `json:"method"`
}

func (c *Client) GetServerCounts() (ServerCountsResponse, error) {
	o, err := c.GetTelnetClient().Command(
		telnet.CommandConfig{Command: "api get_server_counts", EnforceJson: true},
	)
	if err != nil {
		return ServerCountsResponse{}, err
	}

	var zoneList ServerCountsResponse
	err = json.Unmarshal([]byte(o), &zoneList)
	if err != nil {
		return ServerCountsResponse{}, err
	}

	return zoneList, nil
}

type LockStatusResponse struct {
	Data struct {
		Locked bool `json:"locked"`
	} `json:"data"`
	ExecutionTime string `json:"execution_time"`
	Method        string `json:"method"`
}

// GetLockStatus returns the lock status of the server
func (c *Client) GetLockStatus() (bool, error) {
	o, err := c.GetTelnetClient().Command(
		telnet.CommandConfig{Command: "api lock_status", EnforceJson: true},
	)
	if err != nil {
		return false, err
	}

	if !strings.Contains(o, "locked") {
		return false, fmt.Errorf("invalid response from lock_status: %v", o)
	}

	var r LockStatusResponse
	err = json.Unmarshal([]byte(o), &r)
	if err != nil {
		return false, err
	}

	return r.Data.Locked, nil
}

// SetLockStatus returns the lock status of the server
func (c *Client) SetLockStatus(locked bool) error {
	command := "lock"
	if !locked {
		command = "unlock"
	}
	_, err := c.GetTelnetClient().Command(
		telnet.CommandConfig{Command: command},
	)
	if err != nil {
		return err
	}

	return nil
}

type WorldClientList struct {
	Data []struct {
		AccountID            int    `json:"account_id"`
		AccountName          string `json:"account_name,omitempty"`
		Admin                int    `json:"admin,omitempty"`
		Anon                 int    `json:"anon,omitempty"`
		CharacterID          int    `json:"character_id,omitempty"`
		Class                int    `json:"class,omitempty"`
		ClientVersion        int    `json:"client_version,omitempty"`
		Gm                   int    `json:"gm,omitempty"`
		GuildID              int    `json:"guild_id,omitempty"`
		ID                   int    `json:"id,omitempty"`
		Instance             int    `json:"instance,omitempty"`
		IP                   int    `json:"ip,omitempty"`
		IsLocalClient        bool   `json:"is_local_client,omitempty"`
		Level                int    `json:"level,omitempty"`
		Lfg                  bool   `json:"lfg,omitempty"`
		LfgComments          string `json:"lfg_comments,omitempty"`
		LfgFromLevel         int    `json:"lfg_from_level,omitempty"`
		LfgMatchFilter       bool   `json:"lfg_match_filter,omitempty"`
		LfgToLevel           int    `json:"lfg_to_level,omitempty"`
		LoginserverAccountID int    `json:"loginserver_account_id,omitempty"`
		LoginserverID        int    `json:"loginserver_id,omitempty"`
		LoginserverName      string `json:"loginserver_name,omitempty"`
		Name                 string `json:"name"`
		Online               int    `json:"online"`
		Race                 int    `json:"race,omitempty"`
		Server               struct {
			ClientAddress      string `json:"client_address,omitempty"`
			ClientLocalAddress string `json:"client_local_address,omitempty"`
			ClientPort         int    `json:"client_port,omitempty"`
			CompileTime        string `json:"compile_time,omitempty"`
			ID                 int    `json:"id,omitempty"`
			InstanceID         int    `json:"instance_id,omitempty"`
			IP                 string `json:"ip,omitempty"`
			IsBooting          bool   `json:"is_booting,omitempty"`
			LaunchName         string `json:"launch_name,omitempty"`
			LaunchedName       string `json:"launched_name,omitempty"`
			NumberPlayers      int    `json:"number_players,omitempty"`
			Port               int    `json:"port,omitempty"`
			PreviousZoneID     int    `json:"previous_zone_id,omitempty"`
			StaticZone         bool   `json:"static_zone,omitempty"`
			Uui                string `json:"uui,omitempty"`
			ZoneID             int    `json:"zone_id,omitempty"`
			ZoneLongName       string `json:"zone_long_name,omitempty"`
			ZoneName           string `json:"zone_name,omitempty"`
			ZoneOsPid          int    `json:"zone_os_pid,omitempty"`
		} `json:"server,omitempty"`
		TellsOff   int `json:"tells_off,omitempty"`
		WorldAdmin int `json:"world_admin,omitempty"`
		Zone       int `json:"zone,omitempty"`
	} `json:"data"`
	ExecutionTime string `json:"execution_time"`
	Method        string `json:"method"`
}

func (c *Client) GetWorldClientList() (WorldClientList, error) {
	o, err := c.GetTelnetClient().Command(telnet.CommandConfig{Command: "api get_client_list", EnforceJson: true})
	if err != nil {
		return WorldClientList{}, err
	}
	var clientList WorldClientList
	err = json.Unmarshal([]byte(o), &clientList)
	if err != nil {
		return WorldClientList{}, err
	}

	return clientList, nil
}

type ReloadResponse struct {
	Data struct {
		Message string `json:"message"`
	} `json:"data"`
	ExecutionTime string `json:"execution_time"`
	Method        string `json:"method"`
}

func (c *Client) Reload(reloadType string) (ReloadResponse, error) {
	o, err := c.GetTelnetClient().Command(
		telnet.CommandConfig{Command: fmt.Sprintf("api reload %v", reloadType), EnforceJson: true},
	)
	if err != nil {
		return ReloadResponse{}, err
	}
	var r ReloadResponse
	err = json.Unmarshal([]byte(o), &r)
	if err != nil {
		return ReloadResponse{}, err
	}

	return r, nil
}

type ReloadTypesResponse struct {
	Data []struct {
		Command     string `json:"command"`
		Description string `json:"description"`
		Opcode      int    `json:"opcode"`
	} `json:"data"`
	ExecutionTime string `json:"execution_time"`
	Method        string `json:"method"`
}

func (c *Client) GetReloadTypes() (ReloadTypesResponse, error) {
	o, err := c.GetTelnetClient().Command(
		telnet.CommandConfig{Command: "api get_reload_types", EnforceJson: true},
	)
	if err != nil {
		return ReloadTypesResponse{}, err
	}
	var r ReloadTypesResponse
	err = json.Unmarshal([]byte(o), &r)
	if err != nil {
		return ReloadTypesResponse{}, err
	}

	return r, nil
}

func (c *Client) GetWorldUptime() (string, error) {
	o, err := c.GetTelnetClient().Command(
		telnet.CommandConfig{Command: "uptime 0"},
	)
	if err != nil {
		return "", err
	}

	return o, nil
}

func (c *Client) MessageWorld(message string) error {
	_, _ = c.GetTelnetClient().Command(
		telnet.CommandConfig{Command: "emote world 15 " + message},
	)
	message = strings.ReplaceAll(message, "[SERVER MESSAGE] ", "")
	_, _ = c.GetTelnetClient().Command(
		telnet.CommandConfig{Command: "wwmarquee 15 " + message},
	)

	return nil
}

func (c *Client) ReloadQuestsForZone(zone string) error {
	_, err := c.GetTelnetClient().Command(
		telnet.CommandConfig{Command: "reloadzonequests " + zone},
	)
	if err != nil {
		return err
	}

	return nil
}
