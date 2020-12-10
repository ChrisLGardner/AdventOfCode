package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day10.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n")
	adapters := make([]int, len(split))
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
		adapters[i] = n
	}

	sort.Ints(adapters)

	oneJolt := 0
	threeJolt := 1
	previous := 0
	seen := make(map[int]int)
	for i, v := range adapters {
		seen[v] = i
		if diff := v - previous; diff == 1 {
			oneJolt++
		} else if diff == 3 {
			threeJolt++
		}
		previous = v
	}

	fmt.Printf("Total one jolts: %d, thee jolts: %d, final total: %d\n", oneJolt, threeJolt, oneJolt*threeJolt)

	combos := make([]int, len(adapters))
	combos[len(combos)-1] = 1

	for i := len(adapters) - 2; i >= 0; i-- {
		sum := 0
		for j := 1; j <= 3; j++ {
			if pos, ok := seen[adapters[i]+j]; ok {
				sum += combos[pos]
			}
		}
		combos[i] = sum
	}

	ret := 0
	for v := 1; v <= 3; v++ {
		if pos, ok := seen[v]; ok {
			ret += combos[pos]
		}
	}
	fmt.Println(ret)
}
