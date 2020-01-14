package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type application struct {
	Version     string `json:"version"`
	Commit      string `json:"lastCommitSha"`
	Description string `json:"description"`
}

type allApplications []application

var myapplication allApplications

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello ANZ-X!")
}

func getAppVersion(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(myapplication)
}

func initApplications() {
	myapplication = allApplications{
		{
			Version:     version,
			Commit:      commit,
			Description: "Release for ANZ Platform Test",
		},
	}
}
