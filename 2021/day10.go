package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
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
	corruptTotal := 0
	incompleteTotals := []int{}

line:
	for _, s := range split {
		stack := []rune{}
		for _, r := range s {
			if r == '[' || r == '{' || r == '(' || r == '<' {
				stack = append(stack, r)
			} else {
				if r == ')' && stack[len(stack)-1] != '(' {
					corruptTotal += 3
					continue line
				} else if r == ']' && stack[len(stack)-1] != '[' {
					corruptTotal += 57
					continue line
				} else if r == '}' && stack[len(stack)-1] != '{' {
					corruptTotal += 1197
					continue line
				} else if r == '>' && stack[len(stack)-1] != '<' {
					corruptTotal += 25137
					continue line
				} else {
					stack = stack[:len(stack)-1]
				}
			}

		}
		if len(stack) > 0 {
			lineTotal := 0
			for i := len(stack) - 1; i > -1; i-- {
				lineTotal *= 5
				if stack[i] == '(' {
					lineTotal += 1
				} else if stack[i] == '[' {
					lineTotal += 2
				} else if stack[i] == '{' {
					lineTotal += 3
				} else if stack[i] == '<' {
					lineTotal += 4
				}
			}
			incompleteTotals = append(incompleteTotals, lineTotal)
		}
	}

	midTotal := len(incompleteTotals) / 2
	sort.Ints(incompleteTotals)

	fmt.Printf("Corrupt Total is: %d\n", corruptTotal)
	fmt.Printf("Incomplete total is: %d\n", incompleteTotals[midTotal])

}
