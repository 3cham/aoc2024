package day19

import (
	"aoc2024/utils"
	"fmt"
	"strings"
)

func Answer() {
	input := utils.GetInput(19, false)
	patterns, designs := parseInput(input)

	// Part 1
	fmt.Println(countPossibleDesign(patterns, designs))
}

func countPossibleDesign(patterns []string, designs []string) (int, int) {
	counter := make(map[string]int)
	count := 0
	total := 0
	for _, design := range designs {
		if numPossibleDesign(patterns, design, counter) > 0 {
			count++
			total += counter[design]
		}
	}
	return count, total
}

func numPossibleDesign(patterns []string, design string, counter map[string]int) int {
	if len(design) == 0 {
		return 1
	}
	if v, ok := counter[design]; ok {
		return v
	}

	result := 0
	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			result += numPossibleDesign(patterns, design[len(pattern):], counter)
		}
	}
	counter[design] = result
	return result
}

func parseInput(input []string) ([]string, []string) {
	patterns := strings.Split(input[0], ", ")
	designs := input[2:]

	return patterns, designs
}
