package grit

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/go-git/go-git/v5"
)

func Expand(s string) string {
	home, _ := os.UserHomeDir()
	return strings.Replace(s, "~", home, 1)
}

func CloneRepository(repoPath string, remotePath string) {
	log.Printf("Cloning repository %s to %s", remotePath, repoPath)

	_, err := git.PlainClone(Expand(repoPath), false, &git.CloneOptions{
		URL:      remotePath,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func PullRepository(repoPath string, remotePath string, wg *sync.WaitGroup) {
	defer wg.Done()
	expandedPath := Expand(repoPath)
	if _, err := os.Stat(expandedPath); os.IsNotExist(err) {
		CloneRepository(repoPath, remotePath)
		return
	}

	log.Printf("Pulling repository %s", repoPath)
	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Fatal(err)
	}

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}

	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		log.Fatal(err)
	}

	// Print the latest commit that was just pulled
	ref, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		log.Fatal(err)
	}

	log.Println(commit)

}
