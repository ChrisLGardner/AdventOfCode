package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day12.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n")
	equivFacing := []byte{'E', 'N', 'W', 'S'}

	posX := 0
	posY := 0
	facing := 0

	for _, s := range split {
		n, err := strconv.Atoi(s[1:])
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}
		direction := s[0]
		switch s[0] {
		case 'R':
			facing -= n / 90
			facing += 4
			facing %= 4
		case 'L':
			facing += n / 90
			facing %= 4
		case 'F':
			direction = equivFacing[facing]
		}

		switch direction {
		case 'N':
			posY += n
		case 'E':
			posX += n
		case 'S':
			posY -= n
		case 'W':
			posX -= n
		}

	}

	manhattan := math.Abs(float64(posX)) + math.Abs(float64(posY))

	fmt.Println(manhattan)

	// Part 2

	wayX := 10
	wayY := 1

	posX = 0
	posY = 0

	for _, s := range split {
		n, err := strconv.Atoi(s[1:])
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}
		if n == 0 {
			continue
		}

		switch s[0] {
		case 'R':
			for i := 0; i < n/90; i++ {
				temp := wayY
				if wayX >= 0 {
					wayY = wayX * -1
				} else {
					wayY = int(math.Abs(float64(wayX)))
				}
				wayX = temp
			}
		case 'L':
			for i := n / 90; i > 0; i-- {
				temp := wayX
				if wayY >= 0 {
					wayX = wayY * -1
				} else {
					wayX = int(math.Abs(float64(wayY)))
				}
				wayY = temp
			}
		case 'F':
			posX += wayX * n
			posY += wayY * n
		case 'N':
			wayY += n
		case 'E':
			wayX += n
		case 'S':
			wayY -= n
		case 'W':
			wayX -= n
		}
	}

	manhattan = math.Abs(float64(posX)) + math.Abs(float64(posY))

	fmt.Println(manhattan)
}
