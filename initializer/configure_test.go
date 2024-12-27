package initializer_test

import (
	"testing"

	"github.com/Seicrypto/torcontroller/initializer"
)

// TestPlaceTorServiceFile tests the PlaceTorServiceFile method.
func TestPlaceTorServiceFile(t *testing.T) {
	mockTemplates := &MockTemplates{
		Files: map[string][]byte{
			"templates/tor.service": []byte("[Service]\nExecStart=/usr/bin/tor\n"),
		},
	}
	mockRunner := &MockCommandRunner{}
	init := initializer.NewInitializer(mockTemplates, mockRunner, &MockFileSystem{})

	err := init.PlaceTorServiceFile()
	if err != nil {
		t.Errorf("PlaceTorServiceFile failed: %v", err)
	}
}

// TestPlacePrivoxyServiceFile tests the PlacePrivoxyServiceFile method.
func TestPlacePrivoxyServiceFile(t *testing.T) {
	mockTemplates := &MockTemplates{
		Files: map[string][]byte{
			"templates/privoxy.service": []byte("[Service]\nExecStart=/usr/bin/privoxy\n"),
		},
	}
	mockRunner := &MockCommandRunner{}
	init := initializer.NewInitializer(mockTemplates, mockRunner, &MockFileSystem{})

	err := init.PlacePrivoxyServiceFile()
	if err != nil {
		t.Errorf("PlacePrivoxyServiceFile failed: %v", err)
	}
}

// TestPlaceSudoersFile tests the PlaceSudoersFile method.
func TestPlaceSudoersFile(t *testing.T) {
	mockTemplates := &MockTemplates{
		Files: map[string][]byte{
			"templates/sudoers.d/torcontroller": []byte("torcontroller ALL=(ALL) NOPASSWD: /usr/bin/tor"),
		},
	}
	mockRunner := &MockCommandRunner{}
	init := initializer.NewInitializer(mockTemplates, mockRunner, &MockFileSystem{})

	err := init.PlaceSudoersFile()
	if err != nil {
		t.Errorf("PlaceSudoersFile failed: %v", err)
	}
}

// TestPlaceTorrcConfig tests the PlaceTorrcConfig method.
func TestPlaceTorrcConfig(t *testing.T) {
	mockTemplates := &MockTemplates{
		Files: map[string][]byte{
			"templates/tor/torrc": []byte("SOCKSPort 9050"),
		},
	}
	mockRunner := &MockCommandRunner{}
	init := initializer.NewInitializer(mockTemplates, mockRunner, &MockFileSystem{})

	err := init.PlaceTorrcConfig()
	if err != nil {
		t.Errorf("PlaceTorrcConfig failed: %v", err)
	}
}

// TestPlacePrivoxyConfig tests the PlacePrivoxyConfig method.
func TestPlacePrivoxyConfig(t *testing.T) {
	mockTemplates := &MockTemplates{
		Files: map[string][]byte{
			"templates/privoxy/config": []byte("listen-address 127.0.0.1:8118"),
		},
	}
	mockRunner := &MockCommandRunner{}
	init := initializer.NewInitializer(mockTemplates, mockRunner, &MockFileSystem{})

	err := init.PlacePrivoxyConfig()
	if err != nil {
		t.Errorf("PlacePrivoxyConfig failed: %v", err)
	}
}
