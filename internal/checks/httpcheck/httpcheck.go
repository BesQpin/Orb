package httpcheck

import (
	"fmt"
	"net/http"
	"time"
	"github.com/spf13/cobra"
)

func CLICmd() *cobra.Command {
	return &cobra.Command{
		Use:   "http [url]",
		Short: "Check HTTP GET request",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			checkHTTP(args[0])
		},
	}
}

func checkHTTP(url string) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("HTTP error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("HTTP status: %s\n", resp.Status)
}

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "missing 'url' param", http.StatusBadRequest)
		return
	}
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	fmt.Fprintf(w, "HTTP status: %s", resp.Status)
}