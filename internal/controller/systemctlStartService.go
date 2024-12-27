package controller

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// StartServiceFactory is a factory method for starting and checking systemd services
func (h *CommandHandler) StartServiceFactory(service string) error {
	var statusOut, statusErr bytes.Buffer
	statusCmd := exec.Command("sudo", "systemctl", "status", service, "--no-pager")
	statusCmd.Stdout = &statusOut
	statusCmd.Stderr = &statusErr

	h.Logger.Printf("[INFO] Checking %s service status...", service)
	if err := statusCmd.Run(); err != nil {
		output := statusOut.String()
		h.Logger.Printf("[INFO] %s service status output: %s", service, output)

		if strings.Contains(output, "inactive (dead)") {
			h.Logger.Printf("[WARN] %s service is inactive. Attempting to start...", service)
		} else if strings.Contains(output, "could not be found") {
			h.Logger.Printf("[ERROR] %s service not found. Please install and configure %s.", service, service)
			return fmt.Errorf("%s service not found", service)
		} else if strings.Contains(output, "active (running)") {
			h.Logger.Printf("%s service is already running.", service)
			return nil
		} else {
			h.Logger.Printf("[ERROR] Failed to check %s service status: %v", service, err)
			return fmt.Errorf("failed to check %s service status: %w", service, err)
		}
	}

	var startOut, startErr bytes.Buffer
	startCmd := exec.Command("sudo", "systemctl", "start", service)
	startCmd.Stdout = &startOut
	startCmd.Stderr = &startErr

	h.Logger.Printf("[INFO] Starting %s service...", service)
	if err := startCmd.Run(); err != nil {
		h.Logger.Printf("[ERROR] Failed to start %s service: %v", service, err)
		return fmt.Errorf("failed to start %s service: %w", service, err)
	}

	h.Logger.Printf("[INFO] Rechecking %s service status...", service)
	statusOut.Reset() // Clear previous output
	statusErr.Reset()

	if err := statusCmd.Run(); err != nil {
		output := statusOut.String()
		if strings.Contains(output, "active (running)") {
			h.Logger.Printf("[INFO] %s service is now running.", service)
			return nil
		} else if strings.Contains(err.Error(), "already started") {
			h.Logger.Printf("[INFO] %s service already started.", service)
			return nil
		}
		h.Logger.Printf("[ERROR] Failed to recheck %s service status: %v", service, err)
		return fmt.Errorf("failed to recheck %s service status: %w", service, err)
	}

	if strings.Contains(statusOut.String(), "active (running)") {
		h.Logger.Printf("[INFO] %s service started successfully.", service)
		return nil
	}

	h.Logger.Printf("[ERROR] %s service failed to start, still not running.", service)
	return fmt.Errorf("%s service failed to start, still not running", service)
}

// StartTorService starts the Tor service using the factory method
func (h *CommandHandler) StartTorService() error {
	return h.StartServiceFactory("tor")
}

// StartPrivoxyService starts the Privoxy service using the factory method
func (h *CommandHandler) StartPrivoxyService() error {
	return h.StartServiceFactory("privoxy")
}
