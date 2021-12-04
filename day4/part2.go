package main

import (
	"strconv"
	"strings"
)

func countPart2(in []string) (*board, int) {
	calls := strings.Split(in[0], ",")
	boards := loadBoards(in[2:])
var last *board
	lastCall := 0
	for _, callStr := range calls {
		call, _ := strconv.Atoi(callStr)

		for _, b := range boards {
			if b.won {
				continue
			}
			for r := 0; r < 5; r++ {
				for c := 0; c < 5; c++ {
					if b.nums[r][c] == call {
						b.called[r][c] = true
					}
				}
			}

			for r := 0; r < 5; r++ {
				if b.called[r][0] && b.called[r][1] && b.called[r][2] && b.called[r][3] && b.called[r][4] {
					b.won = true
					last = b
					lastCall = call
				}
				if b.called[0][r] && b.called[1][r] && b.called[2][r] && b.called[3][r] && b.called[4][r] {
					b.won = true
					last = b
					lastCall = call
				}
			}
		}
	}

	return last, last.score()*lastCall
}
