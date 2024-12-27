package initializer

import (
	"fmt"
	"os"
	"syscall"

	"github.com/Seicrypto/torcontroller/internal/controller"
)

func sudoersFileVerify() bool {
	runner := &controller.RealCommandRunner{}
	sudoersPath := "/etc/sudoers.d/torcontroller"

	if _, err := os.Stat(sudoersPath); os.IsNotExist(err) {
		fmt.Println("- Sudoers configuration [MISSING]")
		return false
	}

	fileInfo, err := os.Stat(sudoersPath)
	if err != nil {
		fmt.Printf("Failed to check sudoers file: %v\n", err)
		return false
	}

	if fileInfo.Mode() != 0o440 {
		fmt.Println("- Sudoers configuration [INVALID PERMISSIONS]")
		return false
	}

	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok || stat.Uid != 0 || stat.Gid != 0 {
		fmt.Println("- Sudoers configuration [INVALID OWNER]")
		return false
	}

	cmd := []string{"sudo", "visudo", "-cf", sudoersPath}
	if _, err := runner.Run(cmd[0], cmd[1:]...); err != nil {
		fmt.Println("- Sudoers configuration [INVALID SYNTAX]")
		return false
	}

	return true
}
