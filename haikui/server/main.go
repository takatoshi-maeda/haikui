package server

import (
	"net/http"
	"time"
)

func Start() {
	New().ListenAndServe()
}

func New() *http.Server {
	mux := http.NewServeMux()
	registerHandler(mux, "/stats", StatsHandler())

	httpServer := &http.Server{
		Addr:         "localhost:8989",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      mux,
	}

	return httpServer
}

func registerHandler(mux *http.ServeMux, path string, h *Handler) {
	mux.HandleFunc(path, h.Serve)
}
