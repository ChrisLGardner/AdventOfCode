package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day06.txt", "Relative file path to use as input.")

var questions = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n\n")
	total := 0
	for _, s := range split {
		for _, q := range questions {
			if strings.Contains(s, q) {
				total++
			}
		}
	}

	fmt.Printf("Part 1: %d", total)

	//part2
	total = 0
	for _, s := range split {
		passengers := strings.Count(s, "\n") + 1
		for _, q := range questions {
			if strings.Count(s, q) == passengers {
				total++
			}
		}
	}

	fmt.Printf("Part 2: %d", total)
}
