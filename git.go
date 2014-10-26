package main

import (
	"github.com/libgit2/git2go"
	"os"
)

func (a *App) MergePullRequestIntoBranch(pr PullRequest, repo Repository, branch string) error {
	repoPath := a.Config.RepoDataPath + "/" + repo.FullName
	exists, err := exists(repoPath)
	if err != nil {
		return err
	}
	if !exists {
		err = a.InitRepo(repo, repoPath)
		if err != nil {
			return err
		}
	}

}

func (a *App) InitRepo(repo Repository, filePath string) error {
	err := os.MkdirAll(filePath)
	if err != nil {
		return err
	}
	err := gitCmd("git clone " + repo.GetCloneURL(a.Config.GitHubToken) + " " + filePath)
	if err != nil {
		return err
	}
	return nil
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
