package grit

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func Expand(s string) string {
	home, _ := os.UserHomeDir()
	return strings.Replace(s, "~", home, 1)
}

func PullRepository(repoPath string, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("Executing git pull in %s", repoPath)
	cmd := exec.Command("git", "pull")
	cmd.Dir = Expand(repoPath)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", out)
}
