package day5

import (
	"aoc2024/utils"
	"fmt"
	"slices"
	"strings"
)

func parseInput(input []string) ([][]int, [][]int) {
	var rules [][]int
	var updates [][]int
	for _, v := range input {
		if strings.Contains(v, "|") {
			rules = append(rules, utils.ToIntArr(strings.Split(v, "|")))
		} else {
			updates = append(updates, utils.ToIntArr(strings.Split(v, ",")))
		}
	}
	return rules, updates
}

func isCorrect(next map[int][]int, update []int) bool {
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			if slices.Contains(next[update[j]], update[i]) {
				return false
			}
		}
	}
	return true
}

func calculateAdjacent(rules [][]int) map[int][]int {
	next := make(map[int][]int)
	for i := 0; i < len(rules); i++ {
		u, v := rules[i][0], rules[i][1]
		next[u] = append(next[u], v)
	}
	return next
}

func reorder(numbers []int, next map[int][]int) []int {
	for ordered := false; !ordered; {
		ordered = true
		for i := 0; i < len(numbers)-1; i++ {
			for j := i + 1; j < len(numbers); j++ {
				if slices.Contains(next[numbers[j]], numbers[i]) {
					v := numbers[j]
					numbers[j] = numbers[i]
					numbers[i] = v
					ordered = false
				}
			}
		}
	}
	return numbers
}

func Answer() {
	input := utils.GetInput(5, false)
	rules, updates := parseInput(input)
	// Part 1
	next := calculateAdjacent(rules)
	sum := 0
	for i := 0; i < len(updates); i++ {
		if isCorrect(next, updates[i]) {
			sum += updates[i][len(updates[i])/2]
		}
	}
	fmt.Println(sum)
	// Part 2
	var incorrectUpdates [][]int
	for i := 0; i < len(updates); i++ {
		if !isCorrect(next, updates[i]) {
			incorrectUpdates = append(incorrectUpdates, updates[i])
		}
	}
	sum = 0
	for i := 0; i < len(incorrectUpdates); i++ {
		update := reorder(incorrectUpdates[i], next)
		sum += update[len(update)/2]
	}
	fmt.Println(sum)
}
