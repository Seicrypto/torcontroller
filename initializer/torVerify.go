package initializer

import (
	"fmt"

	runneradapter "github.com/Seicrypto/torcontroller/internal/services/runnerAdapter"
)

func verifyTorrcConfig() bool {
	runner := &runneradapter.RealCommandRunner{}
	cmd := []string{"sudo", "tor", "--verify-config"}
	if _, err := runner.Run(cmd[0], cmd[1:]...); err != nil {
		fmt.Printf("- Torrc config validation failed: %v\n", err)
		return false
	}
	return true
}
