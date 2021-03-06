package main

type Config struct {
	HookAddress  string            `json:"hook-address"`
	RepoDataPath string            `json:"repo-data-path"`
	GitHubToken  string            `json:"github-token"`
	Repos        map[string]Events `json:"repos"`
}

type Events []Event
type Event struct {
	Event  string `json:"event"`
	Label  string `json:"label"`
	Action Action `json:"action"`
}

type Action struct {
	Type   string            `json:"type"`
	Params map[string]string `json:"params"`
}
