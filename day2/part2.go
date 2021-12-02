package main

import (
	"strconv"
	"strings"
)

func countPart2(in []string) (distance, depth int) {
	aim := 0
	for _, line := range in {
		items := strings.Split(line, " ")
		move, delta := items[0], items[1]
		num, _ := strconv.Atoi(delta)

		switch move {
		case "up":
			aim -= num
		case "down":
			aim += num
		case "forward":
			distance += num
			depth += aim * num
		}
	}

	return
}
