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

func (c *Client) ClockDrift() (*ClockDrift, error) {
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
