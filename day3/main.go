package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func countPart1(in []string) (gamma, epsilon int) {
	ones := make([]int, 12)
	total := 0
	for _, line := range in {
		total++

		for i, r := range []rune(line) {
			if r == '1' {
				ones[i]++
			}
		}
	}

	pow := 1
	for i := 11; i >= 0; i-- {
		if ones[i] > total/2 {
			gamma += pow
		} else {
			epsilon += pow
		}
		pow *= 2
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
	if isPart1() {
		gamma, epsilon := countPart1(inputs)
		log.Printf("Gamma level %d, epsilon %d - power consumption %d",
			gamma, epsilon, gamma*epsilon)
	} else {
		gen, scrub := countPart2(inputs)
		log.Printf("O2 generator level %d, CO2 scrubber level %d - life support %d",
			gen, scrub, gen*scrub)
	}
}
