package tcp

import (
	"fmt"
	"net"
	"net/http"
	"time"
	"github.com/spf13/cobra"
)

func CLICmd() *cobra.Command {
	return &cobra.Command{
		Use:   "tcp [host:port]",
		Short: "Check TCP connectivity",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			checkTCP(args[0])
		},
	}
}

func checkTCP(address string) {
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		fmt.Println("TCP error:", err)
		return
	}
	defer conn.Close()
	fmt.Println("TCP connection successful")
}

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	addr := r.URL.Query().Get("addr")
	if addr == "" {
		http.Error(w, "missing 'addr' param", http.StatusBadRequest)
		return
	}
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	conn.Close()
	fmt.Fprintf(w, "TCP connection to %s successful", addr)
}