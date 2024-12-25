package initializer

import (
	"bytes"
	"fmt"
	"os/exec"
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
		fmt.Printf("[ERROR] Failed to show service %s configuration: %v", serviceName, err)
		return false
	}
	fmt.Printf("[INFO] Service %s configuration:\n%s", serviceName, out.String())
	return true
}
