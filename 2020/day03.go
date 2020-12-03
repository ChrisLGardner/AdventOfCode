package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day03.txt", "Relative file path to use as input.")

type slope struct {
	right int
	down  int
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n")

	slopes := []slope{
		slope{1, 1},
		slope{3, 1},
		slope{5, 1},
		slope{7, 1},
		slope{1, 2},
	}
	totalTrees := make([]int, len(slopes))

	for i, s := range slopes {
		currpos := 0
		trees := 0
		for j, str := range split {
			if j%s.down != 0 {
				continue
			}

			if j == 0 {
				continue
			}

			currpos += s.right
			if string(rune(str[currpos%len(str)])) == "#" {
				trees++
			}
		}
		totalTrees[i] = trees
	}

	fmt.Println(totalTrees[0] * totalTrees[1] * totalTrees[2] * totalTrees[3] * totalTrees[4])

}
