package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a Torcontroller listener",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		execPath, err := os.Executable()
		if err != nil {
			fmt.Println("Error getting executable path:", err)
			return
		}

		command := exec.Command(execPath, "start-background")
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err = command.Start()
		if err != nil {
			fmt.Println("Error starting background process:", err)
			return
		}

		err = os.WriteFile(pidFile, []byte(fmt.Sprintf("%d", command.Process.Pid)), 0644)
		if err != nil {
			fmt.Printf("Error writing PID file: %v\n", err)
		}

		fmt.Printf("Torcontroller listener started with PID: %d\n", command.Process.Pid)
	},
}
