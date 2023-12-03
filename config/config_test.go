package grit

import (
	"testing"
)

func Test_ParseYml(t *testing.T) {
	res := ParseYml("example_config.yml")
	expected := []string{
		"~/devl/grit",
		"~/devl/dryenv",
		"~/devl/aistr",
	}
	if res[0] != expected[0] {
		t.Fatalf("res %s did not mtch expected %s", res, expected)
	}
	if res[1] != expected[1] {
		t.Fatalf("res %s did not mtch expected %s", res, expected)
	}
	if res[2] != expected[2] {
		t.Fatalf("res %s did not mtch expected %s", res, expected)
	}
}
