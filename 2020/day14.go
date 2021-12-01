package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day14.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n")
	var mask string
	mem := make(map[int]int64)
	memPattern := regexp.MustCompile("mem\\[(?P<address>\\d+)\\] = (?P<value>\\d+)")
	names := memPattern.SubexpNames()

	for _, s := range split {
		if strings.Contains(s, "mask = ") {
			mask = strings.Split(s, " = ")[1]
			continue
		}

		matchingStrings := memPattern.FindAllStringSubmatch(s, -1)
		elements := map[string]string{}
		for j, match := range matchingStrings[0] {
			elements[names[j]] = match
		}
		address, err := strconv.Atoi(elements["address"])
		if err != nil {
			fmt.Printf("Regex went wrong here on %v: %v", elements["address"], err)
		}
		value, err := strconv.ParseInt(elements["value"], 10, 64)
		if err != nil {
			fmt.Printf("Regex went wrong here on %v: %v", elements["value"], err)
		}

		bits := fmt.Sprintf("%0*s", len(mask), strconv.FormatInt(value, 2))

		for i := 0; i < len(mask)-1; i++ {
			pos := len(mask) - (1 + i)
			char := string(mask[pos])
			if char == "X" {
				continue
			}
			bits = bits[:pos] + string(char) + bits[pos+1:]
		}

		value, err = strconv.ParseInt(bits, 2, 64)
		if err != nil {
			fmt.Printf("Regex went wrong here on %v: %v", elements["value"], err)
		}

		mem[address] = value
	}

	fmt.Println(mem)

	var total int64

	for _, v := range mem {
		total += v
	}

	fmt.Println(total)
}
