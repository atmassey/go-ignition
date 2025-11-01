package ignition

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OPCDeviceNames struct {
	Items []struct {
		Name    string `json:"name"`
		Enabled bool   `json:"enabled"`
	} `json:"items"`
	Metadata struct {
		Total    int `json:"total"`
		Matching int `json:"matching"`
		Limit    int `json:"limit"`
		Offset   int `json:"offset"`
	} `json:"metadata"`
}

func (c *Client) GetOPCDeviceNames() (*OPCDeviceNames, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/data/api/v1/resources/names/com.inductiveautomation.opcua/device", c.GetGatewayAddress()), nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req, c.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, status: %v", resp.StatusCode, resp.Status)
	}
	var deviceNames OPCDeviceNames
	if err := json.NewDecoder(resp.Body).Decode(&deviceNames); err != nil {
		return nil, err
	}
	return &deviceNames, nil
}
