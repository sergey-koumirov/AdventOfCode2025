package main

import (
	"aoc2025/day01"
	"aoc2025/day02"
	"aoc2025/day03"
	"aoc2025/day04"
	"aoc2025/day05"
	"aoc2025/day06"
	"aoc2025/day07"
	"aoc2025/day08"
	"aoc2025/day09"
	"aoc2025/day10"
	"aoc2025/day11"
	"aoc2025/day12"
	"flag"
	"fmt"
)

// go run app.go --day=8
func main() {
	day := flag.Int("day", 1, "Day")

	flag.Parse()

	fmt.Println("Day:", *day)

	switch *day {
	case 1:
		day01.Part1()
		day01.Part2()
	case 2:
		day02.Part1()
		day02.Part2()
	case 3:
		day03.Part1()
		day03.Part2()
	case 4:
		day04.Part1()
		day04.Part2()
	case 5:
		day05.Part1()
		day05.Part2()
	case 6:
		day06.Part1()
		day06.Part2()
	case 7:
		day07.Part1()
		day07.Part2()
	case 8:
		day08.Part1()
		day08.Part2()
	case 9:
		day09.Part1()
		day09.Part2()
	case 10:
		day10.Part1()
		day10.Part2()
	case 11:
		day11.Part1()
		day11.Part2()
	case 12:
		day12.Part1()
		day12.Part2()
	default:
		fmt.Println("???")
	}

}
