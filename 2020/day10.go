package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
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

	oneJolt := 1
	threeJolt := 1

	for i := 0; i < (len(adapters) - 1); i++ {
		if diff := adapters[i+1] - adapters[i]; diff == 1 {
			oneJolt++
			continue
		} else if diff == 3 {
			threeJolt++
			continue
		}
	}

	fmt.Printf("Total one jolts: %d, thee jolts: %d, final total: %d\n", oneJolt, threeJolt, oneJolt*threeJolt)

	seen := make(map[int]int)
	total := big.NewInt(0)

	for i := 0; i <= 3; i++ {
		if adapters[i] == 1 || adapters[i] == 2 || adapters[i] == 3 {
			total.Add(total, big.NewInt(1))
			continue
		}
	}
outer:
	for i, v := range adapters {
	inner:
		for j := 1; j <= 3; j++ {
			if i == len(adapters)-1 {
				seen[v]++
				continue outer
			} else if i+j >= len(adapters) {
				continue outer
			}

			for k := 1; k <= 3; k++ {
				if v+j == adapters[i+k] {
					seen[v]++
					continue inner
				}
			}
		}

		total.Add(total, big.NewInt(int64(seen[v]*(i+1))))
	}

	//total.Add(total, big.NewInt(int64()))

	fmt.Println(adapters)
	fmt.Println(seen)
	fmt.Println(total)
}
