package api

import "net/http"

func NewRouter(webhookHandler http.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("POST /webhooks/github", GithubWebhookHandler)

	return mux
}
