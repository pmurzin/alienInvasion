package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const fileName string = "worldmap.txt"

func TestNewMapFromFile(t *testing.T) {
	testMap := newMapFromFile(fileName)

	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Errorf("Error: %v", err)
		os.Exit(1)
	}

	lines := strings.Split(string(bs), "\n")

	if len(testMap) != len(lines) {
		t.Errorf("Expected map length of %v, but got %v", len(lines), len(testMap))
	}

	for _, line := range lines {
		curLine := splitLineByWhitespace(line)
		for c := range testMap {
			if string(c) == curLine[0] {
				connNum := len(curLine) - 1
				if len(testMap[c]) != connNum {
					t.Errorf("Expected connection num for city %v of %v, but got %v", c, connNum, len(testMap[c]))
				}
			}
		}
	}
}
