package privoxy

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Seicrypto/torcontroller/internal/singleton/logger"
)

// StopPrivoxyService stops the Privoxy systemd service
func StopPrivoxyService() error {
	logger := logger.GetLogger()

	var stopOut, stopErr bytes.Buffer
	stopCmd := exec.Command("sudo", "systemctl", "stop", "privoxy")
	stopCmd.Stdout = &stopOut
	stopCmd.Stderr = &stopErr
	stopCmd.Env = append(os.Environ(), "LANG=C", "LC_ALL=C")

	if err := stopCmd.Run(); err != nil {
		logger.Error(fmt.Sprintf("Failed to stop Privoxy service: %v", err))
		return fmt.Errorf("failed to stop Privoxy service: %w", err)
	}

	// Verify that the service is stopped
	var statusOut, statusErr bytes.Buffer
	statusCmd := exec.Command("sudo", "systemctl", "status", "privoxy")
	statusCmd.Stdout = &statusOut
	statusCmd.Stderr = &statusErr
	statusCmd.Env = append(os.Environ(), "LANG=C", "LC_ALL=C")

	if err := statusCmd.Run(); err != nil {
		output := statusOut.String()
		if strings.Contains(output, "inactive (dead)") {
			logger.Info("Privoxy service stopped successfully.")
			return nil
		}

		logger.Error(fmt.Sprintf("Failed to verify Privoxy service status: %v", err))
		return fmt.Errorf("failed to verify Privoxy service status: %w", err)
	}

	output := statusOut.String()
	if strings.Contains(output, "inactive (dead)") {
		logger.Info("privoxy service stopped successfully")
		return nil
	}

	logger.Warn("Privoxy service stop command issued, but service is still running.")
	return errors.New("privoxy service stop command issued, but service is still running")
}
