package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop [socketPath]",
	Short: "Stop a Torcontroller listener",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		// 读取 PID 文件
		data, err := os.ReadFile(pidFile)
		if err != nil {
			fmt.Printf("Error reading PID file: %v\n", err)
			return
		}

		// 转换 PID 并找到进程
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

		// 杀死进程
		err = proc.Kill()
		if err != nil {
			fmt.Printf("Error killing process: %v\n", err)
			return
		}

		// 清理 PID 文件
		os.Remove(pidFile)

		fmt.Printf("Torcontroller listener at %s stopped successfully.\n", socketPath)
	},
}
