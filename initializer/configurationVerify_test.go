package initializer_test

import (
	"os"
	"testing"

	"github.com/Seicrypto/torcontroller/initializer"
)

// TestVerifyConfigFile tests the functionality of VerifyConfigFile.
func TestVerifyConfigFile(t *testing.T) {
	// Mock configuration content
	mockConfigContent := []byte(`
rate_limit:
  min_read_rate: 20000
  min_write_rate: 10000
`)

	mockFileSystem := &MockFileSystem{
		Files: map[string]*MockFileInfo{
			"/etc/torcontroller/torcontroller.yml": {
				content: []byte(mockConfigContent),
				mode:    os.FileMode(0644),
				uid:     0,
				gid:     0,
			},
		},
	}

	mockTemplates := &MockTemplates{}
	mockRunner := &MockCommandRunner{}

	init := initializer.NewInitializer(mockTemplates, mockRunner, mockFileSystem)

	// Valid config file scenario
	valid := init.VerifyConfigFile("/etc/torcontroller/torcontroller.yml")
	if !valid {
		t.Errorf("expected configuration file to be valid")
	}

	// Invalid config file scenario (non-existent path)
	invalid := init.VerifyConfigFile("/invalid/path/torcontroller.yml")
	if invalid {
		t.Errorf("expected configuration file verification to fail for non-existent path")
	}
}
