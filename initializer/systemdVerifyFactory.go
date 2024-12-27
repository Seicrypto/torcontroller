package initializer

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// verifyTorService checks the validity of the Tor service unit file.
func checkTorService() bool {
	return CheckServiceFile("tor")
}

// verifyPrivoxyService checks the validity of the Privoxy service unit file.
func checkPrivoxyService() bool {
	return CheckServiceFile("privoxy")
}

// verifyService is a helper function to validate a given service unit file.
func CheckServiceFile(serviceName string) bool {
	cmd := exec.Command("sudo", "systemctl", "show", serviceName)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		fmt.Printf("[ERROR] Failed to validate service %s: %v\n", serviceName, err)
		return false
	}

	// Parse the output to check for critical fields
	output := out.String()
	// if !strings.Contains(output, "ActiveState=active") {
	// 	fmt.Printf("[ERROR] Service %s is not active.\n", serviceName)
	// 	return false
	// }

	if !strings.Contains(output, "LoadState=loaded") {
		fmt.Printf("[ERROR] Service %s is not loaded properly.\n", serviceName)
		return false
	}

	// If necessary, add more checks for specific fields in the service configuration
	return true
}
