package eqemuserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

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

func (l *Launcher) rpcClientRegister() error {
	cfg, _ := l.serverconfig.Get()
	hostname, _ := os.Hostname()

	body, _ := json.Marshal(RpcClientRegisterRequest{
		ClientAddress: cfg.Server.World.Address,
		Hostname:      hostname,
	})
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/dzs/register", l.getBaseRpcUrl()),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	// add json content type
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("RPC_KEY", cfg.Server.World.Key)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register code: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	contents := string(bodyBytes)
	if !strings.Contains(contents, "Registered") {
		return fmt.Errorf("failed to register, response: %s", contents)
	}

	return nil
}
