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

// RpcClientRegisterRequest is the request payload for client registration.
type RpcClientRegisterRequest struct {
	ClientAddress string `json:"client_address"`
	Hostname      string `json:"hostname"`
}

// RpcZoneCountResponse is the response payload for zone count requests.
type RpcZoneCountResponse struct {
	ZoneCount    int `json:"zone_count"`
	MaxZoneCount int `json:"max_zone_count"`
}

// RpcLaunchZonesRequest is the request payload for launching zones.
type RpcLaunchZonesRequest struct {
	ZoneCount int `json:"zone_count"`
}

const (
	rpcServerActionShutdown = "shutdown"
	rpcServerActionStart    = "start"
	rpcServerActionRestart  = "restart"
)

// RpcServerActionRequest is the request payload for server actions.
type RpcServerActionRequest struct {
	Action string `json:"action"`
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

	// this gets the IP address of the root world server, this is our root launcher
	baseUrl := fmt.Sprintf(
		"http://%v:%v",
		cfg.Server.World.TCP.IP,
		3005,
	)

	url := fmt.Sprintf("%s/api/v1/dzs/register", baseUrl)
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

// rpcClientRootNodeServerAction sends a request to the root node
func (l *Launcher) rpcClientRootNodeServerAction(action string) error {
	return l.makeRequest(
		http.MethodPost,
		"http://127.0.0.1:3005/api/v1/dzs/root-node-action",
		RpcServerActionRequest{Action: action},
		nil,
	)
}

func (l *Launcher) rpcClientServerStart(node LauncherDistributedNode) error {
	return nil
}

func (l *Launcher) rpcClientServerStop(node LauncherDistributedNode) error {
	url := fmt.Sprintf("http://%v:3005/api/v1/dzs/server-stop", node.Address)
	return l.makeRequest(http.MethodPost, url, nil, nil)
}

func (l *Launcher) rpcClientServerRestart(node LauncherDistributedNode) error {
	return nil
}
