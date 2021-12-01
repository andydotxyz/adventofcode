package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func countPart1(in []int) int {
	increase := 0
	for i, depth := range in {
		if i == 0 { // no previous measurement
			continue
		}

		if depth > in[i-1] {
			increase++
		}
	}

	return increase
}

func isPart1() bool {
	return len(os.Args) == 1 || (os.Args[1] != "--part2" && os.Args[1] != "-2")
}

func loadInputs(path string) []int {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Could not read input file " + path)
	}

	lines := strings.Split(string(data), "\n")
	ret := make([]int, len(lines))
	for i, line := range lines {
		ret[i], _ = strconv.Atoi(line)
	}
	return ret
}

func main() {
	inputs := loadInputs("../day1/input.txt")
	count := 0
	if isPart1() {
		count = countPart1(inputs)
	} else {
		count = countPart2(inputs)
	}
	log.Printf("There are %d measurements larger than the previous",
		count)
}
