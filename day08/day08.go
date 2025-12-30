package day08

import (
	"aoc2025/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	ID      int
	GroupID int
	X, Y, Z int
}

type Pair struct {
	P1, P2 int
	Dist   int
}

func Part1() {
	N := 1000
	points := load("./day08/input01.txt")

	pairs := []Pair{}

	for i := 0; i <= len(points)-2; i++ {
		for j := i + 1; j <= len(points)-1; j++ {
			pairs = append(pairs, Pair{
				P1:   i,
				P2:   j,
				Dist: (points[j].X-points[i].X)*(points[j].X-points[i].X) + (points[j].Y-points[i].Y)*(points[j].Y-points[i].Y) + (points[j].Z-points[i].Z)*(points[j].Z-points[i].Z),
			})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Dist < pairs[j].Dist
	})

	groups := map[int][]int{}
	gIndex := 1

	for index := 0; index < N; index++ {
		pair := pairs[index]
		_, ex1 := groups[points[pair.P1].GroupID]
		_, ex2 := groups[points[pair.P2].GroupID]

		// fmt.Printf("%d[%d] %d[%d] %d\n", pair.P1, points[pair.P1].GroupID, pair.P2, points[pair.P2].GroupID, pair.Dist)
		// fmt.Println(ex1, ex2)

		if !ex1 && !ex2 {
			groups[gIndex] = []int{pair.P1, pair.P2}
			points[pair.P1].GroupID = gIndex
			points[pair.P2].GroupID = gIndex
			gIndex++
		}

		if !ex1 && ex2 {
			points[pair.P1].GroupID = points[pair.P2].GroupID
			gKey := points[pair.P2].GroupID
			groups[gKey] = append(groups[gKey], pairs[index].P1)
		}

		if ex1 && !ex2 {
			points[pair.P2].GroupID = points[pair.P1].GroupID
			gKey := points[pair.P1].GroupID
			groups[gKey] = append(groups[gKey], pairs[index].P2)
		}

		if ex1 && ex2 && points[pair.P1].GroupID != points[pair.P2].GroupID {
			gKey := points[pair.P1].GroupID
			deleteKey := points[pair.P2].GroupID

			for _, pIndex := range groups[deleteKey] {
				points[pIndex].GroupID = gKey
			}

			groups[gKey] = append(groups[gKey], groups[deleteKey]...)
			delete(groups, deleteKey)
		}

		// for k, v := range groups {
		// 	fmt.Println(k, v)
		// }
		// fmt.Println()
	}

	// for k, v := range groups {
	// 	fmt.Println(k, v)
	// }

	lengs := []int{}
	for _, v := range groups {
		lengs = append(lengs, len(v))
	}

	sort.Slice(lengs, func(i, j int) bool { return lengs[j] < lengs[i] })

	fmt.Println("Part1:", lengs[0]*lengs[1]*lengs[2])
}

func Part2() {
	points := load("./day08/input01.txt")

	pairs := []Pair{}

	for i := 0; i <= len(points)-2; i++ {
		for j := i + 1; j <= len(points)-1; j++ {
			pairs = append(pairs, Pair{
				P1:   i,
				P2:   j,
				Dist: (points[j].X-points[i].X)*(points[j].X-points[i].X) + (points[j].Y-points[i].Y)*(points[j].Y-points[i].Y) + (points[j].Z-points[i].Z)*(points[j].Z-points[i].Z),
			})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Dist < pairs[j].Dist
	})

	groups := map[int][]int{}
	gIndex := 1
	index := 0

	for {
		pair := pairs[index]
		_, ex1 := groups[points[pair.P1].GroupID]
		_, ex2 := groups[points[pair.P2].GroupID]

		if !ex1 && !ex2 {
			groups[gIndex] = []int{pair.P1, pair.P2}
			points[pair.P1].GroupID = gIndex
			points[pair.P2].GroupID = gIndex
			gIndex++
		}

		if !ex1 && ex2 {
			points[pair.P1].GroupID = points[pair.P2].GroupID
			gKey := points[pair.P2].GroupID
			groups[gKey] = append(groups[gKey], pairs[index].P1)
		}

		if ex1 && !ex2 {
			points[pair.P2].GroupID = points[pair.P1].GroupID
			gKey := points[pair.P1].GroupID
			groups[gKey] = append(groups[gKey], pairs[index].P2)
		}

		if ex1 && ex2 && points[pair.P1].GroupID != points[pair.P2].GroupID {
			gKey := points[pair.P1].GroupID
			deleteKey := points[pair.P2].GroupID

			for _, pIndex := range groups[deleteKey] {
				points[pIndex].GroupID = gKey
			}

			groups[gKey] = append(groups[gKey], groups[deleteKey]...)
			delete(groups, deleteKey)
		}

		allHasSameGroup := true

		for _, p := range points {
			if p.GroupID == 0 || points[0].GroupID != p.GroupID {
				allHasSameGroup = false
				break
			}
		}

		// fmt.Println(groups)

		if allHasSameGroup {
			// fmt.Println(pair)
			// fmt.Println(points[pair.P1], points[pair.P2])
			fmt.Println("Part2:", points[pair.P1].X*points[pair.P2].X)
			break
		}
		index++
	}
}

func load(fn string) []Point {
	lines := common.LoadStrings(fn)

	points := []Point{}
	for i, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		points = append(points, Point{ID: i, X: x, Y: y, Z: z})
	}

	return points
}
