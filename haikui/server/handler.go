package server

import (
	"net/http"
)

type Handler struct {
	Serve func(w http.ResponseWriter, r *http.Request)
}
