package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day15.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n")
	nums := make([]int, 0)
	numbers := make(map[int]map[string]int)
	for i, s := range strings.Split(split[0], ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}

		nums = append(nums, n)
		numbers[n] = map[string]int{"first": i + 1, "last": i + 1}
	}

	lastNumber := nums[len(nums)-1]

	for i := len(nums) + 1; i <= 30000000; i++ {

		if numbers[lastNumber]["first"] == numbers[lastNumber]["last"] {
			lastNumber = 0
		} else {
			lastNumber = numbers[lastNumber]["last"] - numbers[lastNumber]["first"]
		}

		if numbers[lastNumber] == nil {
			numbers[lastNumber] = map[string]int{"first": i, "last": i}
		} else {
			numbers[lastNumber]["first"] = numbers[lastNumber]["last"]
			numbers[lastNumber]["last"] = i
		}
	}
	fmt.Println(lastNumber)
	//fmt.Println(numbers)
}
