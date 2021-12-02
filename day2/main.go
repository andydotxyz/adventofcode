package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func countPart1(in []string) (distance, depth int) {
	for _, line := range in {
		items := strings.Split(line, " ")
		move, delta := items[0], items[1]
		num, _ := strconv.Atoi(delta)

		switch move {
		case "up":
			depth -= num
		case "down":
			depth += num
		case "forward":
			distance += num
		}
	}

	return
}

func isPart1() bool {
	return len(os.Args) == 1 || (os.Args[1] != "--part2" && os.Args[1] != "-2")
}

func loadInputs(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Could not read input file " + path)
	}

	lines := strings.Split(string(data), "\n")
	return lines
}

func main() {
	inputs := loadInputs("./input.txt")
	distance, depth := 0, 0
	if isPart1() {
		distance, depth = countPart1(inputs)
	} else {
		distance, depth = countPart2(inputs)
	}
	log.Printf("Travelled %d, depth is %d - multiple is %d",
		distance, depth, distance*depth)
}
