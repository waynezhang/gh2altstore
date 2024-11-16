package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleRepo(w http.ResponseWriter, r *http.Request) {
	owner := r.PathValue("owner")
	repo := r.PathValue("repo")

	altRepo, err := getReleases(fmt.Sprintf("%s/%s", owner, repo))
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

func serve(addr string) {
	http.HandleFunc("/repos/{owner}/{repo}", handleRepo)
	log.Fatal(http.ListenAndServe(addr, nil))
}
