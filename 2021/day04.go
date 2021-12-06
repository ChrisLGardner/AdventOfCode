package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day04.txt", "Relative file path to use as input.")
var part2 = flag.Bool("part2", false, "Run logic for part 2 of puzzle")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n\n")

	sequence := strings.Split(split[0], ",")
	split = append([]string(nil), split[1:]...)
	boards := make([][][]string, len(split))
	for i, s := range split {

		rawBoard := strings.Split(s, "\n")
		board := make([][]string, 5)
		for j, row := range rawBoard {
			if row == "" {
				continue
			}
			cols := strings.Split(row, " ")
			row := make([]string, 5)
			k := 0
			for _, col := range cols {
				n := strings.Trim(col, " ")
				if n == "" {
					continue
				}
				row[k] = n
				k++
			}
			board[j] = row
		}

		boards[i] = board
	}

	var winner [][]string
	var winningNumber string

outer:
	for _, num := range sequence {

		for _, board := range boards {
			if CheckBoardForNumber(board, num) {
				if CheckBoardForWinning(board) {
					winner = board
					winningNumber = num
					break outer
				}
			}
		}
	}

	fmt.Printf("Winning Number: %d\n", GetWinningTotal(winner, winningNumber))

}

func CheckBoardForNumber(board [][]string, number string) bool {

	for i, row := range board {
		for k, col := range row {
			if col == number {
				board[i][k] = "x"
				return true
			}
		}
	}
	return false
}

func CheckBoardForWinning(board [][]string) bool {

	winning := false
	for _, row := range board {
		count := 0
		for _, col := range row {
			if col == "x" {
				count++
			}
		}
		if count == len(row) {
			winning = true
		}
	}

	for i := range board[0] {
		count := 0
		for j := range board {
			if board[j][i] == "x" {
				count++
			}
		}
		if count == len(board) {
			winning = true
		}
	}

	return winning
}

func GetWinningTotal(board [][]string, number string) int {

	total := 0
	lastNumber, _ := strconv.Atoi(number)

	for _, row := range board {
		for _, col := range row {
			n, _ := strconv.Atoi(col)
			total += n
		}
	}

	return total * lastNumber
}
