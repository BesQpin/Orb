package cli

import (
	"fmt"
	"os"

	"github.com/BesQpin/orb/internal/checks"
	"github.com/BesQpin/orb/internal/server"
	"github.com/spf13/cobra"
)

var (
	mode    string
	rootCmd = &cobra.Command{
		Use:   "orb",
		Short: "Orb is a connectivity testing tool.",
		Run: func(cmd *cobra.Command, args []string) {
			if mode == "cli" {
				fmt.Println("No check provided. Use one of: dns, tcp, http.")
				_ = cmd.Help()
				os.Exit(1)
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&mode, "mode", "cli", "Mode to run in: cli or http")
}

func Execute() error {
	switch mode {
	case "cli":
		checks.RegisterCLIChecks(rootCmd)
		return rootCmd.Execute()
	case "http":
		return server.Start()
	default:
		fmt.Fprintln(os.Stderr, "❌ Invalid mode. Use 'cli' or 'http'")
		os.Exit(1)
		return nil
	}
}
