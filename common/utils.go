package common

import "math"

func Pow10(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res = res * 10
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func IndexOf(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

const EPSILON = 1e-10

func IsInt(f float64) bool {
	return math.Abs(math.Round(f)-f) < EPSILON
}

func IsNegative(f float64) bool {
	return math.Abs(f) > EPSILON && f < 0
}

func IsZero(f float64) bool {
	return math.Abs(f) < EPSILON
}
