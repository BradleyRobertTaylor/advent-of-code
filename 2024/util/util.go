package util

import "strconv"

func Num(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
