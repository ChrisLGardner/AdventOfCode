package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day02.txt", "Relative file path to use as input.")

type password struct {
	min      int
	max      int
	char     string
	password string
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n")

	passwords := make([]password, len(split))
	pattern := regexp.MustCompile("(?P<min>\\d+)-(?P<max>\\d+) (?P<char>\\w): (?P<pass>\\w+)")
	names := pattern.SubexpNames()

	for i, s := range split {
		matchingStrings := pattern.FindAllStringSubmatch(s, -1)
		elements := map[string]string{}
		for j, match := range matchingStrings[0] {
			elements[names[j]] = match
		}
		min, err := strconv.Atoi(elements["min"])
		if err != nil {
			fmt.Printf("Regex went wrong here on %v: %v", elements["min"], err)
		}
		max, err := strconv.Atoi(elements["max"])
		if err != nil {
			fmt.Printf("Regex went wrong here on %v: %v", elements["min"], err)
		}
		passwords[i] = password{min, max, elements["char"], elements["pass"]}

	}

	count := 0

	for _, p := range passwords {

		if c := strings.Count(p.password, p.char); c >= p.min && c <= p.max {
			count++
		}
	}

	fmt.Printf("Found %d passwords matching the old requirements \n", count)

	countPart2 := 0

	for _, p := range passwords {
		pass1 := string(p.password[p.min-1])
		pass2 := string(p.password[p.max-1])
		if (pass1 == p.char && pass2 != p.char) || (pass1 != p.char && pass2 == p.char) {
			countPart2++
		}
	}
	fmt.Printf("Found %d passwords matching the new requirements\n", countPart2)
}
