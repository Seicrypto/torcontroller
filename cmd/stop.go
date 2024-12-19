package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Seicrypto/torcontroller/internal/services/logger"
	"github.com/spf13/cobra"
)

var StopCmd = &cobra.Command{
	Use:   "stop [socketPath]",
	Short: "Stop a Torcontroller listener",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile(pidFile)
		if err != nil {
			fmt.Printf("Error reading PID file: %v\n", err)
			return
		}

		pid, err := strconv.Atoi(string(data))
		if err != nil {
			fmt.Printf("Error parsing PID: %v\n", err)
			return
		}

		proc, err := os.FindProcess(pid)
		if err != nil {
			fmt.Printf("Error finding process: %v\n", err)
			return
		}

		err = proc.Kill()
		if err != nil {
			fmt.Printf("Error killing process: %v\n", err)
			return
		}

		os.Remove(socketPath)
		os.Remove(pidFile)
		logger := logger.GetLogger()
		logger.Printf("Torcontroller listener at %s stopped successfully.\n", socketPath)
		fmt.Printf("Torcontroller listener at %s stopped successfully.\n", socketPath)
	},
}
