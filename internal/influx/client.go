package influx

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"net/http"
	"time"
)

type Client struct {
	client influxdb2.Client // influx client
	writer api.WriteAPI     // write
	alive  bool             // determines whether or not influx is alive to record to or not since its optional
}

// Alive Used to determine if influx is alive or not
// This is so we can check to see if the endpoint is alive before trying to send metrics
func (c *Client) Alive() bool {
	return c.alive
}

func (c *Client) Writer() api.WriteAPI {
	return c.writer
}

func (c *Client) Client() influxdb2.Client {
	return c.client
}

func NewClient() *Client {
	c := &Client{}
	c.alive = c.Ping()
	if c.Alive() {
		c.Init()
	}

	return c
}

// Init initialize a fresh client
func (c *Client) Init() {
	c.client = influxdb2.NewClient("http://influxdb:8086", "")
	c.writer = c.client.WriteAPI("", "db0")
}

// Ping Used during client boot to determine if influx is alive to send metrics to or not
func (c *Client) Ping() bool {
	h := http.Client{
		Timeout: 300 * time.Millisecond,
	}

	_, err := h.Get("http://influxdb:8086/health")
	if err != nil {
		return false
	}

	return true
}
