package utils

import "strconv"

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func ToIntArr(s []string) []int {
	result := []int{}
	for _, v := range s {
		result = append(result, ToInt(v))
	}
	return result
}
