package ignition

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	req.Header.Set(AUTH_HEADER, c.Token)
	req.Header.Set("Accept", "application/json")
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
