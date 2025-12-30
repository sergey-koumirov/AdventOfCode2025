package day10

import (
	"aoc2025/common"
	"bytes"
	"fmt"
	"math"
	"strings"
)

type Record struct {
	Pattern    string
	Buttons    [][]int
	ButtonMaxs []int
	Jolts      []int
	Matrix     [][]float64
	RowOrder   []int
	ColOrder   []int
}

type Variant struct {
	Current, Max, Index int
}

const VERBOSE = false

func Part1() {
	rr := load("./day10/input01.txt")

	sum := 0

	for _, r := range rr {
		maxLevel := len(r.Buttons)
		for cnt := 1; cnt <= maxLevel; cnt++ {
			acc := make([][]int, cnt)
			found, acc := deepFor(0, 0, cnt-1, r, acc)
			if found {
				sum += len(acc)
				break
			}
		}
	}

	fmt.Println("Part1:", sum)
}

func deepFor(level, fromIndex, max int, r Record, acc [][]int) (bool, [][]int) {
	buttonsCnt := len(r.Buttons)

	for index := fromIndex; index < buttonsCnt; index++ {
		acc[level] = r.Buttons[index]
		if level == max {
			pattern := apply(len(r.Pattern), acc)
			if r.Pattern == pattern {
				// fmt.Println(r.Pattern, pattern, acc)
				return true, acc
			}
		} else {
			found, acc := deepFor(level+1, index+1, max, r, acc)
			if found {
				return true, acc
			}
		}
	}
	return false, [][]int{}
}

func apply(lng int, acc [][]int) string {
	res := bytes.Repeat([]byte{'.'}, lng)

	for _, ii := range acc {
		for _, index := range ii {
			if res[index] == '.' {
				res[index] = '#'
			} else {
				res[index] = '.'
			}
		}
	}

	return string(res)
}

func Part2() {
	rr := load("./day10/input01.txt")

	cnt := 0
	for _, r := range rr {
		// PrintRecord(r)
		if VERBOSE {
			PrintMatrix(r)
		}

		MatrixNorm(&r)

		s := CalcMinSolution(&r)
		cnt += s
		// fmt.Println(s)

		// PrintMatrix(r)

		// fmt.Println("-----------------")
	}

	fmt.Println("Part2:", cnt)
}

func CalcMinSolution(r *Record) int {
	rcMin := common.Min(len(r.Matrix), len(r.Buttons)) - 1

	variants := []Variant{}

	for col := rcMin + 1; col < len(r.Buttons); col++ {
		realCol := r.ColOrder[col]
		variants = append(variants, Variant{Current: 0, Max: r.ButtonMaxs[realCol], Index: col})
	}

	maxVariant := len(variants) - 1

	solution := -1.0

	for {
		answers := make([]float64, rcMin+1)
		isWrong := false
		for rc := rcMin; rc >= 0; rc-- {
			expected := r.Matrix[rc][len(r.Matrix[rc])-1]
			sum := 0.0
			for _, v := range variants {
				sum += r.Matrix[rc][v.Index] * float64(v.Current)
			}

			for col := rc + 1; col <= rcMin; col++ {
				sum += r.Matrix[rc][col] * answers[col]
			}

			answers[rc] = (expected - sum) / r.Matrix[rc][rc]
			if common.IsNegative(answers[rc]) || !common.IsInt(answers[rc]) {
				isWrong = true
				break
			}
		}

		if !isWrong {
			// fmt.Print(answers)
			cnt := 0.0
			for _, a := range answers {
				cnt += a
			}

			for _, v := range variants {
				// fmt.Print(v.Current, " ")
				cnt += float64(v.Current)
			}
			// fmt.Printf("[%0.2f]\n", cnt)

			if solution == -1.0 || solution > math.Round(cnt) {
				solution = math.Round(cnt)
			}
		}

		if len(variants) == 0 {
			break
		}

		variants[0].Current += 1
		for i := 0; i < maxVariant; i++ {
			if variants[i].Current > variants[i].Max {
				variants[i].Current = 0
				variants[i+1].Current += 1
			}
		}

		if variants[maxVariant].Current > variants[maxVariant].Max {
			break
		}
	}

	return int(solution)
}

func MatrixNorm(r *Record) {
	rcIndex := 0

	for {
		rcMin := common.Min(len(r.Matrix), len(r.Buttons)) - 1

		if rcIndex >= rcMin {
			if rcIndex == rcMin && common.IsZero(r.Matrix[rcIndex][rcIndex]) {
				// fmt.Println("IsZero", r.Matrix[rcIndex][rcIndex])
				MatrixSwapCols(r, rcIndex)
				if VERBOSE {
					fmt.Println("Last Zero")
					PrintMatrix(*r)
				}
			}

			if rcIndex == rcMin && r.Matrix[rcIndex][rcIndex] != 1 {
				// fmt.Println(r.Matrix[rcIndex])
				// fmt.Println(r.Matrix[rcIndex][rcIndex])

				// PrintMatrix(*r)/

				MatrixNormRow(r, rcIndex)
			}

			break
		}

		if common.IsZero(r.Matrix[rcIndex][rcIndex]) {
			firstOneIndex := -1

			for i := rcIndex + 1; i <= rcMin; i++ {
				if r.Matrix[i][rcIndex] != 0 {
					firstOneIndex = i
					break
				}
			}

			if firstOneIndex > -1 {
				MatrixSwap(*r, firstOneIndex, rcIndex)
			} else {
				MatrixSwapCols(r, rcIndex)
				if VERBOSE {
					fmt.Println("No Swap")
					PrintMatrix(*r)
				}
			}
		}

		MatrixNormRow(r, rcIndex)
		if VERBOSE {
			PrintMatrix(*r)
		}

		rcIndex++
	}
}

func MatrixSwapCols(r *Record, rcIndex int) {
	swapIndex := -1

	for col := rcIndex + 1; col < len(r.Buttons); col++ {
		if !common.IsZero(r.Matrix[rcIndex][col]) {
			swapIndex = col
			break
		}
	}

	if swapIndex > -1 {
		tempIndex := r.ColOrder[rcIndex]
		r.ColOrder[rcIndex] = r.ColOrder[swapIndex]
		r.ColOrder[swapIndex] = tempIndex

		for i := 0; i < len(r.Matrix); i++ {
			temp := r.Matrix[i][swapIndex]
			r.Matrix[i][swapIndex] = r.Matrix[i][rcIndex]
			r.Matrix[i][rcIndex] = temp
		}
	}
}

func MatrixSwap(r Record, rowIndex1, rowIndex2 int) {
	// fmt.Println("Swap", rowIndex1, "<->", rowIndex2)

	t1 := r.RowOrder[rowIndex1]
	r.RowOrder[rowIndex1] = r.RowOrder[rowIndex2]
	r.RowOrder[rowIndex2] = t1

	temp := r.Matrix[rowIndex2]
	r.Matrix[rowIndex2] = r.Matrix[rowIndex1]
	r.Matrix[rowIndex1] = temp
}

func MatrixNormRow(r *Record, rcIndex int) {
	// fmt.Println("Norm", rowIndex, colIndex)

	if common.IsZero(r.Matrix[rcIndex][rcIndex]) {
		fmt.Println("MatrixNormRow Zero")
		MatrixSwapCols(r, rcIndex)
		return
	}

	result := [][]float64{}

	k0 := 1 / r.Matrix[rcIndex][rcIndex]

	for rIndex := 0; rIndex < rcIndex; rIndex++ {
		result = append(result, r.Matrix[rIndex])
	}

	for rIndex := rcIndex; rIndex < len(r.Matrix); rIndex++ {
		k := r.Matrix[rIndex][rcIndex] / r.Matrix[rcIndex][rcIndex]

		allZeros := true
		for cIndex := range r.Matrix[rIndex] {
			if rIndex == rcIndex {
				r.Matrix[rIndex][cIndex] = r.Matrix[rIndex][cIndex] * k0
			} else {
				r.Matrix[rIndex][cIndex] = r.Matrix[rIndex][cIndex] - r.Matrix[rcIndex][cIndex]*k
			}

			if r.Matrix[rIndex][cIndex] != 0 {
				allZeros = false
			}
		}

		if !allZeros {
			result = append(result, r.Matrix[rIndex])
		}
	}

	r.Matrix = result
}

func PrintRecord(r Record) {
	// fmt.Println(r)

	fmt.Print("    ")
	for i := range r.Buttons {
		fmt.Printf("%4s", fmt.Sprintf("B%d", i))
	}
	fmt.Println()

	fmt.Print("     ")
	for _, m := range r.ButtonMaxs {
		fmt.Printf("%4s", fmt.Sprintf("(%d)", m))
	}
	fmt.Println()

	for jIndex, j := range r.Jolts {
		fmt.Printf("%4s", fmt.Sprintf("J%d", jIndex))

		for _, bb := range r.Buttons {
			if common.IndexOf(bb, jIndex) == -1 {
				fmt.Print("   .")
			} else {
				fmt.Print("   1")
			}
		}
		fmt.Printf("%6s", fmt.Sprintf("=%d", j))
		fmt.Println()
	}
	fmt.Println()
}

func PrintMatrix(r Record) {
	fmt.Print("    ")
	for _, n := range r.ColOrder {
		fmt.Printf("  (%d)", n)
	}
	fmt.Println()

	for jIndex := range r.Matrix {
		fmt.Printf("  J%d", r.RowOrder[jIndex])
		for i, n := range r.Matrix[jIndex] {
			if i == len(r.Matrix[jIndex])-1 {
				if n == 0 {
					fmt.Print("   0.0")
				} else {
					fmt.Printf("%6.1f", n)
				}
			} else {
				if n == 0 {
					fmt.Print("  0.0")
				} else {
					fmt.Printf("%5.1f", n)
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func load(fn string) []Record {
	lines := common.LoadStrings(fn)

	res := []Record{}

	for _, line := range lines {
		temp := Record{Buttons: [][]int{}}
		index1 := strings.Index(line, "]")
		temp.Pattern = line[1:index1]

		index2 := strings.Index(line, "{")
		str1 := line[index2+1 : len(line)-1]

		temp.Jolts = common.StrToInts(str1)

		str2 := line[index1+3 : index2-2]
		parts := strings.Split(str2, ") (")

		for _, part := range parts {
			temp.Buttons = append(temp.Buttons, common.StrToInts(part))
		}

		temp.Matrix = [][]float64{}
		temp.RowOrder = make([]int, len(temp.Jolts))

		for range temp.Jolts {
			temp.Matrix = append(temp.Matrix, []float64{})
		}

		for jIndex, jolt := range temp.Jolts {
			for _, bb := range temp.Buttons {
				if common.IndexOf(bb, jIndex) == -1 {
					temp.Matrix[jIndex] = append(temp.Matrix[jIndex], 0)
				} else {
					temp.Matrix[jIndex] = append(temp.Matrix[jIndex], 1)
				}
			}

			temp.Matrix[jIndex] = append(temp.Matrix[jIndex], float64(jolt))
			temp.RowOrder[jIndex] = jIndex
		}

		temp.ButtonMaxs = make([]int, len(temp.Buttons))

		for bi, jj := range temp.Buttons {
			min := -1

			for _, ji := range jj {
				if min == -1 || min > temp.Jolts[ji] {
					min = temp.Jolts[ji]
				}
			}

			temp.ButtonMaxs[bi] = min
		}

		temp.ColOrder = make([]int, len(temp.Buttons))
		for i := range temp.ColOrder {
			temp.ColOrder[i] = i
		}

		res = append(res, temp)
	}

	return res
}
