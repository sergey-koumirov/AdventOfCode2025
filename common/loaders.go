package common

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func LoadInts(fileName string) []int {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)

	res := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		num, _ := strconv.Atoi(line)
		res = append(res, num)
	}

	return res
}

func StrToInts(line string) []int {
	parts := strings.Split(line, ",")

	res := make([]int, 0)

	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		res = append(res, num)
	}

	return res
}

func LoadStrings(fileName string) []string {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)

	res := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	return res
}

func LoadMap(fileName string) ([][]rune, int, int) {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)

	res := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()

		temp := make([]rune, 0)

		for _, c := range line {
			temp = append(temp, c)
		}

		res = append(res, temp)
	}

	return res, len(res), len(res[0])
}
