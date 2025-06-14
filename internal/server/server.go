package server

import (
	"log"
	"net/http"

	"github.com/BesQpin/orb/internal/checks/dns"
	"github.com/BesQpin/orb/internal/checks/httpcheck"
	"github.com/BesQpin/orb/internal/checks/tcp"
)

func Start() error {
	http.HandleFunc("/healthz/live", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("live"))
	})

	http.HandleFunc("/healthz/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ready"))
	})

	http.HandleFunc("/check/dns", dns.HTTPHandler)
	http.HandleFunc("/check/tcp", tcp.HTTPHandler)
	http.HandleFunc("/check/http", httpcheck.HTTPHandler)

	log.Println("🌐 Starting HTTP server on :8080")
	return http.ListenAndServe(":8080", nil)
}
