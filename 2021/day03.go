package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input/day03.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(strings.TrimRight(contents, "\n"), "\n")

	gammaBinary := ""
	epsilonBinary := ""

	for i := 0; i < len(split[0]); i++ {
		ones := 0
		zeroes := 0
		for _, v := range split {
			if string(v[i]) == "1" {
				ones++
			} else {
				zeroes++
			}
		}

		if ones > zeroes {
			gammaBinary += "1"
			epsilonBinary += "0"
		} else {
			gammaBinary += "0"
			epsilonBinary += "1"
		}
	}

	gamma, _ := strconv.ParseInt(gammaBinary, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonBinary, 2, 64)
	fmt.Printf("Power is %d * %d = %d\n", gamma, epsilon, gamma*epsilon)

	oxygenSlice := append([]string(nil), split...)
	co2Slice := append([]string(nil), split...)

	// Oxygen
	var oxygen string
	for i := 0; i < len(split[0]); i++ {
		commonBit := GetMostCommonBit(oxygenSlice, i, "1")
		var tempSlice []string

		if len(oxygenSlice) == 1 {
			oxygen = oxygenSlice[0]
		}
		for _, s := range oxygenSlice {
			if rune(s[i]) == rune(commonBit[0]) {
				tempSlice = append(tempSlice, s)
			}
		}
		oxygenSlice = append([]string(nil), tempSlice...)
	}

	if len(oxygenSlice) == 1 {
		oxygen = oxygenSlice[0]
	}

	// CO2
	var co2 string
	for i := 0; i < len(split[0]); i++ {
		commonBit := GetLeastCommonBit(co2Slice, i, "0")
		var tempSlice []string

		if len(co2Slice) == 1 {
			co2 = co2Slice[0]
		}
		for _, s := range co2Slice {
			if rune(s[i]) == rune(commonBit[0]) {
				tempSlice = append(tempSlice, s)
			}
		}

		co2Slice = append([]string(nil), tempSlice...)
	}

	if len(co2Slice) == 1 {
		co2 = co2Slice[0]
	}

	oxygenVal, _ := strconv.ParseInt(oxygen, 2, 64)
	co2Val, _ := strconv.ParseInt(co2, 2, 64)
	fmt.Printf("Life Support is %d * %d = %d\n", oxygenVal, co2Val, oxygenVal*co2Val)
}

func GetMostCommonBit(slice []string, position int, tiedBit string) string {

	ones := 0
	zeroes := 0
	for _, v := range slice {
		if string(v[position]) == "1" {
			ones++
		} else {
			zeroes++
		}
	}

	if ones == zeroes {
		return tiedBit
	} else if ones > zeroes {
		return "1"
	} else {
		return "0"
	}
}

func GetLeastCommonBit(slice []string, position int, tiedBit string) string {

	ones := 0
	zeroes := 0
	for _, v := range slice {
		if string(v[position]) == "1" {
			ones++
		} else {
			zeroes++
		}
	}

	if ones == zeroes {
		return tiedBit
	} else if ones > zeroes {
		return "0"
	} else {
		return "1"
	}
}
