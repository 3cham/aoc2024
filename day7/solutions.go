package day7

import (
	"aoc2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func Answer() {
	input := utils.GetInput(7, false)
	tests, operators := parseInput(input)
	// Part 1
	sum := int64(0)
	for i := 0; i < len(tests); i++ {
		sum += tests[i] * check(tests[i], operators[i])
	}
	fmt.Println(sum)

	// Part 2
	sum = int64(0)
	for i := 0; i < len(tests); i++ {
		sum += tests[i] * checkWithCombineDigit(tests[i], operators[i])
	}
	fmt.Println(sum)
}

func check(currentValue int64, ops []int64) int64 {
	if len(ops) == 1 && currentValue == ops[0] {
		return 1
	}

	if len(ops) == 1 && currentValue != ops[0] {
		return 0
	}

	lastOp := ops[len(ops)-1]

	case1 := check(currentValue-lastOp, ops[:len(ops)-1])
	case2 := int64(0)

	if currentValue%lastOp == 0 {
		case2 = check(currentValue/lastOp, ops[:len(ops)-1])
	}

	return max(case1, case2)
}

func checkWithCombineDigit(currentValue int64, ops []int64) int64 {
	if len(ops) == 1 && currentValue == ops[0] {
		return 1
	}

	if len(ops) == 1 && currentValue != ops[0] {
		return 0
	}

	lastOp := ops[len(ops)-1]

	case1 := checkWithCombineDigit(currentValue-lastOp, ops[:len(ops)-1])
	case2 := int64(0)

	if currentValue%lastOp == 0 {
		case2 = checkWithCombineDigit(currentValue/lastOp, ops[:len(ops)-1])
	}

	case3 := int64(0)

	if prefix, ok := checkSuffix(currentValue, lastOp); ok {
		case3 = checkWithCombineDigit(prefix, ops[:len(ops)-1])
	}
	return max(case1, case2, case3)
}

func checkSuffix(value int64, op int64) (int64, bool) {
	s := strconv.FormatInt(value, 10)
	o := strconv.FormatInt(op, 10)

	ns := utils.ToInt64(strings.TrimSuffix(s, o))
	return ns, ns < value
}

func parseInput(input []string) ([]int64, [][]int64) {
	var tests []int64
	var operators [][]int64
	for _, v := range input {
		row := strings.Split(v, ":")
		tests = append(tests, utils.ToInt64(row[0]))
		operators = append(operators, utils.ToInt64Arr(strings.Split(row[1], " ")))
	}
	return tests, operators
}
