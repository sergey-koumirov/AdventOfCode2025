package day07

import (
	"aoc2025/common"
	"bytes"
	"fmt"
	"strings"
)

type Operation struct {
	Nums     []int
	Operator string
}

func Part1() {
	lines := load("./day07/input01.txt")

	cnt := 0

	var state []byte

	for lineIndex, line := range lines {
		newState := bytes.Repeat([]byte{' '}, len(line))

		for index, char := range line {
			if char == 'S' {
				newState[index] = '|'
			}
			if lineIndex > 0 && char == '^' && state[index] == '|' {
				cnt++
				if index > 0 {
					newState[index-1] = '|'
				}
				if index < len(line)-1 {
					newState[index+1] = '|'
				}
			} else if lineIndex > 0 && char != '^' && state[index] == '|' {
				newState[index] = '|'
			}
		}

		state = newState
	}

	fmt.Println("Part1:", cnt)
}

func Part2() {
	lines := load("./day07/input01.txt")

	sIndex := strings.Index(lines[0], "S")

	cntMem := map[int]int{}

	cnt := deepGetCnt(lines, cntMem, 0, sIndex)

	// common.PrintStrings(lines)

	fmt.Println("Part2:", cnt)
}

func deepGetCnt(lines []string, cntMem map[int]int, lineIndex, charIndex int) int {

	cnt, ex := cntMem[1000*lineIndex+charIndex]

	if ex {
		return cnt
	}

	if lineIndex == len(lines) {
		return 1
	}

	char := lines[lineIndex][charIndex]
	if char == 'S' || char == '.' {
		nextCnt := deepGetCnt(lines, cntMem, lineIndex+1, charIndex)
		cntMem[1000*lineIndex+charIndex] = nextCnt
		return nextCnt
	}

	if char == '^' {
		nextCntL := deepGetCnt(lines, cntMem, lineIndex, charIndex-1)
		nextCntR := deepGetCnt(lines, cntMem, lineIndex, charIndex+1)
		cntMem[1000*lineIndex+charIndex] = nextCntL + nextCntR
		return nextCntL + nextCntR
	}

	fmt.Println(char, string(char))
	panic("???")
}

func load(fn string) []string {
	lines := common.LoadStrings(fn)

	return lines
}
