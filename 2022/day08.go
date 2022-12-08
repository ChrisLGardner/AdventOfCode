package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day08.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n")
	grid := make([][]int, len(split))
	for i, s := range split {
		arr := make([]int, len(s))
		for k, v := range s {
			n, err := strconv.Atoi(string(v))
			if err != nil {
				fmt.Printf("Failed to parse %s\n", s)
				break
			}
			if n < 0 {
				fmt.Printf("Optimization invariant broken: %d <= 0 \n", n)
				break
			}
			arr[k] = n
		}
		grid[i] = arr
	}

	//fmt.Println(grid)

	count := (len(grid)+len(grid[0]))*2 - 4
	for i, row := range grid {
		if i == 0 || i == len(grid)-1 {
			continue
		}

		for k := range row {
			if k == 0 || k == len(row)-1 {
				continue
			}

			if isVisible(i, k, grid) {
				count += 1
			}
		}
	}

	fmt.Printf("Visible trees: %d\n", count)

	scenicScore := 0
	for i, row := range grid {
		if i == 0 || i == len(grid)-1 {
			continue
		}

		for k := range row {
			if k == 0 || k == len(row)-1 {
				continue
			}

			if score := getScenicScore(i, k, grid); score > scenicScore {
				scenicScore = score
			}
		}
	}
	fmt.Printf("Scenic Score: %d\n", scenicScore)
}

func isVisible(row int, col int, grid [][]int) bool {

	left := true
	right := true
	top := true
	bottom := true
	// check left -> right
	for i := 0; i < col; i++ {
		if grid[row][i] >= grid[row][col] {
			left = false
			break
		}
	}
	// check right -> left
	for i := len(grid[row]) - 1; i > col; i-- {
		if grid[row][i] >= grid[row][col] {
			right = false
			break
		}
	}
	// check top -> bottom
	for i := 0; i < row; i++ {
		if grid[i][col] >= grid[row][col] {
			top = false
			break
		}
	}
	// check bottom -> top
	for i := len(grid) - 1; i > row; i-- {
		if grid[i][col] >= grid[row][col] {
			bottom = false
			break
		}
	}

	if left || right || top || bottom {
		return true
	}
	//fmt.Printf("%v:%v is not visible\n", row, col)
	return false
}

func getScenicScore(row, col int, grid [][]int) int {

	left := 0
	right := 0
	top := 0
	bottom := 0
	// check tree -> left
	for i := col; i >= 0; i-- {
		if (i - 1) < 0 {
			break
		}
		left++
		if grid[row][i-1] >= grid[row][col] {
			break
		}
	}
	// check tree -> right
	for i := col; i < len(grid[row]); i++ {
		if (i + 1) >= len(grid[row]) {
			break
		}
		right++
		if grid[row][i+1] >= grid[row][col] {
			break
		}
	}
	// check tree -> top
	for i := row; i >= 0; i-- {
		if (i - 1) < 0 {
			break
		}
		top++
		if grid[i-1][col] >= grid[row][col] {
			break
		}
	}
	// check tree -> bottom
	for i := row; i < len(grid); i++ {
		if (i + 1) >= len(grid) {
			break
		}
		bottom++
		if grid[i+1][col] >= grid[row][col] {
			break
		}
	}
	return left * right * top * bottom
}
