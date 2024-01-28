package eqemuserver

import (
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/telnet"
)

type Client struct {
	telnet *telnet.Client
}

func NewClient(telnet *telnet.Client) *Client {
	return &Client{telnet: telnet}
}

type WorldZoneList struct {
	Data []struct {
		BootingUp          bool   `json:"booting_up"`
		ClientAddress      string `json:"client_address"`
		ClientLocalAddress string `json:"client_local_address"`
		ClientPort         int    `json:"client_port"`
		CompileTime        string `json:"compile_time"`
		ID                 int    `json:"id"`
		InstanceID         int    `json:"instance_id"`
		IP                 string `json:"ip"`
		IsStaticZone       bool   `json:"is_static_zone"`
		LaunchName         string `json:"launch_name"`
		LaunchedName       string `json:"launched_name"`
		NumberPlayers      int    `json:"number_players"`
		Port               int    `json:"port"`
		PreviousZoneID     int    `json:"previous_zone_id"`
		UUID               string `json:"uuid"`
		ZoneID             int    `json:"zone_id"`
		ZoneLongName       string `json:"zone_long_name"`
		ZoneName           string `json:"zone_name"`
		ZoneOsPid          int    `json:"zone_os_pid"`
	} `json:"data"`
	ExecutionTime string `json:"execution_time"`
	Method        string `json:"method"`
}

func (c *Client) GetZoneList() (WorldZoneList, error) {
	o, err := c.telnet.Command(
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

type LockStatusResponse struct {
	Data struct {
		Locked bool `json:"locked"`
	} `json:"data"`
	ExecutionTime string `json:"execution_time"`
	Method        string `json:"method"`
}

// GetLockStatus returns the lock status of the server
func (c *Client) GetLockStatus() (bool, error) {
	o, err := c.telnet.Command(
		telnet.CommandConfig{Command: "api lock_status", EnforceJson: true},
	)
	if err != nil {
		return false, err
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
	_, err := c.telnet.Command(
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
		AccountName          string `json:"account_name"`
		Admin                int    `json:"admin"`
		Anon                 int    `json:"anon"`
		CharacterID          int    `json:"character_id"`
		Class                int    `json:"class"`
		ClientVersion        int    `json:"client_version"`
		Gm                   int    `json:"gm"`
		GuildID              int    `json:"guild_id"`
		ID                   int    `json:"id"`
		Instance             int    `json:"instance"`
		IP                   int    `json:"ip"`
		IsLocalClient        bool   `json:"is_local_client"`
		Level                int    `json:"level"`
		Lfg                  bool   `json:"lfg"`
		LfgComments          string `json:"lfg_comments"`
		LfgFromLevel         int    `json:"lfg_from_level"`
		LfgMatchFilter       bool   `json:"lfg_match_filter"`
		LfgToLevel           int    `json:"lfg_to_level"`
		LoginserverAccountID int    `json:"loginserver_account_id"`
		LoginserverID        int    `json:"loginserver_id"`
		LoginserverName      string `json:"loginserver_name"`
		Name                 string `json:"name"`
		Online               int    `json:"online"`
		Race                 int    `json:"race"`
		Server               struct {
			ClientAddress      string `json:"client_address"`
			ClientLocalAddress string `json:"client_local_address"`
			ClientPort         int    `json:"client_port"`
			CompileTime        string `json:"compile_time"`
			ID                 int    `json:"id"`
			InstanceID         int    `json:"instance_id"`
			IP                 string `json:"ip"`
			IsBooting          bool   `json:"is_booting"`
			LaunchName         string `json:"launch_name"`
			LaunchedName       string `json:"launched_name"`
			NumberPlayers      int    `json:"number_players"`
			Port               int    `json:"port"`
			PreviousZoneID     int    `json:"previous_zone_id"`
			StaticZone         bool   `json:"static_zone"`
			Uui                string `json:"uui"`
			ZoneID             int    `json:"zone_id"`
			ZoneLongName       string `json:"zone_long_name"`
			ZoneName           string `json:"zone_name"`
			ZoneOsPid          int    `json:"zone_os_pid"`
		} `json:"server"`
		TellsOff   int `json:"tells_off"`
		WorldAdmin int `json:"world_admin"`
		Zone       int `json:"zone"`
	} `json:"data"`
	ExecutionTime string `json:"execution_time"`
	Method        string `json:"method"`
}

func (c *Client) GetWorldClientList() (WorldClientList, error) {
	o, err := c.telnet.Command(telnet.CommandConfig{Command: "api get_client_list", EnforceJson: true})
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
	o, err := c.telnet.Command(
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
	o, err := c.telnet.Command(
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
	o, err := c.telnet.Command(
		telnet.CommandConfig{Command: "uptime 0"},
	)
	if err != nil {
		return "", err
	}

	return o, nil
}
