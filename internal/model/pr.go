package model

type WebhookPayload struct {
	Action string `json:"action"`
	Number int    `json:"number"`
}

type PullRequest struct {
}
