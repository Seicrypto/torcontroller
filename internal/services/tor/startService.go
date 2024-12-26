package tor

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Seicrypto/torcontroller/internal/singleton/logger"
)

func StartTorService() error {
	logger := logger.GetLogger()

	var statusOut, statusErr bytes.Buffer
	statusCmd := exec.Command("sudo", "systemctl", "status", "tor", "--no-pager")
	statusCmd.Stdout = &statusOut
	statusCmd.Stderr = &statusErr
	statusCmd.Env = append(os.Environ(), "LANG=C", "LC_ALL=C")

	if err := statusCmd.Run(); err != nil {
		output := statusOut.String()
		//  will return 3 as an exit code if tor is stopped
		logger.Info(fmt.Sprintf("Current stdout during checking: %s", output))
		logger.Info(fmt.Sprintf("Current stderr during checking: %s", statusErr.String()))
		logger.Info(fmt.Sprintf("Current cmd err during checking: %s", err))
		if strings.Contains(output, "inactive (dead)") {
			logger.Warn("tor service is inactive. attempting to start...")
		} else if strings.Contains(output, "could not be found") {
			logger.Warn("tor service not found. please install and configure tor")
			return errors.New("tor service not found. please install and configure tor")
		} else if strings.Contains(output, "active (running)") {
			logger.Info("Tor service is already running.")
			return nil
		} else {
			logger.Error(fmt.Sprintf("failed to check tor service status: %v", err))
			logger.Error(fmt.Sprintf("stderr: %s", statusErr.String()))
			return fmt.Errorf("failed to check tor service status: %w", err)
		}
	}

	var startOut, startErr bytes.Buffer
	startCmd := exec.Command("sudo", "systemctl", "start", "tor")
	startCmd.Stdout = &startOut
	startCmd.Stderr = &startErr
	startCmd.Env = append(os.Environ(), "LANG=C", "LC_ALL=C")

	if err := startCmd.Run(); err != nil {
		logger.Error(fmt.Sprintf("failed to start tor service: %v", err))
		return fmt.Errorf("failed to start tor service: %w", err)
	}

	logger.Info("Rechecking Tor service status...")
	statusOut.Reset() // Clear the last output
	statusErr.Reset()

	if err := statusCmd.Run(); err != nil {
		output := statusOut.String()
		if strings.Contains(output, "active (running)") {
			logger.Info("Tor service is now running.")
			return nil
		} else if strings.Contains(err.Error(), "already started") {
			logger.Info("Tor service already started.")
			return nil
		}
		logger.Error(fmt.Sprintf("Failed to recheck Tor service status: %v", err))
		logger.Error(fmt.Sprintf("stderr: %s", statusErr.String()))
		return fmt.Errorf("failed to recheck tor service status: %w", err)
	}

	if strings.Contains(statusOut.String(), "active (running)") {
		logger.Info("Tor service started successfully.")
		return nil
	}
	logger.Error("Tor service failed to start, still not running.")
	return errors.New("tor service failed to start, still not running")
}
