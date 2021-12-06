package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day06.txt", "Relative file path to use as input.")
var part2 = flag.Bool("part2", false, "Run logic for part 2 of puzzle")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), ",")
	fish := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	for _, s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}
		if n <= 0 {
			fmt.Printf("Optimization invariant broken: %d <= 0 \n", n)
			break
		}
		fish[n] += 1
	}

	days := 80

	if *part2 {
		days = 256
	}

	for i := 1; i <= days; i++ {
		fish = DoDay(fish)
	}

	total := 0

	for _, v := range fish {
		total += v
	}
	fmt.Printf("Number of fish: %d\n", total)
}

func DoDay(fish map[int]int) map[int]int {

	zeroes := fish[0]

	for i := 1; i < 9; i++ {
		fish[i-1] = fish[i]
	}

	fish[6] += zeroes
	fish[8] = zeroes

	return fish
}
