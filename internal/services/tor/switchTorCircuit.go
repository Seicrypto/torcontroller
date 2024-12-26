package tor

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/Seicrypto/torcontroller/internal/services/logger"
)

// SwitchTorCircuit authenticates and switches Tor nodes using control.authcookie
func SwitchTorCircuit() error {
	logger := logger.GetLogger()

	conn, err := net.Dial("tcp", "127.0.0.1:9051")
	if err != nil {
		return fmt.Errorf("failed to connect to Tor control port: %v", err)
	}
	defer conn.Close()

	authCookiePath := "/var/lib/tor/control.authcookie"
	cookie, err := os.ReadFile(authCookiePath)
	if err != nil {
		return fmt.Errorf("failed to read control.authcookie: %v", err)
	}

	authCommand := fmt.Sprintf("AUTHENTICATE %s\r\n", hex.EncodeToString(cookie))
	_, err = conn.Write([]byte(authCommand))
	if err != nil {
		return fmt.Errorf("failed to send authenticate command: %v", err)
	}
	logger.Info("AUTHENTICATE command sent.")

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read AUTHENTICATE response: %v", err)
		}
		line = strings.TrimSpace(line)
		// logger.Info(fmt.Sprintf("AUTHENTICATE response: %s", line))
		if line == "250 OK" {
			break
		} else if strings.HasPrefix(line, "5") { // Error Code
			return fmt.Errorf("authentication failed: %s", line)
		}
	}

	_, err = conn.Write([]byte("SIGNAL NEWNYM\r\n"))
	if err != nil {
		return fmt.Errorf("failed to send SIGNAL NEWNYM command: %v", err)
	}
	logger.Info("SIGNAL NEWNYM command sent.")

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read SIGNAL NEWNYM response: %v", err)
		}
		line = strings.TrimSpace(line)
		// logger.Info(fmt.Sprintf("SIGNAL NEWNYM response: %s", line))
		if line == "250 OK" {
			logger.Info("Tor circuit switched successfully.")
			return nil
		} else if strings.HasPrefix(line, "5") { // Error Code
			return fmt.Errorf("SIGNAL NEWNYM failed: %s", line)
		}
	}
}
