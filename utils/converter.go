package utils

import "strconv"

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func ToIntArr(s []string) []int {
	var result []int
	for _, v := range s {
		result = append(result, ToInt(v))
	}
	return result
}

func ToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func ToInt64Arr(s []string) []int64 {
	var result []int64
	for _, v := range s {
		result = append(result, ToInt64(v))
	}
	return result
}
