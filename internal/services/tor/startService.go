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

func StartTorService() error {
	logger := logger.GetLogger()

	var statusOut, statusErr bytes.Buffer
	statusCmd := exec.Command("sudo", "systemctl", "status", "tor")
	statusCmd.Stdout = &statusOut
	statusCmd.Stderr = &statusErr
	statusCmd.Env = append(os.Environ(), "LANG=C", "LC_ALL=C")

	if err := statusCmd.Run(); err != nil {
		output := statusOut.String()
		//  will return 3 as an exit code if tor is stopped
		if strings.Contains(output, "inactive (dead)") {
			logger.Warn("tor service is inactive. attempting to start...")
		} else if strings.Contains(output, "could not be found") {
			logger.Warn("tor service not found. please install and configure tor")
			return errors.New("tor service not found. please install and configure tor")
		} else {
			logger.Error(fmt.Sprintf("failed to check tor service status: %v", err))
			logger.Error(fmt.Sprintf("stderr: %s", statusErr.String()))
			return fmt.Errorf("failed to check tor service status: %w", err)
		}
	}

	output := statusOut.String()
	switch {
	case strings.Contains(output, "active (running)"):
		logger.Info("tor service is already running")
		return nil
	case strings.Contains(output, "inactive (dead)"):
		logger.Warn("tor service is inactive. attempting to start...")

		var startOut, startErr bytes.Buffer
		startCmd := exec.Command("sudo", "systemctl", "start", "tor")
		startCmd.Stdout = &startOut
		startCmd.Stderr = &startErr
		startCmd.Env = append(os.Environ(), "LANG=C", "LC_ALL=C")

		if err := startCmd.Run(); err != nil {
			logger.Error(fmt.Sprintf("failed to start tor service: %v", err))
			return fmt.Errorf("failed to start tor service: %w", err)
		}

		if err := statusCmd.Run(); err != nil {
			logger.Error(fmt.Sprintf("failed to recheck tor service status: %v", err))
			return fmt.Errorf("failed to recheck tor service status: %w", err)
		}

		if strings.Contains(statusOut.String(), "active (running)") {
			logger.Info("tor service started successfully")
			return nil
		}
		logger.Error("failed to start tor service, still not running")
		return errors.New("failed to start tor service, still not running")
	default:
		logger.Error("unexpected tor service status or service not configured properly")
		return errors.New("unexpected tor service status or service not configured properly")
	}
}
