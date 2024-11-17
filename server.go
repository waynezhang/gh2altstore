package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/syumai/workers"
)

func handleRepo(w http.ResponseWriter, r *http.Request) {
	// owner := r.PathValue("owner")
	// repo := r.PathValue("repo")
	path := r.URL.Path
	segments := strings.Split(strings.Trim(path, "/"), "/")
	if len(segments) != 3 {
		http.NotFound(w, r)
		return
	}

	owner := segments[1]
	repo := segments[2]

	altRepo, err := getReleases(fmt.Sprintf("%s/%s", owner, repo), r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, err := json.MarshalIndent(altRepo, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(jsonData)
}

func serve() {
	// http.HandleFunc("/repos/{owner}/{repo}", handleRepo)
	http.HandleFunc("/repos/", handleRepo)
	workers.Serve(nil)
}
