package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (a *App) Receive(w http.ResponseWriter, r *http.Request) {
	body, err := getBody(r)
	if err != nil {
		a.Logger.Error(err.Error())
	}

	action := r.Header.Get("X-Github-Event")
	switch {
	case action == "pull_request":
		pr := PullRequestEvent{}
		err = json.Unmarshal(body, pr)
		if err != nil {
			a.Logger.Error(err.Error())
		}

		err = a.ProcessPullRequestPayload(pr)
		if err != nil {
			a.Logger.Error(err.Error())
		}
		break
	}
}

func getBody(r *http.Request) ([]byte, error) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return body, err
	}
	return body, nil
}
