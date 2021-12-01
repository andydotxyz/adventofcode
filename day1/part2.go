package main

func countPart2(in []int) int {
	increase := 0
	for i, depth := range in {
		if i <= 2 { // no previous measurement
			continue
		}

		if depth+in[i-1]+in[i-2] > in[i-1]+in[i-2]+in[i-3] {
			increase++
		}
	}

	return increase
}
