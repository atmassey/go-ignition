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

type OPCDeviceConfig struct {
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Enabled     bool     `json:"enabled"`
	Version     int      `json:"version"`
	Collection  string   `json:"collection"`
	Collections []string `json:"collections"`
	Signature   string   `json:"signature"`
	Config      struct {
		Profile struct {
			Type                   string `json:"type"`
			BrowsePath             string `json:"browsePath"`
			RolePermissionMappings []struct {
				Role        string   `json:"role"`
				Permissions []string `json:"permissions"`
			} `json:"rolePermissionMappings"`
		} `json:"profile"`
		Settings struct {
			Connectivity struct {
				Hostname     string `json:"hostname"`
				LocalAddress string `json:"localAddress"`
				Timeout      int    `json:"timeout"`
				Path         string `json:"path"`
			} `json:"connectivity"`
			Advanced struct {
				DisableAutomaticBrowse bool `json:"disableAutomaticBrowse"`
				ShowStringArrays       bool `json:"showStringArrays"`
				StatusPollRate         int  `json:"statusPollRate"`
				ConcurrentRequests     int  `json:"concurrentRequests"`
			} `json:"advanced"`
		} `json:"settings"`
	} `json:"config"`
	BackupConfig struct {
		Profile struct {
			Type                   string `json:"type"`
			BrowsePath             string `json:"browsePath"`
			RolePermissionMappings []struct {
				Role        string   `json:"role"`
				Permissions []string `json:"permissions"`
			} `json:"rolePermissionMappings"`
		} `json:"profile"`
		Settings struct {
			Connectivity struct {
				Hostname     string `json:"hostname"`
				LocalAddress string `json:"localAddress"`
				Timeout      int    `json:"timeout"`
				Path         string `json:"path"`
			} `json:"connectivity"`
			Advanced struct {
				DisableAutomaticBrowse bool `json:"disableAutomaticBrowse"`
				ShowStringArrays       bool `json:"showStringArrays"`
				StatusPollRate         int  `json:"statusPollRate"`
				ConcurrentRequests     int  `json:"concurrentRequests"`
			} `json:"advanced"`
		} `json:"settings"`
	} `json:"backupConfig"`
	Data       []string `json:"data"`
	Attributes struct {
	} `json:"attributes"`
	Healthchecks struct {
		Status struct {
			Name   string `json:"name"`
			Result struct {
				Healthy bool   `json:"healthy"`
				Message string `json:"message"`
				Error   struct {
					Message    string   `json:"message"`
					Stacktrace []string `json:"stacktrace"`
				} `json:"error"`
				Time     string `json:"time"`
				Duration int    `json:"duration"`
			} `json:"result"`
		} `json:"status"`
	} `json:"healthchecks"`
}

func (c *Client) GetOPCDeviceConfig(deviceName string) (*OPCDeviceConfig, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/data/api/v1/resources/config/com.inductiveautomation.opcua/device/%s", c.GetGatewayAddress(), deviceName), nil)
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
	var deviceConfig OPCDeviceConfig
	if err := json.NewDecoder(resp.Body).Decode(&deviceConfig); err != nil {
		return nil, err
	}
	return &deviceConfig, nil
}
