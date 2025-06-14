package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/BesQpin/orb/internal/cli"
	"github.com/BesQpin/orb/internal/server"
)

func main() {
	mode := flag.String("mode", "cli", "Mode to run: cli or http")
	flag.Parse()

	switch *mode {
	case "cli":
		if err := cli.Execute(); err != nil {
			fmt.Fprintln(os.Stderr, "❌ Error:", err)
			os.Exit(1)
		}
	case "http":
		if err := server.Start(); err != nil {
			log.Fatal("❌ HTTP server failed:", err)
		}
	default:
		fmt.Fprintln(os.Stderr, "❌ Invalid mode:", *mode)
		os.Exit(1)
	}
}
