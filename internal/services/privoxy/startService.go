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

// StartPrivoxyService starts the Privoxy systemd service
func StartPrivoxyService() error {
	logger := logger.GetLogger()

	var statusOut, statusErr bytes.Buffer
	statusCmd := exec.Command("sudo", "systemctl", "status", "privoxy", "--no-pager")
	statusCmd.Stdout = &statusOut
	statusCmd.Stderr = &statusErr
	statusCmd.Env = append(os.Environ(), "LANG=C", "LC_ALL=C")

	if err := statusCmd.Run(); err != nil {
		output := statusOut.String()
		//  will return 3 as an exit code if privoxy is stopped
		if strings.Contains(output, "inactive (dead)") {
			logger.Warn("Privoxy service is inactive. attempting to start...")
		} else if strings.Contains(output, "could not be found") {
			logger.Warn("Privoxy service not found. Please install and configure Privoxy.")
			return errors.New("privoxy service not found. please install and configure privoxy")
		} else if strings.Contains(output, "active (running)") {
			logger.Info("Privoxy service is already running.")
			return nil
		} else {
			logger.Error(fmt.Sprintf("failed to check privoxy service status: %v", err))
			logger.Error(fmt.Sprintf("stderr: %s", statusErr.String()))
			return fmt.Errorf("failed to check privoxy service status: %w", err)
		}
	}

	var startOut, startErr bytes.Buffer
	startCmd := exec.Command("sudo", "systemctl", "start", "privoxy")
	startCmd.Stdout = &startOut
	startCmd.Stderr = &startErr
	startCmd.Env = append(os.Environ(), "LANG=C", "LC_ALL=C")

	if err := startCmd.Run(); err != nil {
		logger.Error(fmt.Sprintf("Failed to start Privoxy service: %v", err))
		return fmt.Errorf("failed to start Privoxy service: %w", err)
	}

	logger.Info("Rechecking Privoxy service status...")
	statusOut.Reset() // Clear the last output
	statusErr.Reset()

	if err := statusCmd.Run(); err != nil {
		output := statusOut.String()
		if strings.Contains(output, "active (running)") {
			logger.Info("Privoxy service is now running.")
			return nil
		} else if strings.Contains(err.Error(), "already started") {
			logger.Info("Privoxy service already started.")
			return nil
		}
		logger.Error(fmt.Sprintf("Failed to recheck Privoxy service status: %v", err))
		logger.Error(fmt.Sprintf("stderr: %s", statusErr.String()))
		return fmt.Errorf("failed to recheck Privoxy service status: %w", err)
	}

	if strings.Contains(statusOut.String(), "active (running)") {
		logger.Info("Privoxy service started successfully.")
		return nil
	}
	logger.Error("Failed to start Privoxy service, still not running.")
	return errors.New("failed to start Privoxy service, still not running")
}
