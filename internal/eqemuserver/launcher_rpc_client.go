package eqemuserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// getBaseRpcUrl retrieves the base RPC URL from the configuration.
func (l *Launcher) getBaseRpcUrl() string {
	cfg, _ := l.serverconfig.Get()
	var baseRpcUrl string
	if cfg.WebAdmin != nil && cfg.WebAdmin.Launcher != nil && cfg.WebAdmin.Launcher.LeafNodeConfig != nil {
		if cfg.WebAdmin.Launcher.LeafNodeConfig.RootSpireUrl != "" {
			baseRpcUrl = cfg.WebAdmin.Launcher.LeafNodeConfig.RootSpireUrl
		}
		if cfg.WebAdmin.Launcher.LeafNodeConfig.RootSpirePort != "" {
			baseRpcUrl = fmt.Sprintf("%s:%s", baseRpcUrl, cfg.WebAdmin.Launcher.LeafNodeConfig.RootSpirePort)
		}
	}

	return baseRpcUrl
}

type RpcClientRegisterRequest struct {
	ClientAddress string `json:"client_address"`
	Hostname      string `json:"hostname"`
}

type RpcZoneCountResponse struct {
	ZoneCount    int `json:"zone_count"`
	MaxZoneCount int `json:"max_zone_count"`
}

type RpcLaunchZonesRequest struct {
	ZoneCount int `json:"zone_count"`
}

// makeRequest handles generic HTTP requests and response processing.
func (l *Launcher) makeRequest(method, url string, payload any, result any) error {
	cfg, _ := l.serverconfig.Get()

	// Marshal payload if provided
	var body io.Reader
	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(jsonData)
	}

	// Create the HTTP request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	// Add headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("RPC_KEY", cfg.Server.World.Key)

	// Send the request
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read and handle the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with code: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	// If a result object is provided, unmarshal the response into it
	if result != nil {
		if err := json.Unmarshal(bodyBytes, result); err != nil {
			return err
		}
	}

	return nil
}

// rpcClientRegister registers the client.
func (l *Launcher) rpcClientRegister() error {
	hostname, _ := os.Hostname()
	cfg, _ := l.serverconfig.Get()

	url := fmt.Sprintf("%s/api/v1/dzs/register", l.getBaseRpcUrl())
	payload := RpcClientRegisterRequest{
		ClientAddress: cfg.Server.World.Address,
		Hostname:      hostname,
	}

	return l.makeRequest(http.MethodPost, url, payload, nil)
}

// rpcClientGetZoneCount retrieves the zone count from a node.
func (l *Launcher) rpcClientGetZoneCount(node LauncherDistributedNode) (RpcZoneCountResponse, error) {
	url := fmt.Sprintf("http://%v:3005/api/v1/dzs/zone-count", node.Address)

	var response RpcZoneCountResponse
	err := l.makeRequest(http.MethodGet, url, nil, &response)
	return response, err
}

// rpcClientSetTargetZoneCount sets the target zone count on a node.
func (l *Launcher) rpcClientSetTargetZoneCount(node LauncherDistributedNode, zoneCount int) error {
	url := fmt.Sprintf("http://%v:3005/api/v1/dzs/set-zone-count", node.Address)
	payload := RpcLaunchZonesRequest{ZoneCount: zoneCount}

	return l.makeRequest(http.MethodPost, url, payload, nil)
}
