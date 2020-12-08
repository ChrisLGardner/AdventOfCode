package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day08.txt", "Relative file path to use as input.")

type instruction struct {
	action string // acc, nop, jmp
	sign   string // +, -
	amount int
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n")
	instructionSet := make([]instruction, 0)
	for _, s := range split {
		parts := strings.Split(s, " ")
		sign := string(parts[1][0])
		value, err := strconv.Atoi(parts[1][1:])
		if err != nil {
			fmt.Printf("Failed to parse %s", parts[1])
			continue
		}

		ins := instruction{
			parts[0],
			sign,
			value,
		}

		instructionSet = append(instructionSet, ins)
	}

	//part 1
	var accumulator int
	visited := make(map[int]bool)

	i := 0
	for {
		current := instructionSet[i]

		if previous := visited[i]; previous == true {
			break
		}

		visited[i] = true

		if current.action == "acc" {
			if current.sign == "+" {
				accumulator += current.amount
			} else if current.sign == "-" {
				accumulator -= current.amount
			}
		} else if current.action == "jmp" {
			if current.sign == "+" {
				i += current.amount
			} else if current.sign == "-" {
				i -= current.amount
			}
			continue
		}

		i++
		continue
	}

	fmt.Println(accumulator)

	//part 2
	for i, ins := range instructionSet {
		if ins.action == "acc" {
			continue
		}

		originalAction := ins.action

		if ins.action == "jmp" {
			ins.action = "nop"
		} else if ins.action == "nop" {
			ins.action = "jmp"
		}

		instructionSet[i] = ins

		if output := runInstructions(instructionSet); output != 0 {
			fmt.Println(output)
			break
		}

		ins.action = originalAction
		instructionSet[i] = ins

	}
}

func runInstructions(input []instruction) int {
	var accumulator int
	visited := make(map[int]bool)

	i := 0
	for {
		current := input[i]

		if previous := visited[i]; previous == true {
			return 0
		}

		visited[i] = true

		if current.action == "acc" {
			if current.sign == "+" {
				accumulator += current.amount
			} else if current.sign == "-" {
				accumulator -= current.amount
			}
		} else if current.action == "jmp" {
			if current.sign == "+" {
				i += current.amount
			} else if current.sign == "-" {
				i -= current.amount
			}
			if i >= len(input) {
				return accumulator
			}
			continue
		}

		i++
		if i >= len(input) {
			break
		}
		continue
	}

	return accumulator
}
