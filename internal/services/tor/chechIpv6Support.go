package tor

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strings"
)

// CheckTorIPv6Support checks if Tor supports IPv6 by querying the control port.
func CheckTorIPv6Support() (bool, error) {
	// Connect to the Tor control port
	conn, err := net.Dial("tcp", "127.0.0.1:9051")
	if err != nil {
		return false, fmt.Errorf("failed to connect to Tor control port: %v", err)
	}
	defer conn.Close()

	// Read control.authcookie
	authCookiePath := "/var/lib/tor/control.authcookie"
	cookie, err := os.ReadFile(authCookiePath)
	if err != nil {
		return false, fmt.Errorf("failed to read control.authcookie: %v", err)
	}

	// Authenticate using the cookie
	authCommand := fmt.Sprintf("AUTHENTICATE %s\r\n", hex.EncodeToString(cookie))
	_, err = conn.Write([]byte(authCommand))
	if err != nil {
		return false, fmt.Errorf("failed to send authenticate command: %v", err)
	}

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return false, fmt.Errorf("failed to read AUTHENTICATE response: %v", err)
		}
		line = strings.TrimSpace(line)
		if line == "250 OK" {
			break // Authentication successful
		} else if strings.HasPrefix(line, "5") {
			return false, fmt.Errorf("authentication failed: %s", line)
		}
	}

	// Query IPv6 support
	_, err = conn.Write([]byte("GETINFO ip-to-country/ipv6-available\r\n"))
	if err != nil {
		return false, fmt.Errorf("failed to send GETINFO command: %v", err)
	}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return false, fmt.Errorf("failed to read GETINFO response: %v", err)
		}
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "250-ip-to-country/ipv6-available=") {
			value := strings.TrimPrefix(line, "250-ip-to-country/ipv6-available=")
			return value == "1", nil // Return true if "1", false otherwise
		} else if strings.HasPrefix(line, "250 OK") {
			break // End of response
		}
	}

	return false, fmt.Errorf("unexpected response: failed to determine IPv6 support")
}
