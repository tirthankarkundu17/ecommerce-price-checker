package handler

import (
	"encoding/json"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Hello World!"})

}
