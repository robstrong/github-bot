package main

type PullRequestEvent struct {
	Action      string      `json:"action"`
	Number      int         `json:"number"`
	PullRequest PullRequest `json:"pull_request"`
}

type PullRequest struct {
	Number int    `json:"number"`
	State  string `json:"state"`
	Head   Branch `json:"head"`
	Base   Branch `json:"base"`
}

type Branch struct {
	Ref  string     `json:"ref"`
	Sha  string     `json:"sha"`
	Repo Repository `json:"repo"`
}

type Repository struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Owner    User   `json:"user"`
}

type User struct {
	Login string `json:"login"`
	Id    int    `json:"id"`
}
