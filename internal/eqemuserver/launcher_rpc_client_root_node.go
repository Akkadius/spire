package eqemuserver

import "net/http"

// rpcClientRootNodeServerShutdown sends a request to the root node
func (l *Launcher) rpcClientRootNodeServerShutdown() error {
	return l.makeRequest(http.MethodPost, "http://127.0.0.1:3005/api/v1/dzs/root-node-shutdown", nil, nil)
}

// rpcClientRootNodeKillProcess sends a request to the root node to kill a process
func (l *Launcher) rpcClientRootNodeKillProcess(zone *ZoneServer) error {
	return l.makeRequest(http.MethodPost, "http://127.0.0.1:3005/api/v1/dzs/root-node-kill-process", zone, nil)
}
