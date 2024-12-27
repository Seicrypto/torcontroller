package initializer

import (
	"fmt"

	"github.com/Seicrypto/torcontroller/internal/controller"
)

func verifyTorrcConfig() bool {
	runner := &controller.RealCommandRunner{}
	cmd := []string{"sudo", "tor", "--verify-config"}
	if _, err := runner.Run(cmd[0], cmd[1:]...); err != nil {
		fmt.Printf("- Torrc config validation failed: %v\n", err)
		return false
	}
	return true
}
