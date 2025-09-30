package ignition

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type CurrentPerformanceData struct {
	CPU        float64 `json:"cpu"`
	HeapMemory float64 `json:"heapMemory"`
	MaxMemory  int64   `json:"maxMemory"`
}

func (c *Client) CurrentPerformanceData() (*CurrentPerformanceData, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/data/api/v1/systemPerformance/currentGauges", c.GetGatewayAddress()), nil)
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
	var data CurrentPerformanceData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type HistoricalPerformanceData struct {
	CPUChartDatapoints []struct {
		HistID    int     `json:"histId"`
		Timestamp int     `json:"timestamp"`
		Value     float32 `json:"value"`
	} `json:"cpuChartDatapoints"`
	MemoryChartDatapoints struct {
		HeapMemoryDatapoints []struct {
			HistID    int     `json:"histId"`
			Timestamp int     `json:"timestamp"`
			Value     float32 `json:"value"`
		} `json:"heapMemoryDatapoints"`
		NonHeapMemoryDatapoints []struct {
			HistID    int     `json:"histId"`
			Timestamp int     `json:"timestamp"`
			Value     float32 `json:"value"`
		} `json:"nonHeapMemoryDatapoints"`
	} `json:"memoryChartDatapoints"`
}

func (c *Client) HistoricalPerformanceData() (*HistoricalPerformanceData, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/data/api/v1/systemPerformance/charts", c.GetGatewayAddress()), nil)
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
	var data HistoricalPerformanceData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type ClockDrift struct {
	ClockDrift float32 `json:"clockDrift"`
}

func (c *Client) CurrentClockDrift() (*ClockDrift, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/data/api/v1/systemPerformance/driftGauge", c.GetGatewayAddress()), nil)
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
	var data ClockDrift
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type ThreadExecutionData struct {
	Running      int `json:"running"`
	Waiting      int `json:"waiting"`
	TimedWaiting int `json:"timedWaiting"`
	Blocked      int `json:"blocked"`
}

func (c *Client) ThreadExecutionData() (*ThreadExecutionData, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/data/api/v1/systemPerformance/threads", c.GetGatewayAddress()), nil)
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
	var data ThreadExecutionData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type ClockDriftEvents struct {
	ClockDriftEvents []struct {
		Description string `json:"description"`
		Status      string `json:"status"`
		Timestamp   int    `json:"timestamp"`
	} `json:"clockDriftEvents"`
}

func (c *Client) ClockDriftEvents() (*ClockDriftEvents, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/data/api/v1/systemPerformance/driftEvents", c.GetGatewayAddress()), nil)
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
	var data ClockDriftEvents
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) RestartGateway() error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/data/api/v1/restart-tasks/restart", c.GetGatewayAddress()), nil)
	if err != nil {
		return err
	}
	setHeaders(req, c.Token)
	q := req.URL.Query()
	q.Add("confirm", "true")
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d, status: %v", resp.StatusCode, resp.Status)
	}
	return nil
}

type GatewayInfo struct {
	Name                 string `json:"name"`
	RedundancyRole       string `json:"redundancyRole"`
	Edition              string `json:"edition"`
	Hostname             string `json:"hostname"`
	Port                 string `json:"port"`
	IgnitionVersion      string `json:"ignitionVersion"`
	DeploymentMode       string `json:"deploymentMode"`
	TimeZone             string `json:"timeZone"`
	TimeZoneID           string `json:"timeZoneId"`
	JvmVersion           string `json:"jvmVersion"`
	AllowUnsignedModules bool   `json:"allowUnsignedModules"`
	License              struct {
		Mode                string `json:"mode"`
		ValidForVersion     int    `json:"validForVersion"`
		ExpirationDate      string `json:"expirationDate"`
		LicenseRestrictions []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"licenseRestrictions"`
	} `json:"license"`
}

func (c *Client) GatewayInfo() (*GatewayInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/data/api/v1/gateway-info", c.GetGatewayAddress()), nil)
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
	var data GatewayInfo
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) DownloadSystemLogs() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/data/api/v1/logs/download", c.GetGatewayAddress()), nil)
	if err != nil {
		return err
	}
	setHeaders(req, c.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d, status: %v", resp.StatusCode, resp.Status)
	}
	dateStr := time.Now().Format("20060102_150405")
	outFile, err := os.Create(fmt.Sprintf("syslog_%v.sqlite", dateStr))
	if err != nil {
		return err
	}
	defer outFile.Close()
	_, err = outFile.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

type AuditLog struct {
	Items []struct {
	} `json:"items"`
	Metadata struct {
		Total    int `json:"total"`
		Matching int `json:"matching"`
		Limit    int `json:"limit"`
		Offset   int `json:"offset"`
	} `json:"metadata"`
}

func (c *Client) GetAuditLog(name string, params *map[string]string) (*AuditLog, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/data/api/v1/audit/log/%s", c.GetGatewayAddress(), name), nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req, c.Token)
	if params != nil {
		addQueryParams(req, params)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, status: %v", resp.StatusCode, resp.Status)
	}
	var data AuditLog
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type AuditProfileNames struct {
	Items []struct {
		Name string `json:"name"`
	} `json:"items"`
	Metadata struct {
		Total    int `json:"total"`
		Matching int `json:"matching"`
		Limit    int `json:"limit"`
		Offset   int `json:"offset"`
	} `json:"metadata"`
}

func (c *Client) GetAuditProfileNames(serverId string, params *map[string]string) (*AuditProfileNames, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/data/api/v1/audit/remote-profiles/%s", c.GetGatewayAddress(), serverId), nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req, c.Token)
	if params != nil {
		addQueryParams(req, params)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, status: %v", resp.StatusCode, resp.Status)
	}
	var data AuditProfileNames
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
