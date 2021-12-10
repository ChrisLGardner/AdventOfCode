package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day09.txt", "Relative file path to use as input.")

type coord struct {
	row int
	col int
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n")
	seen := make([][]int, len(split))
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
		seen[i] = arr
	}

	lowPointsTotal := 0
	lowPoints := []coord{}

	for i, row := range seen {
		for j, col := range row {
			if j > 0 && row[j-1] <= col {
				continue
			} else if i > 0 && seen[i-1][j] <= col {
				continue
			} else if i < (len(seen)-1) && seen[i+1][j] <= col {
				continue
			} else if j < (len(row)-1) && row[j+1] <= col {
				continue
			} else {
				lowPoints = append(lowPoints, coord{i, j})
				lowPointsTotal += col + 1
			}
		}
	}

	fmt.Printf("Total of all low points: %d \n", lowPointsTotal)

	basinSize := 1
	basins := make([]int, len(lowPoints))
	for i, loc := range lowPoints {
		basin, _ := FindHigherPoints(seen, loc.row, loc.col, map[coord]bool{})
		basins[i] = basin
	}

	sort.Ints(basins)
	for i := len(basins) - 1; i >= len(basins)-3; i-- {
		basinSize *= basins[i]
	}
	fmt.Printf("Total basin size: %d\n", basinSize)
}

func FindHigherPoints(grid [][]int, row, col int, visited map[coord]bool) (int, map[coord]bool) {

	visited[coord{row, col}] = true
	count := 1
	// check above
	if row > 0 {
		if !visited[coord{row - 1, col}] && grid[row-1][col] != 9 && grid[row][col] < grid[row-1][col] {
			points, locs := FindHigherPoints(grid, row-1, col, visited)
			count += points
			visited = locs
		}
	}
	// check below
	if row < (len(grid) - 1) {
		if !visited[coord{row + 1, col}] && grid[row+1][col] != 9 && grid[row][col] < grid[row+1][col] {
			points, locs := FindHigherPoints(grid, row+1, col, visited)
			count += points
			visited = locs
		}
	}
	// check left
	if col > 0 {
		if !visited[coord{row, col - 1}] && grid[row][col-1] != 9 && grid[row][col] < grid[row][col-1] {
			points, locs := FindHigherPoints(grid, row, col-1, visited)
			count += points
			visited = locs
		}
	}
	// check right
	if col < (len(grid[0]) - 1) {
		if !visited[coord{row, col + 1}] && grid[row][col+1] != 9 && grid[row][col] < grid[row][col+1] {
			points, locs := FindHigherPoints(grid, row, col+1, visited)
			count += points
			visited = locs
		}
	}

	return count, visited

}

func checkVisited(visited []coord, row, col int) bool {
	for _, visit := range visited {
		if visit.row == row && visit.col == col {
			return true
		}
	}
	return false
}
