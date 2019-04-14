package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// home returns a simple HTTP handler function which writes a response.
func home(version, commit, repo string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		info := struct {
			Version   string `json:"version"`
			Commit    string `json:"commit"`
			Repo      string `json:"repo"`
		}{
			version, commit, repo,
		}

		body, err := json.Marshal(info)
		if err != nil {
			log.Printf("Could not encode info data: %v", err)
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}
