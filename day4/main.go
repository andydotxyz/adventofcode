package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type board struct {
	nums   [5][5]int
	called [5][5]bool
	won bool
}

func (b *board) score() int {
	total := 0
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if !b.called[r][c] {
				total += b.nums[r][c]
			}
		}
	}

	return total
}

func countPart1(in []string) (*board, int) {
	calls := strings.Split(in[0], ",")
	boards := loadBoards(in[2:])
	for _, callStr := range calls {
		call, _ := strconv.Atoi(callStr)

		for _, b := range boards {
			for r := 0; r < 5; r++ {
				for c := 0; c < 5; c++ {
					if b.nums[r][c] == call {
						b.called[r][c] = true
					}
				}
			}

			for r := 0; r < 5; r++ {
				if b.called[r][0] && b.called[r][1] && b.called[r][2] && b.called[r][3] && b.called[r][4] {
					return b, b.score() * call
				}
				if b.called[0][r] && b.called[1][r] && b.called[2][r] && b.called[3][r] && b.called[4][r] {
					return b, b.score() * call
				}
			}
		}
	}

	return nil, 0
}

func isPart1() bool {
	return len(os.Args) == 1 || (os.Args[1] != "--part2" && os.Args[1] != "-2")
}

func loadBoards(in []string) []*board {
	var boards = []*board{}
	var current *board
	for i, line := range in {
		if i%6 == 0 {
			current = &board{}
			boards = append(boards, current)
		} else if i%6 == 5 {
			continue
		}

		var a, b, c, d, e int
		_, _ = fmt.Sscanf(line, "%d %d %d %d %d", &a, &b, &c, &d, &e)
		current.nums[i%6] = [5]int{a, b, c, d, e}
	}

	return boards
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
		board, score := countPart1(inputs)
		log.Printf("Winning board score %d, for state %#v",
			score, board)
	} else {
		board, score := countPart2(inputs)
		log.Printf("Losing board score %d, for state %#v",
			score, board)
	}
}
