package controller

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// StopServiceFactory is a factory method for stopping and checking systemd services
func (h *CommandHandler) StopServiceFactory(service string) error {
	var stopOut, stopErr bytes.Buffer
	stopCmd := exec.Command("sudo", "systemctl", "stop", service)
	stopCmd.Stdout = &stopOut
	stopCmd.Stderr = &stopErr

	h.Logger.Printf("[INFO] Stopping %s service...", service)
	if err := stopCmd.Run(); err != nil {
		output := stopOut.String()
		if strings.Contains(output, "could not be found") {
			h.Logger.Printf("[WARN] %s service not found. Please install and configure %s.", service, service)
			return fmt.Errorf("%s service not found", service)
		}
		h.Logger.Printf("[ERROR] Failed to stop %s service: %v", service, err)
		return fmt.Errorf("failed to stop %s service: %w", service, err)
	}

	// Verify that the service is stopped
	var statusOut, statusErr bytes.Buffer
	statusCmd := exec.Command("sudo", "systemctl", "status", service, "--no-pager")
	statusCmd.Stdout = &statusOut
	statusCmd.Stderr = &statusErr

	h.Logger.Printf("[INFO] Verifying %s service status...", service)
	if err := statusCmd.Run(); err != nil {
		output := statusOut.String()
		if strings.Contains(output, "inactive (dead)") {
			h.Logger.Printf("[INFO] %s service stopped successfully.", service)
			return nil
		}
		h.Logger.Printf("[ERROR] Failed to verify %s service status: %v", service, err)
		return fmt.Errorf("failed to verify %s service status: %w", service, err)
	}

	output := statusOut.String()
	if strings.Contains(output, "inactive (dead)") {
		h.Logger.Printf("[INFO] %s service stopped successfully.", service)
		return nil
	}

	h.Logger.Printf("[WARN] %s service stop command issued, but service is still running.", service)
	return fmt.Errorf("%s service stop command issued, but service is still running", service)
}

// StopTorService stops the Tor service using the factory method
func (h *CommandHandler) StopTorService() error {
	return h.StopServiceFactory("tor")
}

// StopPrivoxyService stops the Privoxy service using the factory method
func (h *CommandHandler) StopPrivoxyService() error {
	return h.StopServiceFactory("privoxy")
}
