package devicedetector

import (
	"fmt"
	"testing"
)

func TestNewDeviceDetector(t *testing.T) {
	userAgent := "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36"
	detector, err := NewDeviceDetector(DeviceDetectorOptions{})
	if err != nil {
		t.Errorf("failed to create device detector: %v", err)
	}
	result, err := detector.Parse(userAgent)
	if err != nil {
		t.Errorf("failed to parse user agent: %v", err)
	}
	if result == nil {
		t.Errorf("result is nil")
	}
	if result == "" {
		t.Errorf("result is empty")
	}

	fmt.Printf("%s\n", result)
}

func BenchmarkDeviceDetector(b *testing.B) {
	userAgent := "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36"
	detector, err := NewDeviceDetector(DeviceDetectorOptions{})
	if err != nil {
		b.Errorf("failed to create device detector: %v", err)
	}
	for i := 0; i < b.N; i++ {
		_, err := detector.Parse(userAgent)
		if err != nil {
			b.Errorf("failed to parse user agent: %v", err)
		}
	}
}
