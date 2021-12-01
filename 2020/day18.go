package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day18.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n")
	total := 0
	for _, s := range split {

		lineTotal, _ := parseLine(strings.Replace(s, " ", "", -1))

		total += lineTotal
	}
	fmt.Println(total)
}

func parseLine(input string) (total, lastIndex int) {

	total = 0
	nextOp := "+"
	lastIndex = -1
	for i, s := range input {
		if i <= lastIndex {
			continue
		}
		switch s {
		case '(':
			extra, newLastIndex := parseLine(input[i+1:])
			if nextOp == "+" {
				total += extra
			} else if nextOp == "*" {
				total *= extra
			}
			lastIndex = newLastIndex + i + 1
		case ')':
			return total, i
		case '+':
			nextOp = "+"
		case '*':
			nextOp = "*"
		default:
			n, err := strconv.Atoi(string(s))
			if err != nil {
				fmt.Printf("Failed to parse %s\n", s)
				break
			}

			if nextOp == "+" {
				total += n
			} else if nextOp == "*" {
				total *= n
			}
		}
	}

	return total, 0
}
