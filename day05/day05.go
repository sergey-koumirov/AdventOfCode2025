package day05

import (
	"aoc2025/common"
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	Left, Right int
}

func Part1() {
	rr, nn := load("./day05/input01.txt")

	cnt := 0

	for _, n := range nn {
		for _, r := range rr {
			if r.Left <= n && n <= r.Right {
				cnt++
				break
			}
		}
	}

	fmt.Println("Part1:", cnt)
}

func Part2() {
	rr, _ := load("./day05/input01.txt")

	for {
		changes := false

		lastIndex := len(rr) - 1

		for index1 := 0; index1 < lastIndex; index1++ {
			for index2 := index1 + 1; index2 <= lastIndex; index2++ {
				r1 := rr[index1]
				r2 := rr[index2]

				if r1.Right >= r2.Left && r1.Left <= r2.Right || r2.Right >= r1.Left && r2.Left <= r1.Right {
					rr[index1].Left = common.Min(r1.Left, r2.Left)
					rr[index1].Right = common.Max(r1.Right, r2.Right)

					rr = append(rr[:index2], rr[index2+1:]...)
					changes = true
					break
				}
			}

			if changes {
				break
			}
		}

		if !changes {
			break
		}
	}

	cnt := 0

	for _, r := range rr {
		cnt += r.Right - r.Left + 1
	}

	fmt.Println("Part2:", cnt)
}

func load(fn string) ([]Range, []int) {
	lines := common.LoadStrings(fn)

	firstPart := true

	ranges := []Range{}
	numbers := []int{}

	for _, line := range lines {
		if line == "" {
			firstPart = false
		} else if firstPart {
			parts := strings.Split(line, "-")

			left, _ := strconv.Atoi(parts[0])
			right, _ := strconv.Atoi(parts[1])

			ranges = append(ranges, Range{Left: left, Right: right})
		} else {
			num, _ := strconv.Atoi(line)
			numbers = append(numbers, num)
		}
	}

	return ranges, numbers
}
