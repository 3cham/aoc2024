package day2

import (
	"aoc2024/utils"
	"fmt"
	"math"
	"slices"
	"strings"
)

func isInRange(level []int) bool {
	for i := 1; i < len(level); i++ {
		if math.Abs(float64(level[i]-level[i-1])) > 3.0 || math.Abs(float64(level[i]-level[i-1])) < 1.0 {
			return false
		}
	}
	return true
}

func isMonotonic(level []int) bool {
	for i := 0; i < len(level)-2; i++ {
		if level[i] == level[i+1] {
			return false
		}
		if level[i] > level[i+1] && level[i+1] <= level[i+2] {
			return false
		}
		if level[i] < level[i+1] && level[i+1] >= level[i+2] {
			return false
		}
	}
	return true
}

func Answer() {
	var input = utils.GetInput(2, false)
	safeLevels := 0

	for _, v := range input {
		var level = strings.Split(v, " ")
		if isInRange(utils.ToIntArr(level)) && isMonotonic(utils.ToIntArr(level)) {
			safeLevels += 1
		}
	}

	fmt.Println(safeLevels)

	safeLevels = 0
	for _, v := range input {
		var level = strings.Split(v, " ")
		for i := 0; i < len(level); i++ {
			modifiedLevel := slices.Delete(strings.Split(v, " "), i, i+1)
			if isInRange(utils.ToIntArr(modifiedLevel)) && isMonotonic(utils.ToIntArr(modifiedLevel)) {
				safeLevels += 1
				break
			}
		}
	}
	fmt.Println(safeLevels)
}
