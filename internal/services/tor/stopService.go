package tor

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Seicrypto/torcontroller/internal/services/logger"
)

// StopTorService stops the Tor service using systemctl
func StopTorService() error {
	logger := logger.GetLogger()

	var stopOut, stopErr bytes.Buffer
	stopCmd := exec.Command("sudo", "systemctl", "stop", "tor")
	stopCmd.Stdout = &stopOut
	stopCmd.Stderr = &stopErr
	stopCmd.Env = append(os.Environ(), "LANG=C", "LC_ALL=C")

	if err := stopCmd.Run(); err != nil {
		output := stopOut.String()
		if strings.Contains(output, "could not be found") {
			logger.Warn("tor service not found. please install and configure tor")
			return errors.New("tor service not found. please install and configure tor")
		}
		logger.Error(fmt.Sprintf("failed to stop tor service: %v", err))
		logger.Error(fmt.Sprintf("stderr: %s", stopErr.String()))
		return fmt.Errorf("failed to stop tor service: %w", err)
	}

	var statusOut, statusErr bytes.Buffer
	statusCmd := exec.Command("sudo", "systemctl", "status", "tor", "--no-pager")
	statusCmd.Stdout = &statusOut
	statusCmd.Stderr = &statusErr
	statusCmd.Env = append(os.Environ(), "LANG=C", "LC_ALL=C")

	if err := statusCmd.Run(); err != nil {
		output := statusOut.String()
		if strings.Contains(output, "inactive (dead)") {
			logger.Info("tor service stopped successfully")
			return nil
		}
		logger.Error(fmt.Sprintf("failed to recheck tor service status: %v", err))
		// logger.Error(fmt.Sprintf("stderr: %s", statusErr.String()))
		return fmt.Errorf("failed to recheck tor service status: %w", err)
	}

	output := statusOut.String()
	if strings.Contains(output, "inactive (dead)") {
		logger.Info("tor service stopped successfully")
		return nil
	}

	logger.Error("tor service is still running after attempting to stop it")
	return errors.New("tor service is still running after attempting to stop it")
}
