package initializer_test

import (
	"errors"
	"testing"

	"github.com/Seicrypto/torcontroller/initializer"
)

func TestCheckTorService(t *testing.T) {
	mockRunner := &MockCommandRunner{
		Output: "LoadState=loaded",
	}

	initializer := initializer.NewInitializer(&MockTemplates{}, mockRunner, &MockFileSystem{})

	if !initializer.CheckTorService() {
		t.Errorf("expected CheckTorService to return true for valid service")
	}

	mockRunner.Output = "LoadState=failed"
	if initializer.CheckTorService() {
		t.Errorf("expected CheckTorService to return false for invalid service")
	}
}

func TestCheckPrivoxyService(t *testing.T) {
	mockRunner := &MockCommandRunner{
		Output: "LoadState=loaded",
	}

	initializer := initializer.NewInitializer(&MockTemplates{}, mockRunner, &MockFileSystem{})

	if !initializer.CheckPrivoxyService() {
		t.Errorf("expected CheckPrivoxyService to return true for valid service")
	}

	mockRunner.Output = "LoadState=inactive"
	if initializer.CheckPrivoxyService() {
		t.Errorf("expected CheckPrivoxyService to return false for invalid service")
	}
}

func TestCheckServiceFile(t *testing.T) {
	tests := []struct {
		serviceName   string
		mockOutput    string
		mockError     error
		expectedValid bool
	}{
		{
			serviceName:   "tor",
			mockOutput:    "LoadState=loaded",
			mockError:     nil,
			expectedValid: true,
		},
		{
			serviceName:   "privoxy",
			mockOutput:    "LoadState=failed",
			mockError:     nil,
			expectedValid: false,
		},
		{
			serviceName:   "tor",
			mockOutput:    "",
			mockError:     errors.New("command failed"),
			expectedValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.serviceName, func(t *testing.T) {
			mockRunner := &MockCommandRunner{
				Output: tt.mockOutput,
				Error:  tt.mockError,
			}

			init := initializer.NewInitializer(&MockTemplates{}, mockRunner, &MockFileSystem{})
			valid := init.CheckServiceFile(tt.serviceName)

			if valid != tt.expectedValid {
				t.Errorf("expected validity to be %v, got %v for service %s", tt.expectedValid, valid, tt.serviceName)
			}
		})
	}
}
