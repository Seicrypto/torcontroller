package cmd

import (
	"github.com/Seicrypto/torcontroller/internal/singleton/logger"
	"github.com/spf13/cobra"
)

// Root Command
var rootCmd = &cobra.Command{
	Use:   "torcontroller",
	Short: "Tor Controller CLI",
	Long:  "A CLI to control Tor and Privoxy services.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Initialization Log
		// logger.GetLogger()

		return nil
	},
}

var socketPath = "/tmp/torcontroller.sock"

var pidFile = "/tmp/torcontroller.pid"

var log *logger.Logger

// Initialize Root Command
func InitCommands() *cobra.Command {
	// Initialization Log
	log = logger.GetLogger()

	rootCmd.AddCommand(
		VersionCmd,
		CheckCmd,
		StartCmd,
		StartBackgroundCmd,
		TrafficCmd,
		StatusCmd,
		SwitchCmd,
		StopCmd,
		NewPasswordCmd,
	)

	CheckCmd.Flags().BoolVarP(&fixFlag, "fix", "f", false, "Fix missing or incorrect results")

	return rootCmd
}
