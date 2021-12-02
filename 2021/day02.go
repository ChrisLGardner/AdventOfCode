package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day02.txt", "Relative file path to use as input.")

type instruction struct {
	direction string
	distance  int
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(strings.TrimRight(contents, "\n"), "\n")
	instructions := make([]instruction, len(lines))
	for i, s := range lines {
		split := strings.Split(s, " ")
		n, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}

		instructions[i] = instruction{
			direction: split[0],
			distance:  n,
		}
	}

	horizontal := 0
	vertical := 0
	aim := 0

	for _, ins := range instructions {
		if ins.direction == "forward" {
			horizontal += ins.distance
			vertical += aim * ins.distance
		} else if ins.direction == "down" {
			aim += ins.distance
		} else if ins.direction == "up" {
			aim -= ins.distance
		}
	}

	fmt.Printf("final position horizontal: %d, vertical: %d. Total: %d\n", horizontal, vertical, horizontal*vertical)
}
