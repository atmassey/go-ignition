package ignition

import (
	"fmt"
	"net/http"
	"os"
)

type Client struct {
	GatewayAddress string
	GatewayPort    int
	Token          string
	SSLEnabled     bool
}

const (
	AUTH_HEADER = "X-Ignition-API-Token"
)

func NewClient(GatewayAddress string, GatewayPort int, SSLFlag bool) (*Client, error) {
	client := &Client{
		GatewayAddress: GatewayAddress,
		GatewayPort:    GatewayPort,
		SSLEnabled:     SSLFlag,
	}

	token := os.Getenv("API_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("API_TOKEN environment variable is not set")
	}
	client.Token = token

	return client, nil
}

func (c *Client) GetGatewayAddress() string {
	var addr string
	if c.SSLEnabled {
		addr = fmt.Sprintf("https://%s:%d", c.GatewayAddress, c.GatewayPort)
	} else {
		addr = fmt.Sprintf("http://%s:%d", c.GatewayAddress, c.GatewayPort)
	}
	return addr
}

func setHeaders(req *http.Request, token string) {
	req.Header.Set(AUTH_HEADER, token)
	req.Header.Set("Accept", "application/json")
}
