package devicedetector

import (
	"fmt"
	"testing"
)

// TestDeviceDetector tests initialization and parsing errors dependent js lib.
func TestDeviceDetector(t *testing.T) {
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

// BenchmarkDeviceDetector benchmarks the device detector.
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

// ExampleDeviceDetector shows how to use the device detector.
func ExampleNewDeviceDetector() {
	userAgent := "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36"
	detector, err := NewDeviceDetector(DeviceDetectorOptions{})
	if err != nil {
		fmt.Printf("failed to create device detector: %v", err)
	}
	result, err := detector.Parse(userAgent)
	if err != nil {
		fmt.Printf("failed to parse user agent: %v", err)
	}
	if result == nil {
		fmt.Printf("result is nil")
	}
	if result == "" {
		fmt.Printf("result is empty")
	}
	fmt.Printf("%s\n", result)
	// Output:
	// {"client":{"type":"browser","name":"Chrome Mobile","version":"41.0","engine":"Blink","engineVersion":""},"os":{"name":"Android","version":"6.0","platform":""},"device":{"type":"smartphone","brand":"Google","model":"Nexus 5X"},"bot":null}
}
