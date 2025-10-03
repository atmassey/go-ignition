package ignition

import (
	"os"
	"testing"
)

func TestRestart(t *testing.T) {
	err := os.Setenv("API_TOKEN", "dev:5-nl5YiEKvKRCQs0sSce5zKKS_0HDwxL_WElZ0jOfWM")
	if err != nil {
		t.Fatalf("Failed to set env var: %v", err)
	}
	client, err := NewClient("localhost", int(8088), false)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	performance, err := client.CurrentPerformanceData()
	if err != nil {
		t.Fatalf("Failed to get current performance data: %v", err)
	}
	t.Logf("Current Performance Data: CPU: %f, HeapMemory: %f, MaxMemory: %d", performance.CPU, performance.HeapMemory, performance.MaxMemory)
	// err = client.RestartGateway()
	// if err != nil {
	// 	t.Fatalf("Failed to restart: %v", err)
	// }
	gatewayInfo, err := client.GatewayInfo()
	if err != nil {
		t.Fatalf("Failed to get gateway info: %v", err)
	}
	t.Logf("Gateway Info: Name: %s, Hostname: %s, Port: %s, ExpirationDate: %s", gatewayInfo.Name, gatewayInfo.Hostname, gatewayInfo.Port, gatewayInfo.License.ExpirationDate)
}
