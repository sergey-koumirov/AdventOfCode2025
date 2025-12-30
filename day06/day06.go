package day06

import (
	"aoc2025/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Operation struct {
	Nums     []int
	Operator string
}

func Part1() {
	ops := load("./day06/input01.txt")

	sum := 0

	for _, op := range ops {
		if op.Operator == "+" {
			temp := 0
			for _, n := range op.Nums {
				temp += n
			}
			sum += temp
		} else if op.Operator == "*" {
			temp := 1
			for _, n := range op.Nums {
				temp *= n
			}
			sum += temp
		} else {
			fmt.Println("???", op)
		}
	}

	fmt.Println("Part1:", sum)
}

func Part2() {
	ops := load2("./day06/input01.txt")

	sum := 0

	for _, op := range ops {
		if op.Operator == "+" {
			temp := 0
			for _, n := range op.Nums {
				temp += n
			}
			sum += temp
		} else if op.Operator == "*" {
			temp := 1
			for _, n := range op.Nums {
				temp *= n
			}
			sum += temp
		} else {
			fmt.Println("???", op)
		}
	}

	fmt.Println("Part2:", sum)
}

func load2(fn string) []Operation {
	lines := common.LoadStrings(fn)

	nums := []int{}
	res := []Operation{}

	for i := len(lines[0]) - 1; i >= 0; i-- {
		temp := []byte{}

		for j := 0; j < len(lines)-1; j++ {
			temp = append(temp, lines[j][i])
		}

		str := strings.Trim(string(temp), " ")
		if str != "" {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)

			if lines[len(lines)-1][i] != ' ' {
				res = append(res, Operation{Operator: string(lines[len(lines)-1][i]), Nums: nums})
				nums = []int{}
			}
		}
	}

	return res
}

func load(fn string) []Operation {
	lines := common.LoadStrings(fn)

	re := regexp.MustCompile(`\s+`) // запятая или пробелы

	res := []Operation{}

	temp := make([][]string, len(lines))
	for i, line := range lines {
		temp[i] = re.Split(strings.Trim(line, " "), -1)
	}

	for i := 0; i < len(temp[0]); i++ {
		nums := []int{}
		for j := 0; j < len(temp)-1; j++ {
			num, _ := strconv.Atoi(temp[j][i])
			nums = append(nums, num)
		}

		res = append(res, Operation{Operator: temp[len(temp)-1][i], Nums: nums})
	}

	return res
}
