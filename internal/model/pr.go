package model

type WebhookPayload struct {
	Action      string      `json:"action"`
	Number      int         `json:"number"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
}

type PullRequest struct {
	DiffUrl string `json:"diff_url"`
	Head    Commit `json:"head"`
}

type Repository struct {
	Name  string `json:"name"`
	Owner Owner  `json:"owner"`
}

type Commit struct {
	SHA string `json:"sha"`
}

type Owner struct {
	Login string `json:"login"`
}
