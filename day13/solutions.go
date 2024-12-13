package day13

import (
	"aoc2024/utils"
	"fmt"
	"regexp"
)

func Answer() {
	input := utils.GetInput(13, false)
	buttons, targets := parseInput(input)
	// Part 1
	fmt.Println(calculateMinimumCost(buttons, targets))

	// Part 2
	fmt.Println(calculateMinimumCost2(buttons, targets))
}

type Button struct {
	A utils.Pair
	B utils.Pair
}

func parseInput(input []string) ([]Button, []utils.Pair) {
	buttons := make([]Button, 0)
	targets := make([]utils.Pair, 0)
	for i := 0; i < len(input); i += 4 {
		buttonA := parseNums(input[i])
		buttonB := parseNums(input[i+1])
		buttons = append(buttons, Button{buttonA, buttonB})
		targets = append(targets, parseNums(input[i+2]))
	}

	return buttons, targets
}

func parseNums(s string) utils.Pair {
	r := regexp.MustCompile("[0-9]{1,10}")

	nums := r.FindAllString(s, -1)
	return utils.Pair{utils.ToInt(nums[0]), utils.ToInt(nums[1])}
}

func calculateMinimumCost(buttons []Button, goals []utils.Pair) int {
	sum := 0
	for i := 0; i < len(buttons); i++ {
		sum += minPush(buttons[i], goals[i])
	}
	return sum
}

func calculateMinimumCost2(buttons []Button, goals []utils.Pair) int {
	sum := 0
	for i := 0; i < len(buttons); i++ {
		sum += minPush(buttons[i], utils.Pair{
			goals[i].X + 10000000000000,
			goals[i].Y + 10000000000000,
		})
	}
	return sum
}

func minPush(button Button, pair utils.Pair) int {
	X1, Y1, X2, Y2 := button.A.X, button.A.Y, button.B.X, button.B.Y
	X, Y := pair.X, pair.Y

	y := (X*Y1 - Y*X1) / (X2*Y1 - Y2*X1)
	x := (X - y*X2) / X1

	if x*X1+y*X2 == X && x*Y1+y*Y2 == Y {
		return x*3 + y
	}
	return 0
}
