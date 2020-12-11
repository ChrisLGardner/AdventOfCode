package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day11.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n")
	seats := make([][]string, len(split))
	for i, s := range split {
		s = strings.Replace(s, "L", "#", -1)
		n := strings.Split(s, "")

		seats[i] = n
	}

	changed := true
	for {
		seats, changed = tick(seats)

		if !changed {
			break
		}
	}

	count := 0
	for _, row := range seats {
		for _, col := range row {
			if col == "#" {
				count++
			}
		}
	}

	fmt.Println(count)

}

func tick(state [][]string) ([][]string, bool) {
	original := make([][]string, len(state))
	for i, v := range state {
		new := make([]string, len(v))
		copy(new, v)
		original[i] = new
	}

	changed := false
	for i, row := range state {
		for j, col := range row {
			if col == "." || (i == 0 && j == 0) || (i == 0 && j == len(row)-1) || (i == len(state)-1 && j == 0) || (i == len(state)-1 && j == len(row)-1) {
				continue
			}
			seat := changeSeat(original, i, j)

			if seat != col {
				changed = true
				state[i][j] = seat
			}
		}
	}

	return state, changed
}

func changeSeat(orig [][]string, row, col int) string {

	adjacentCount := 0
	var rowLeft, rowRight, colAbove, colBelow, leftAbove, leftBelow, rightAbove, rightBelow bool
	if row == 0 {
		rowLeft = true
		rowRight = true
		colBelow = true
		leftBelow = true
		rightBelow = true
	} else if row == len(orig)-1 {
		rowLeft = true
		rowRight = true
		colAbove = true
		leftAbove = true
		rightAbove = true
	} else {
		if col == 0 {
			rowRight = true
			colAbove = true
			colBelow = true
			leftAbove = true
			rightAbove = true
		} else if col == len(orig[row])-1 {
			rowLeft = true
			colAbove = true
			colBelow = true
			leftAbove = true
			rightAbove = true
		} else {
			rowLeft = true
			rowRight = true
			colAbove = true
			colBelow = true
			leftAbove = true
			rightAbove = true
			leftBelow = true
			rightBelow = true
		}
	}

	switch {
	case rowLeft:
	rowLeft:
		for i := 0; i < col; i++ {
			if orig[row][i] == "#" {
				adjacentCount++
				break rowLeft
			}
		}
		fallthrough
	case rowRight:
	rowRight:
		for i := col + 1; i < len(orig[row])-1; i++ {
			if orig[row][i] == "#" {
				adjacentCount++
				break rowRight
			}
		}
		fallthrough
	case colBelow:
	colBelow:
		for i := 1; i < len(orig)-1; i++ {
			if orig[i][col] == "#" {
				adjacentCount++
				break colBelow
			}
		}
		fallthrough
	case colAbove:
	colAbove:
		for i := 1; i < len(orig)-1; i++ {
			if orig[i][col] == "#" {
				adjacentCount++
				break colAbove
			}
		}
		fallthrough
	case leftBelow:
	rowLeftBelow:
		for i := 0; i < len(orig)-1; i++ {
			for j := 0; j < col; j++ {
				if orig[i][j] == "#" {
					adjacentCount++
					break rowLeftBelow
				}
			}
		}
		fallthrough
	case rightBelow:
	rowRightBelow:
		for i := 1; i < len(orig)-1; i++ {
			for j := col + 1; j < len(orig[row])-1; j++ {
				if orig[i][j] == "#" {
					adjacentCount++
					break rowRightBelow
				}
			}
		}
		fallthrough
	case leftAbove:
	rowLeftAbove:
		for i := 1; i < len(orig)-1; i++ {
			for j := 0; j < col; j++ {
				if orig[i][j] == "#" {
					adjacentCount++
					break rowLeftAbove
				}
			}
		}
		fallthrough
	case rightAbove:
	rowRightAbove:
		for i := 1; i < len(orig)-1; i++ {
			for j := col + 1; j < len(orig[row])-1; j++ {
				if orig[i][j] == "#" {
					adjacentCount++
					break rowRightAbove
				}
			}
		}
	}
	if adjacentCount == 0 {
		return "#"
	} else if adjacentCount >= 5 {
		return "L"
	}
	return orig[row][col]
}
