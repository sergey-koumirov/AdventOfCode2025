package day09

import (
	"aoc2025/common"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	Row, Col int
}

type Pair struct {
	P1, P2                            Point
	S, minRow, maxRow, minCol, maxCol int
}

func Part1() {
	points := load("./day09/input01.txt")

	max := 0

	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			temp := (common.Max(points[i].Row, points[j].Row) - common.Min(points[i].Row, points[j].Row) + 1) *
				(common.Max(points[i].Col, points[j].Col) - common.Min(points[i].Col, points[j].Col) + 1)

			if temp > max {
				max = temp
			}
		}
	}

	fmt.Println("Part1:", max)
}

func Part2() {
	points := load("./day09/input01.txt")

	pairsSet1 := []Pair{}
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			temp := (common.Max(points[i].Row, points[j].Row) - common.Min(points[i].Row, points[j].Row) + 1) *
				(common.Max(points[i].Col, points[j].Col) - common.Min(points[i].Col, points[j].Col) + 1)

			pairsSet1 = append(
				pairsSet1,
				Pair{
					P1:     points[i],
					P2:     points[j],
					minRow: common.Min(points[i].Row, points[j].Row),
					maxRow: common.Max(points[i].Row, points[j].Row),
					minCol: common.Min(points[i].Col, points[j].Col),
					maxCol: common.Max(points[i].Col, points[j].Col),
					S:      temp,
				},
			)
		}
	}
	fmt.Println("Pairs Set 1 #", len(pairsSet1))

	pairsSet2 := []Pair{}
	for _, pair := range pairsSet1 {
		if noPointsInside(pair, points) {
			pairsSet2 = append(pairsSet2, pair)
		}
	}
	fmt.Println("Pairs Set 2 #", len(pairsSet2))

	lines := []Pair{{P1: points[0], P2: points[len(points)-1]}}
	for i := 0; i < len(points)-1; i++ {
		lines = append(lines, Pair{P1: points[i], P2: points[i+1]})
	}

	pairsSet3 := []Pair{}
	for _, pair := range pairsSet2 {
		if noLineInside(pair, lines) {
			pairsSet3 = append(pairsSet3, pair)
		}
	}

	fmt.Println("Pairs Set 3 #", len(pairsSet3))

	maxS := 0
	for _, pair := range pairsSet3 {
		// fmt.Println(pair.S)
		if pair.S > maxS {
			maxS = pair.S
		}
	}

	fmt.Println("Part2:", maxS)
}

func noLineInside(pair Pair, lines []Pair) bool {
	res := true

	// fmt.Println("-----")
	// fmt.Println(pair)

	for _, l := range lines {
		// fmt.Println(l)
		if l.P1.Row == l.P2.Row && l.P1.Row != pair.maxRow && l.P1.Row != pair.minRow {

			horLineOutside := l.P1.Col <= pair.minCol && l.P2.Col <= pair.minCol || pair.maxCol <= l.P1.Col && pair.maxCol <= l.P2.Col
			rowInside := pair.minRow < l.P1.Row && l.P1.Row < pair.maxRow

			// fmt.Println(horLineOutside, rowInside)
			if !horLineOutside && rowInside {
				res = false
				break
			}
		}

		if l.P1.Col == l.P2.Col && l.P1.Col != pair.maxCol && l.P1.Col != pair.minCol {
			verLineOutside := l.P1.Row <= pair.minRow && l.P2.Row <= pair.minRow || pair.maxRow <= l.P1.Row && pair.maxRow <= l.P2.Row
			colInside := pair.minCol < l.P1.Col && l.P1.Col < pair.maxCol

			// fmt.Println(verLineOutside, colInside)
			if !verLineOutside && colInside {
				res = false
				break
			}
		}

	}

	return res
}

func noPointsInside(pair Pair, points []Point) bool {
	res := true

	for _, p := range points {
		if pair.minRow < p.Row && p.Row < pair.maxRow && pair.minCol < p.Col && p.Col < pair.maxCol {
			res = false
			break
		}
	}

	return res
}

func load(fn string) []Point {
	lines := common.LoadStrings(fn)

	points := []Point{}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		col, _ := strconv.Atoi(parts[0])
		row, _ := strconv.Atoi(parts[1])
		points = append(points, Point{Row: row, Col: col})
	}

	return points
}
