package server

import (
	"encoding/json"
	"net/http"
)

type statusResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	response := statusResponse{"OK", Version}

	bytes, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.Write(bytes)
}
