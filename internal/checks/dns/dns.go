package dns

import (
	"fmt"
	"net"
	"net/http"
	"github.com/spf13/cobra"
)

func CLICmd() *cobra.Command {
	return &cobra.Command{
		Use:   "dns [hostname]",
		Short: "Check DNS resolution",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ips, err := net.LookupHost(args[0])
			if err != nil {
				fmt.Println("DNS error:", err)
				return
			}
			fmt.Println("Resolved IPs:", ips)
		},
	}
}

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	host := r.URL.Query().Get("host")
	if host == "" {
		http.Error(w, "missing 'host' param", http.StatusBadRequest)
		return
	}
	ips, err := net.LookupHost(host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Resolved IPs: %v", ips)
}