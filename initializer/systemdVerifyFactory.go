package initializer

import (
	"bytes"
	"fmt"
	"os/exec"
)

// verifyTorService checks the validity of the Tor service unit file.
func verifyTorService() bool {
	return VerifyService("/etc/systemd/system/tor.service")
}

// verifyPrivoxyService checks the validity of the Privoxy service unit file.
func verifyPrivoxyService() bool {
	return VerifyService("/etc/systemd/system/privoxy.service")
}

// verifyService is a helper function to validate a given service unit file.
func VerifyService(servicePath string) bool {
	var out, errBuf bytes.Buffer

	// Execute `systemd-analyze verify` command.
	cmd := exec.Command("sudo", "systemd-analyze", "verify", servicePath)
	cmd.Stdout = &out
	cmd.Stderr = &errBuf

	err := cmd.Run()

	// Log the result and handle errors.
	if err != nil {
		fmt.Printf("Verification failed for %s\nError: %v\nDetails: %s\n", servicePath, err, errBuf.String())
		return false
	}

	if out.Len() > 0 {
		fmt.Printf("Verification output for %s: %s\n", servicePath, out.String())
	}

	// fmt.Printf("Service %s is valid.\n", servicePath)
	return true
}
