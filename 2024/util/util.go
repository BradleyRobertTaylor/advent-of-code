package util

import "strconv"

func ToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func ToFloat64(s string) float64 {
	num, _ := strconv.ParseFloat(s, 64)
	return num
}

func Map[T, V any](l []T, myFunc func(val T) V) []V {
	result := make([]V, 0, len(l))
	for _, v := range l {
		result = append(result, myFunc(v))
	}
	return result
}

func Sum(l ...int) int {
	sum := 0
	for _, v := range l {
		sum += v
	}
	return sum
}
