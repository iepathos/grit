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

func ParseYml(ymlpath string) (map[string]string, error) {
	conf := Config{}

	data, err := os.ReadFile(ymlpath)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}

	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}
	// log.Printf("--- conf:\n%v\n\n", conf)
	return conf.Repositories, nil
}
