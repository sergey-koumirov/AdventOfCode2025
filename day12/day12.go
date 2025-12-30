package day12

import (
	"aoc2025/common"
	"fmt"
	"strconv"
	"strings"
)

type Figure struct {
	Shape []string
	Index int
	Mass  int
}

type Task struct {
	Width, Height int
	Counts        []int
}

func Part1() {
	ff, tt := load("./day12/input01.txt")

	fmt.Println(ff)

	res := 0

	for _, t := range tt {
		fullMass := 0
		allCnt := 0
		for i, cnt := range t.Counts {
			fullMass += ff[i].Mass * cnt
			allCnt += cnt
		}

		if allCnt*9 <= (t.Height/3)*3*(t.Width/3)*3 {
			// fmt.Println(t, "free")
			res++
		} else if fullMass <= t.Height*t.Width {
			fmt.Println(t, fullMass, t.Height*t.Width)
		} else {
			// fmt.Println(t, "Not possible")
		}

	}

	fmt.Println("Part1:", res)
}

func Part2() {
	// points := load("./day09/input01.txt")
	fmt.Println("Part2:")
}

func load(fn string) ([]Figure, []Task) {
	lines := common.LoadStrings(fn)

	figures := []Figure{}
	tasks := []Task{}

	mode := "F"

	acc := []string{}
	index := 0

	for _, line := range lines {
		if strings.Contains(line, "x") {
			mode = "T"
		}

		if mode == "F" && !strings.Contains(line, ":") {
			acc = append(acc, line)
			if line == "" {
				mass := 0
				for _, l := range acc {
					for _, c := range l {
						if c == '#' {
							mass++
						}
					}
				}

				figures = append(figures, Figure{Index: index, Shape: acc, Mass: mass})
				acc = []string{}
				index++
			}
		}

		if mode == "T" {
			temp := Task{Counts: []int{}}
			parts1 := strings.Split(line, ": ")

			parts2 := strings.Split(parts1[0], "x")

			temp.Width, _ = strconv.Atoi(parts2[0])
			temp.Height, _ = strconv.Atoi(parts2[1])

			parts3 := strings.Split(parts1[1], " ")

			for _, s := range parts3 {
				num, _ := strconv.Atoi(s)
				temp.Counts = append(temp.Counts, num)
			}

			tasks = append(tasks, temp)
		}
	}

	return figures, tasks
}
