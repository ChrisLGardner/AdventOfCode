package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day09.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n")
	input := make([]int, len(split))
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
		input[i] = n
	}

	problemIndex := 0

outer:
	for i := 25; i < len(input); i++ {

		for j := i - 25; j < i; j++ {

		second:
			for k := i - 24; k < i; k++ {
				if j == k {
					continue second
				}
				if input[j]+input[k] == input[i] {
					continue outer
				}
			}
		}
		problemIndex = i
		break
	}
	fmt.Println(problemIndex)
	fmt.Println(input[problemIndex])

main:
	for i := 0; i < problemIndex; i++ {

		total := input[i]
		contiguous := []int{input[i]}

	inner:
		for j := i + 1; j < problemIndex; j++ {
			total += input[j]
			if total > input[problemIndex] {
				continue main
			}
			contiguous = append(contiguous, input[j])

			if total == input[problemIndex] {
				break inner
			}
		}

		lowest := input[problemIndex]
		highest := 0

		for _, v := range contiguous {
			if v < lowest {
				lowest = v
			}
			if v > highest {
				highest = v
			}
		}

		fmt.Println(lowest + highest)
		break main

	}
}
