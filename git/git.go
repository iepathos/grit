package grit

import (
	"log"
	"os"
	"os/exec"
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

	_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:      remotePath,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatal(err)
	}

	// get parent directory of local repo path for calling
	// parentDir := path.Dir(Expand(repoPath))

	// cmd := exec.Command("git", "clone", remotePath)
	// cmd.Dir = parentDir
	// out, err := cmd.Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%s", out)
}

func PullRepository(repoPath string, remotePath string, wg *sync.WaitGroup) {
	defer wg.Done()
	expandedPath := Expand(repoPath)
	if _, err := os.Stat(expandedPath); os.IsNotExist(err) {
		CloneRepository(repoPath, remotePath)
		return
	}

	log.Printf("Pulling repository %s", repoPath)
	cmd := exec.Command("git", "pull")
	cmd.Dir = expandedPath
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", out)
}
