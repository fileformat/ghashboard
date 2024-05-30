package main

import (
	"os"
	"regexp"
)

var lineSplitter = regexp.MustCompile(`[\s,]+`)

func loadAtFile(filename string) ([]string, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := lineSplitter.Split(string(contents), -1)
	return lines, nil
}
