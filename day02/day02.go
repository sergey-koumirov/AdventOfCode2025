package day02

import (
	"aoc2025/common"
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	First, Last string
}

type Ranges []Range

func Part1() {
	rr := load("./day02/input01.txt")

	sum := 0

	for _, r := range rr {
		len1 := len(r.First)
		len2 := len(r.Last)

		for lenX := len1; lenX <= len2; lenX++ {
			if lenX%2 == 0 {
				var minRange, maxRange int

				if len(r.First) < lenX {
					minRange = common.Pow10(lenX/2 - 1)
				} else {
					num1, _ := strconv.Atoi(r.First[:lenX/2])
					num2, _ := strconv.Atoi(r.First[lenX/2:])
					if num1 >= num2 {
						minRange = num1
					} else {
						minRange = num1 + 1
					}
				}

				if len(r.Last) > lenX {
					maxRange = common.Pow10(lenX/2) - 1
				} else {
					num1, _ := strconv.Atoi(r.Last[:lenX/2])
					num2, _ := strconv.Atoi(r.Last[lenX/2:])

					if num1 <= num2 {
						maxRange = num1
					} else {
						maxRange = num1 - 1
					}
				}

				p := common.Pow10(lenX / 2)
				for x := minRange; x <= maxRange; x++ {
					sum += x*p + x
				}
			}
		}
	}

	fmt.Println("Part1:", sum)
}

func Part2() {
	rr := load("./day02/input01.txt")

	sum := 0

	used := map[int]bool{}

	for _, r := range rr {
		for digits := 1; digits <= len(r.Last)/2; digits++ {
			if len(r.First)%digits == 0 || len(r.Last)%digits == 0 {
				var lenX int
				if len(r.First)%digits == 0 {
					lenX = len(r.First) / digits
				} else {
					lenX = len(r.Last) / digits
				}

				var normFirst, normLast string

				if len(r.First)%digits != 0 {
					normFirst = strings.Repeat("1"+strings.Repeat("0", digits-1), lenX)
				} else {
					normFirst = r.First
				}

				if len(r.Last)%digits != 0 {
					normLast = strings.Repeat(strings.Repeat("9", digits), lenX)
				} else {
					normLast = r.Last
				}

				if digits == 1 && len(normFirst) < len(normLast) {
					sum += analyze(normFirst, strings.Repeat("9", len(normFirst)), digits, used)
					sum += analyze(strings.Repeat("1"+strings.Repeat("0", digits-1), len(normLast)), normLast, digits, used)
				} else {
					sum += analyze(normFirst, normLast, digits, used)
				}
			}
		}
	}

	fmt.Println("Part2:", sum)
}

func analyze(normFirst, normLast string, digits int, used map[int]bool) int {
	if len(normFirst) == 1 {
		return 0
	}

	var minRange, maxRange int

	groups := len(normFirst) / digits

	for n := 1; n <= groups; n++ {
		firstNum, _ := strconv.Atoi(normFirst[(n-1)*digits : (n)*digits])
		if n == 1 {
			minRange = firstNum
		} else {
			if minRange > firstNum {
				break
			} else if minRange < firstNum {
				minRange += 1
				break
			}
		}
	}

	for n := 1; n <= groups; n++ {
		lastNum, _ := strconv.Atoi(normLast[(n-1)*digits : (n)*digits])
		if n == 1 {
			maxRange = lastNum
		} else {
			if maxRange < lastNum {
				break
			} else if maxRange > lastNum {
				maxRange = maxRange - 1
				break
			}
		}
	}

	subSum := 0
	p := common.Pow10(digits)
	for x := minRange; x <= maxRange; x++ {
		wrongN := repeatInt(x, groups, p)

		if !used[wrongN] {
			subSum += wrongN
			used[wrongN] = true
		}
	}

	return subSum
}

func repeatInt(x, n, p int) int {
	res := x
	for i := 1; i < n; i++ {
		res = res*p + x
	}
	return res
}

func load(fn string) Ranges {
	res := Ranges{}

	lines := common.LoadStrings(fn)

	parts := strings.Split(lines[0], ",")

	for _, part := range parts {
		subparts := strings.Split(part, "-")
		res = append(res, Range{First: subparts[0], Last: subparts[1]})
	}

	return res
}
