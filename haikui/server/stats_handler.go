package server

import (
	"net/http"
)

func StatsHandler() *Handler {
	handler := &Handler{
		Serve: statsServe,
	}
	return handler
}

func statsServe(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
