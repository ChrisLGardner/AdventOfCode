package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day01.txt", "Relative file path to use as input.")
var part2 = flag.Bool("part2", false, "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n")
	seen := make([]int, len(split))
	for i, s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}
		if n <= 0 {
			fmt.Printf("Optimization invariant broken: %d <= 0 \n", n)
			break
		}
		seen[i] = n
	}

	count := 0
	prev := 0

	if !*part2 {
		for _, v := range seen {
			if v > prev && prev > 0 {
				count++
			}
			prev = v
		}
	} else {
		for i, _ := range seen {
			if i+2 >= len(seen) {
				break
			}

			total := seen[i] + seen[i+1] + seen[i+2]

			if prev == 0 {
				prev = total
				continue
			}

			if total > prev {
				count++
			}

			prev = total
		}
	}

	fmt.Printf("Count is %d \n", count)
}
