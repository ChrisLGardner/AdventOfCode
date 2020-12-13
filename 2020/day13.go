package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day13.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n")
	arrivalTime, err := strconv.Atoi(split[0])
	if err != nil {
		fmt.Printf("Failed to parse %s\n", split[0])
	}

	earliest := 2 * arrivalTime
	bus := 0
	difference := 0
	buses := strings.Split(split[1], ",")
	busInts := make(map[int]int)
	for i, s := range buses {
		if s == "x" {
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}
		busInts[n] = i
		// closest a bus will arrive to arrival time without going over
		// then add bus time again to get closest after arrival time
		busEarliest := (int(math.Floor(float64(arrivalTime/n))) * n) + n

		if busEarliest < earliest && busEarliest > arrivalTime {
			earliest = busEarliest
			bus = n
			difference = busEarliest - arrivalTime
		}
	}

	fmt.Println(bus * difference)

	// iterate until we find a number which i % target == 0
	// then iterate by that number until we find the next one (i + 1) % target1 == 0
	// then for maths reasons we can target * target1 and iterate by that
	// once we do that each iteration will match the previous two checks
	// this means we can then start checking (i + 2) % target 2 == 0
	// then repeat the process of target * target1 * target2...
	minValue := 0
	runningProduct := 1
	for k, v := range busInts {
		for (minValue+v)%k != 0 {
			minValue += runningProduct
		}
		runningProduct *= k
		fmt.Printf("t + %d === 0 mod %d\n", v, k)
		fmt.Printf("Sum so far: %d, product so far: %d\n", minValue, runningProduct)
	}
	fmt.Println(minValue)

}
