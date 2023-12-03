package grit

import (
	// "fmt"
	"log"

	"gopkg.in/yaml.v3"
	// "bufio"
	// "io"
	"os"
)

type Config struct {
	Repositories []string
}

func ParseYml(ymlpath string) []string {
	conf := Config{}

	data, err := os.ReadFile(ymlpath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("--- t:\n%v\n\n", conf)
	return conf.Repositories
}
