package day11

import (
	"aoc2025/common"
	"fmt"
	"strings"
)

func Part1() {
	rr, _ := load("./day11/input01.txt")

	cnt := 0

	deepCntYou(rr, "you", &cnt)

	fmt.Println("Part1:", cnt)
}

func deepCntYou(rr map[string][]string, current string, cnt *int) {
	for _, next := range rr[current] {
		if next == "out" {
			*cnt += 1
		} else {
			deepCntYou(rr, next, cnt)
		}
	}
}

func Part2() {
	_, rrBack := load("./day11/input01.txt")

	cntSrvDac := dfs(rrBack, "svr", "dac", "fft", map[string]bool{"svr": true}, map[string]int{"svr": 1})

	cntSrvFft := dfs(rrBack, "svr", "fft", "dac", map[string]bool{"svr": true}, map[string]int{"svr": 1})

	cntDacFft := dfs(rrBack, "dac", "fft", "svr", map[string]bool{"dac": true}, map[string]int{"dac": 1})

	cntFftDac := dfs(rrBack, "fft", "dac", "svr", map[string]bool{"fft": true}, map[string]int{"fft": 1})

	cntFftOut := dfs(rrBack, "fft", "out", "dac", map[string]bool{"fft": true}, map[string]int{"fft": 1})

	cntDacOut := dfs(rrBack, "dac", "out", "fft", map[string]bool{"dac": true}, map[string]int{"dac": 1})

	cnt := cntSrvDac*cntDacFft*cntFftOut + cntSrvFft*cntFftDac*cntDacOut

	fmt.Println("Part2:", cnt)
}

func dfs(backwardTree map[string][]string, fromNode, toNode, ignore string, visited map[string]bool, cnts map[string]int) int {
	_, ex := visited[toNode]
	if ex {
		return cnts[toNode]
	}

	temp := 0
	for _, n := range backwardTree[toNode] {
		if n != ignore {
			temp += dfs(backwardTree, fromNode, n, ignore, visited, cnts)
		}
	}
	visited[toNode] = true
	cnts[toNode] = temp
	return temp
}

func possibleNodes(nodeName string, backwardTree map[string][]string) map[string]bool {
	res := map[string]bool{}

	newNodes := []string{nodeName}
	res[nodeName] = true

	for {
		temp := []string{}
		for _, n := range newNodes {
			for _, bNode := range backwardTree[n] {
				_, ex := res[bNode]
				if !ex {
					temp = append(temp, bNode)
					res[bNode] = true
				}
			}
		}

		if len(temp) == 0 {
			break
		}
		newNodes = temp
	}

	return res
}

func deepCntXY(rr map[string][]string, current, finish, ignore string, possible map[string]bool, cnt *int, level int) {
	// if *cnt%100000 == 0 {
	// 	fmt.Println(level, *cnt)
	// }
	for _, next := range rr[current] {
		_, ex := possible[next]

		if next == ignore || !ex {
			continue
		}

		if next == finish {
			*cnt += 1
		} else {
			deepCntXY(rr, next, finish, ignore, possible, cnt, level+1)
		}
	}
}

func load(fn string) (map[string][]string, map[string][]string) {
	lines := common.LoadStrings(fn)

	res := map[string][]string{}
	backwardRes := map[string][]string{}

	for _, line := range lines {
		parts1 := strings.Split(line, ": ")
		parts2 := strings.Split(parts1[1], " ")

		nodeName := parts1[0]
		res[nodeName] = parts2

		for _, p := range parts2 {
			backwardRes[p] = append(backwardRes[p], nodeName)
		}
	}

	return res, backwardRes
}
