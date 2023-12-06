package grit

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func Expand(s string) string {
	home, _ := os.UserHomeDir()
	return strings.Replace(s, "~", home, 1)
}

func GetCurrentBranch(repoPath string) (string, error) {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	h, err := r.Head()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return h.Name().Short(), nil
}

func CheckoutBranch(repoPath string, branchName string, wg *sync.WaitGroup, errCh chan error) {
	defer wg.Done()
	log.Printf("Checking out repository %s at %s", repoPath, branchName)
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Fatal(err)
		errCh <- err
		return
	}

	w, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
		errCh <- err
		return
	}

	// branchRef, err := r.Branch(branchName)
	// if err != nil {
	// 	log.Fatal(err)
	// 	errCh <- err
	// 	return
	// }
	branchref := "refs/heads/" + branchName
	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(branchref),
	})
	if err != nil {
		log.Println("failed checkout")
		log.Fatal(err)
		errCh <- err
		return
	}
}

func CloneRepository(repoPath string, remotePath string) error {
	log.Printf("Cloning repository %s to %s", remotePath, repoPath)

	_, err := git.PlainClone(Expand(repoPath), false, &git.CloneOptions{
		URL:      remotePath,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func PullRepository(repoPath string, remotePath string, wg *sync.WaitGroup, errCh chan error) {
	defer wg.Done()
	expandedPath := Expand(repoPath)
	if _, err := os.Stat(expandedPath); os.IsNotExist(err) {
		err = CloneRepository(repoPath, remotePath)
		if err != nil {
			log.Fatal(err)
			errCh <- err
			return
		}
		return
	}

	log.Printf("Pulling repository %s", repoPath)
	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Fatal(err)
		errCh <- err
		return
	}

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
		errCh <- err
		return
	}

	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		log.Fatal(err)
		errCh <- err
		return
	}

	ref, err := r.Head()
	if err != nil {
		log.Fatal(err)
		errCh <- err
		return
	}

	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		log.Fatal(err)
		errCh <- err
		return
	}

	log.Println(commit)
}
