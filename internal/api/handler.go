package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GithubWebhookHandler(w http.ResponseWriter, r *http.Request) {
	event := r.Header["X-Github-Event"]
	if event[0] == "pull_request" {
		req := json.NewDecoder(r.Body)
		fmt.Print(req)

	}
	fmt.Fprintf(w, "test", r)
}
