package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day05.txt", "Relative file path to use as input.")

type seat struct {
	row    string //F|B{7} 0 - 127
	column string // R|L{3} 0 - 8
	id     int    // row * 8 + column
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n")

	seats := make(map[int]seat)
	highest := 0

	for _, s := range split {
		if len(s) != 10 {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}

		row := 0
		rowMax := 127

		for i := 0; i < 7; i++ {
			row, rowMax = binSearch(row, rowMax, string(s[i]))

			if i == 6 {
				if string(s[i]) == "B" {
					row = rowMax
				}
			}
		}

		column := 0
		columnMax := 7

		for i := 7; i < 10; i++ {
			column, columnMax = binSearch(column, columnMax, string(s[i]))

			if i == 9 {
				if string(s[i]) == "R" {
					column = columnMax
				}
			}
		}

		id := row*8 + column
		seats[id] = seat{s[0 : len(s)-3], s[7:], id}

		if id > highest {
			highest = id
		}
	}

	fmt.Printf("Highest ticket id: %d\n", highest)

	for k := range seats {
		if seats[k+2].row != "" && seats[k+1].row == "" {
			fmt.Printf("Ticket should be %d because %d found\n", k+1, k+2)
		}
	}
}

func binSearch(min, max int, half string) (newMin, newMax int) {
	if half == "F" || half == "L" {
		newMin := min
		newMax := ((max-min)/2 + min)
		return newMin, newMax
	} else if half == "B" || half == "R" {
		newMin := ((max-min)/2 + min) + 1
		newMax := max
		return newMin, newMax
	}

	return
}
