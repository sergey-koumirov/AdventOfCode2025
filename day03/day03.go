package day03

import (
	"aoc2025/common"
	"fmt"
)

func Part1() {
	rr := load("./day03/input01.txt")

	sum := 0

	for _, r := range rr {
		max1 := len(r) - 2

		for index := len(r) - 2; index >= 0; index-- {
			if r[index] >= r[max1] {
				max1 = index
			}
		}

		max2 := len(r) - 1

		for index := len(r) - 1; index > max1; index-- {
			if r[index] >= r[max2] {
				max2 = index
			}
		}

		sum += (int(r[max1])-48)*10 + int(r[max2]) - 48
	}

	fmt.Println("Part1:", sum)
}

func Part2() {
	rr := load("./day03/input01.txt")

	sum := 0

	for _, r := range rr {
		subSum := 0
		prevMax := -1

		for digit := 1; digit <= 12; digit++ {
			fromIndex := len(r) - 1 - (12 - digit)
			max := fromIndex
			// if prevMax > -1 {
			// 	fmt.Println(r[prevMax+1 : max+1])
			// } else {
			// 	fmt.Println(r[:max+1])
			// }

			for index := fromIndex - 1; index > prevMax; index-- {
				if r[index] >= r[max] {
					max = index
				}
			}

			// fmt.Println(digit, max, r[max:max+1])

			prevMax = max
			subSum += (int(r[max]) - 48) * common.Pow10(12-digit)
		}
		fmt.Println(subSum)

		sum += subSum
	}

	fmt.Println("Part2:", sum)
}

func load(fn string) []string {
	return common.LoadStrings(fn)
}
