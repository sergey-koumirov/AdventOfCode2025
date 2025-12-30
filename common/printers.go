package common

import "fmt"

func PrintMap(m [][]rune) {
	for _, row := range m {
		for _, ch := range row {
			fmt.Print(string(ch))
		}
		fmt.Println()
	}
}

func PrintStrings(m []string) {
	for _, row := range m {
		fmt.Println(row)
	}
}
