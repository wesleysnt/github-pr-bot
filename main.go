package main

import (
	"net/http"
	"wesleysnt/github-pr-bot/internal/api"
)

func main() {
	s := http.NewServeMux()

	s = api.NewRouter(s)

	server := &http.Server{
		Addr:    ":8000",
		Handler: s,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
