package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wesleysnt/github-pr-bot/internal/github"
	"wesleysnt/github-pr-bot/internal/model"
)

func GithubWebhookHandler(w http.ResponseWriter, r *http.Request) {
	event := r.Header["X-Github-Event"]
	if event[0] != "pull_request" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var payload model.WebhookPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		fmt.Printf("Error: %v", err)
		http.Error(w, "Failed to parse json payload", http.StatusBadRequest)
		return
	}

	if payload.Action == "opened" || payload.Action == "synchronized" || payload.Action == "reopened" {
		fmt.Printf("Received PR #%d from %s/%s\n",
			payload.Number,
			payload.Repository.Owner.Login,
			payload.Repository.Name,
		)
		fmt.Printf("Diff URL: %s\n", payload.PullRequest.DiffUrl)
		github.ProcessDiff(payload.PullRequest.DiffUrl)
	}
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, "Webhook received successfully")
}
