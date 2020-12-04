package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day04.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n\n")

	byrPattern := regexp.MustCompile("byr:(19[2-9]\\d|200[12])\\b")
	iyrPattern := regexp.MustCompile("iyr:(201\\d|2020)\\b")
	eyrPattern := regexp.MustCompile("eyr:(202\\d|2030)\\b")
	hgtPattern := regexp.MustCompile("hgt:((1[5-8]\\d|19[0-3])cm|(59|6\\d|7[0-6])in)\\b")
	hclPattern := regexp.MustCompile("hcl:#[0-9a-f]{6}\\b")
	eclPattern := regexp.MustCompile("ecl:(amb|blu|brn|gry|grn|hzl|oth){1}\\b")
	pidPattern := regexp.MustCompile("pid:\\d{9}\\s?")

	count := 0
	for _, s := range split {
		if byrPattern.MatchString(s) && iyrPattern.MatchString(s) && eyrPattern.MatchString(s) && hgtPattern.MatchString(s) && hclPattern.MatchString(s) && eclPattern.MatchString(s) && pidPattern.MatchString(s) {
			count++
		}
	}

	fmt.Println(count)

}
