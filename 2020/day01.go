package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day01.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n")
	//split = split[:len(split)-1]
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

partA:
	for i, m := range seen {
		for j, n := range seen {

			if i >= j {
				continue
			}
			if m+n == 2020 {
				fmt.Println(m * n)
				break partA
			}
		}
	}

partB:
	for i, m := range seen {
		for j, n := range seen {
			for k, o := range seen {

				if i >= j || j >= k {
					continue
				}
				if m+n+o == 2020 {
					fmt.Println(m * n * o)
					break partB
				}
			}
		}
	}
}
