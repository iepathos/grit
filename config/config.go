package grit

import (
	"log"

	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Repositories map[string]string
}

func FileExists(filePath string) (bool, error) {
	info, err := os.Stat(filePath)
	if err == nil {
		return !info.IsDir(), nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func GetDefaultYml() string {
	// either grit.yml in local directory will be used
	// or if it doesn't exist ~/.grit.yml will be used
	relative_grit := "grit.yml"
	exists, err := FileExists(relative_grit)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if exists {
		return relative_grit
	} else {
		return "~/.grit.yml"
	}
}

func ParseYml(ymlpath string) map[string]string {
	conf := Config{}

	data, err := os.ReadFile(ymlpath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	// log.Printf("--- conf:\n%v\n\n", conf)
	return conf.Repositories
}
