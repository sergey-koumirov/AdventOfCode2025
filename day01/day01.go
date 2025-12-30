package day01

import (
	"aoc2025/common"
	"fmt"
	"strconv"
)

type Action struct {
	Dir string
	Val int
}

type Actions []Action

const N = 100
const START = 50

func Part1() {
	aa := load("./day01/input01.txt")
	init := START
	cnt := 0

	for _, a := range aa {
		rest := a.Val % N

		if a.Dir == "L" {
			if init < rest {
				init = init + N - rest
			} else {
				init = init - rest
			}
		} else {
			init = (init + rest) % N
		}

		if init == 0 {
			cnt += 1
		}
	}

	fmt.Println("Part1:", cnt)
}

func Part2() {
	aa := load("./day01/input01.txt")
	init := START
	cnt := 0

	for _, a := range aa {
		rest := a.Val % N
		taps := a.Val / N

		if a.Dir == "L" {
			if init < rest {
				if init != 0 {
					taps += 1
				}
				init = init + N - rest
			} else {
				init = init - rest
				if init == 0 {
					taps += 1
				}
			}
		} else {
			if init+rest >= N {
				taps += 1
			}
			init = (init + rest) % N
		}

		cnt += taps
	}

	fmt.Println("Part2:", cnt)
}

func load(fn string) Actions {
	res := Actions{}

	lines := common.LoadStrings(fn)

	for _, line := range lines {
		valStr := line[1:]
		val, _ := strconv.Atoi(valStr)

		res = append(res, Action{Dir: line[0:1], Val: val})
	}

	return res
}
