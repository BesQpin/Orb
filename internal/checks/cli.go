package checks

import (
	"github.com/BesQpin/orb/internal/checks/dns"
	"github.com/BesQpin/orb/internal/checks/httpcheck"
	"github.com/BesQpin/orb/internal/checks/tcp"
	"github.com/spf13/cobra"
)

func RegisterCLIChecks(cmd *cobra.Command) {
	cmd.AddCommand(dns.CLICmd())
	cmd.AddCommand(tcp.CLICmd())
	cmd.AddCommand(httpcheck.CLICmd())
}
