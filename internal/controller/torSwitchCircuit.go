package controller

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// SwitchTorCircuit authenticates and switches Tor nodes using control.authcookie.
func (h *CommandHandler) SwitchTorCircuit() error {
	// Establish a connection to the Tor control port
	conn, err := h.Socket.Dial()
	if err != nil {
		h.Logger.Printf("[ERROR] Failed to connect to Tor control port: %v", err)
		return fmt.Errorf("failed to connect to Tor control port: %v", err)
	}
	defer conn.Close()

	// read control.authcookie
	cookie, err := os.ReadFile("/var/lib/tor/control.authcookie")
	if err != nil {
		h.Logger.Printf("[ERROR] Failed to read control.authcookie: %v", err)
		return fmt.Errorf("failed to read control.authcookie: %v", err)
	}

	// Send the AUTHENTICATE command
	authCommand := fmt.Sprintf("AUTHENTICATE %s\r\n", hex.EncodeToString(cookie))
	_, err = conn.Write([]byte(authCommand))
	if err != nil {
		h.Logger.Printf("[ERROR] Failed to send authenticate command: %v", err)
		return fmt.Errorf("failed to send authenticate command: %v", err)
	}
	h.Logger.Println("[INFO] AUTHENTICATE command sent.")

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			h.Logger.Printf("[ERROR] Failed to read AUTHENTICATE response: %v", err)
			return fmt.Errorf("failed to read AUTHENTICATE response: %v", err)
		}
		line = strings.TrimSpace(line)
		h.Logger.Printf("[DEBUG] AUTHENTICATE response: %s", line)

		if line == "250 OK" {
			break
		} else if strings.HasPrefix(line, "5") { // Error Code
			h.Logger.Printf("[ERROR] Authentication failed: %s", line)
			return fmt.Errorf("authentication failed: %s", line)
		}
	}

	// Send the SIGNAL NEWNYM command.
	_, err = conn.Write([]byte("SIGNAL NEWNYM\r\n"))
	if err != nil {
		h.Logger.Printf("[ERROR] Failed to send SIGNAL NEWNYM command: %v", err)
		return fmt.Errorf("failed to send SIGNAL NEWNYM command: %v", err)
	}
	h.Logger.Println("[INFO] SIGNAL NEWNYM command sent.")

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			h.Logger.Printf("[ERROR] Failed to read SIGNAL NEWNYM response: %v", err)
			return fmt.Errorf("failed to read SIGNAL NEWNYM response: %v", err)
		}
		line = strings.TrimSpace(line)
		h.Logger.Printf("[DEBUG] SIGNAL NEWNYM response: %s", line)

		if line == "250 OK" {
			h.Logger.Println("[INFO] Tor circuit switched successfully.")
			return nil
		} else if strings.HasPrefix(line, "5") { // Error Code
			h.Logger.Printf("[ERROR] SIGNAL NEWNYM failed: %s", line)
			return fmt.Errorf("SIGNAL NEWNYM failed: %s", line)
		}
	}
}
