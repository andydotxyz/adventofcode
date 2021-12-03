package main

import "strconv"

func countPart2(in []string) (int, int) {
	return filter(in, in, 0)
}

func filter(o2in, co2in []string, level int) (int, int) {
	var o2out, co2out []string
	if len(o2in) == 1 {
		o2out = o2in
	} else {
		ones := 0
		total := 0
		for _, line := range o2in {
			total++

			if line[level] == '1' {
				ones++
			}
		}

		for _, line := range o2in {
			if ones*2 >= total {
				if line[level] == '1' {
					o2out = append(o2out, line)
				}
			} else {
				if line[level] == '0' {
					o2out = append(o2out, line)
				}
			}
		}
	}

	if len(co2in) == 1 {
		co2out = co2in
	} else {
		ones := 0
		total := 0
		for _, line := range co2in {
			total++

			if line[level] == '1' {
				ones++
			}
		}

		for _, line := range co2in {
			if ones*2 >= total {
				if line[level] == '0' {
					co2out = append(co2out, line)
				}
			} else {
				if line[level] == '1' {
					co2out = append(co2out, line)
				}
			}
		}
	}

	if level == 11 {
		gen, _ := strconv.ParseInt(o2out[0], 2, 64)
		scrub, _ := strconv.ParseInt(co2out[0], 2, 64)

		return int(gen), int(scrub)
	} else {
		return filter(o2out, co2out, level+1)
	}
}
