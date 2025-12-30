package day04

import (
	"aoc2025/common"
	"fmt"
)

func Part1() {
	rr := load("./day04/input01.txt")

	cnt := 0

	for r := 0; r < len(rr); r++ {
		for c := 0; c < len(rr[0]); c++ {
			if rr[r][c] == '@' && countAround(rr, r, c) < 4 {
				cnt++
			}
		}
	}

	fmt.Println("Part1:", cnt)
}

func Part2() {
	rr := load("./day04/input01.txt")

	cnt := 0

	for {
		removed := 0

		copied := make([]string, len(rr))
		copy(copied, rr)

		for r := 0; r < len(rr); r++ {
			for c := 0; c < len(rr[0]); c++ {
				if rr[r][c] == '@' && countAround(rr, r, c) < 4 {
					removed++

					bs := []byte(copied[r])
					bs[c] = '.'
					copied[r] = string(bs)
				}
			}
		}

		cnt += removed
		if removed == 0 {
			break
		}

		rr = copied
	}

	fmt.Println("Part2:", cnt)
}

func countAround(f []string, r, c int) int {
	res := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			rowOk := r+i >= 0 && r+i < len(f)
			colOk := c+j >= 0 && c+j < len(f[0])
			if rowOk && colOk && (i != 0 || j != 0) && f[r+i][c+j] == '@' {
				res++
			}
		}
	}

	return res
}

func load(fn string) []string {
	return common.LoadStrings(fn)
}
