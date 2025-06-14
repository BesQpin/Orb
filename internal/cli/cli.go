package cli

import (
	"fmt"
	"os"

	"github.com/BesQpin/orb/internal/checks"
	"github.com/BesQpin/orb/internal/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "orb",
	Short: "Orb is a connectivity testing tool.",
}

var (
	mode string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&mode, "mode", "cli", "Mode to run in: cli or http")
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		switch mode {
		case "cli":
			checks.RegisterCLIChecks(cmd)
			cmd.Execute()
		case "http":
			server.Start()
		default:
			fmt.Println("Invalid mode. Use 'cli' or 'http'")
			os.Exit(1)
		}
	}
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return rootCmd.Execute()
}
