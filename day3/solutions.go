package day3

import (
	"aoc2024/utils"
	"fmt"
	"regexp"
)

func extractMulOps(input string) []string {
	r := regexp.MustCompile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")

	return r.FindAllString(input, -1)
}

func extractNums(ops string) []string {
	r := regexp.MustCompile("[0-9]{1,3}")

	return r.FindAllString(ops, -1)
}

func extractMulDoDontOps(input string) []string {
	r := regexp.MustCompile("mul\\([0-9]{1,3},[0-9]{1,3}\\)|do\\(\\)|don't\\(\\)")

	return r.FindAllString(input, -1)
}

func Answer() {
	var input = utils.GetInput(3, false)
	sum := 0

	for i := 0; i < len(input); i++ {
		ops := extractMulOps(input[i])

		for _, op := range ops {
			nums := extractNums(op)
			sum += utils.ToInt(nums[0]) * utils.ToInt(nums[1])
		}
	}
	fmt.Println(sum)

	sum = 0
	enabled := true
	for i := 0; i < len(input); i++ {
		ops := extractMulDoDontOps(input[i])

		for _, op := range ops {
			switch op {
			case "do()":
				enabled = true
			case "don't()":
				enabled = false
			default:
				nums := extractNums(op)
				if enabled {
					sum += utils.ToInt(nums[0]) * utils.ToInt(nums[1])
				}
			}

		}
	}
	fmt.Println(sum)
}
