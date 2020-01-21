package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type wmap map[city][]city

func newMapFromFile(filename string) wmap {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	worldMap := make(wmap)

	lines := strings.Split(string(bs), "\n")

	for _, line := range lines {
		curLine := splitLineByWhitespace(line)

		for i := 0; i < len(curLine); i++ {
			curCity := city(curLine[0])
			var connections []city
			for _, d := range curLine[1:] {
				connections = append(connections, city(splitLineByEqual(d)[1]))
			}

			worldMap[curCity] = connections
		}
	}

	return worldMap
}

func (wm wmap) print() {
	for city, connections := range wm {
		fmt.Println(city, ":", connections)
	}
}

func splitLineByWhitespace(s string) []string {
	return strings.Split(s, " ")
}

func splitLineByEqual(s string) []string {
	return strings.Split(s, "=")
}
