package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day07.txt", "Relative file path to use as input.")

type bagContents struct {
	count  int
	colour string
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n")
	contains := make(map[string][]bagContents)
	containedBy := make(map[string][]string)

	outerBagPattern := regexp.MustCompile("(\\w+ \\w+) bags contain")
	innerBagPattern := regexp.MustCompile("contain (.+)")
	bagTypePattern := regexp.MustCompile("(?P<count>\\d) (?P<colour>\\w+ \\w+)")

	for _, s := range split {
		outerBagColour := outerBagPattern.FindAllStringSubmatch(s, -1)[0][1]

		innerBags := innerBagPattern.FindAllStringSubmatch(s, -1)[0][1]

		if innerBags == "no other bags." {
			contains[outerBagColour] = nil
			continue
		}

		for _, v := range strings.Split(innerBags, ",") {
			bagRegex := bagTypePattern.FindAllStringSubmatch(v, -1)
			bagCount, err := strconv.Atoi(bagRegex[0][1])
			if err != nil {
				fmt.Printf("Failed to get count of bags from %s", bagRegex[0][0])
			}
			bag := bagContents{bagCount, bagRegex[0][2]}
			contains[outerBagColour] = append(contains[outerBagColour], bag)
			containedBy[bag.colour] = append(containedBy[bag.colour], outerBagColour)
		}
	}

	newlyFound := []string{"shiny gold"}
	found := map[string]bool{
		"shiny gold": true,
	}
	shinyGoldParents := 0
	for {
		nextCycle := make([]string, 0)
		for _, v := range newlyFound {
			validParents := containedBy[v]
			for _, p := range validParents {
				if !found[p] {
					shinyGoldParents++
					found[p] = true
					nextCycle = append(nextCycle, p)
				}
			}
		}

		if len(nextCycle) == 0 {
			break
		}
		newlyFound = nextCycle
	}
	fmt.Println(shinyGoldParents)

	newlyFoundBags := map[string]int{
		"shiny gold": 1,
	}

	bagsInside := make(map[string]int)

	for {
		nextCycle := make(map[string]int)

		for k, v := range newlyFoundBags {
			bagsContained := contains[k]
			for _, bag := range bagsContained {
				bagsInside[bag.colour] += bag.count * v
				nextCycle[bag.colour] += bag.count * v

			}
		}

		if len(nextCycle) == 0 {
			break
		}
		newlyFoundBags = nextCycle
	}

	children := 0
	for _, v := range bagsInside {
		children += v
	}

	fmt.Println(children)
}
