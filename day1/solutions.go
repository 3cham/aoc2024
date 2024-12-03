package day1

import (
	"aoc2024/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

func Diff(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)

	var diff = 0.0
	for i := 0; i < len(a); i++ {
		diff += math.Abs(float64(a[i]) - float64(b[i]))
	}
	return int(diff)
}

func WeightedDiff(a []int, b []int) int {
	var diff = 0.0
	for i := 0; i < len(a); i++ {
		var counter = 0
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				counter++
			}
		}
		diff += float64(a[i] * counter)
	}
	return int(diff)
}

func Answer() {
	var input = utils.GetInput(1, false)
	var a []int
	var b []int

	for _, v := range input {
		var x = strings.Split(v, "   ")
		a = append(a, utils.ToInt(x[0]))
		b = append(b, utils.ToInt(x[1]))
	}

	fmt.Println(Diff(a, b))
	fmt.Println(WeightedDiff(a, b))
}
