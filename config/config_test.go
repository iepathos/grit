package grit

import (
	"testing"
)

func Test_ParseYml(t *testing.T) {
	res := ParseYml("example_config.yml")
	expected := map[string]string{
		"~/devl/grit":               "git@github.com:iepathos/grit.git",
		"~/devl/dryenv":             "git@github.com:iepathos/dryenv.git",
		"~/devl/aistr":              "git@github.com:iepathos/aistr.git",
		"/path/that/does/not/exist": "git@github.com:iepathos/someinvalidrepo",
	}
	for localPath, remotePath := range res {
		if expected[localPath] != remotePath {
			t.Fatalf("localPath %s %s remotePath did not match expected %s", localPath, remotePath, expected[localPath])
		}
	}
}
