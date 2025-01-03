package cmd

import (
	"fmt"

	"github.com/Seicrypto/torcontroller/internal/singleton/configuration"
	"github.com/Seicrypto/torcontroller/internal/singleton/logger"
	"github.com/spf13/cobra"
)

// Root Command
var rootCmd = &cobra.Command{
	Use:   "torcontroller",
	Short: "Tor Controller CLI",
	Long:  "A CLI to control Tor and Privoxy services.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		// Initialize configuration
		configurationPath := "/etc/torcontroller/torcontroller.yml"
		if err := configuration.LoadConfig(configurationPath); err != nil {
			return fmt.Errorf("failed to load configuration: %v", err)
		}

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
		InitCmd,
		StartCmd,
		StartBackgroundCmd,
		TrafficCmd,
		SwitchCmd,
		StopCmd,
		NewPasswordCmd,
	)

	CheckCmd.Flags().BoolVarP(&fixFlag, "fix", "f", false, "Fix missing or incorrect results")

	return rootCmd
}
