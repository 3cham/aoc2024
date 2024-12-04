package day4

import (
	"aoc2024/utils"
	"fmt"
)

func countXmas(input []string) int {
	counter := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			counter += isXmas(input, i, j)
		}
	}

	return counter
}

func isXmas(input []string, i int, j int) int {
	counter := 0
	if j < len(input[i])-3 && input[i][j] == 'X' && input[i][j+1] == 'M' && input[i][j+2] == 'A' && input[i][j+3] == 'S' {
		counter += 1
	}

	if j < len(input[i])-3 && input[i][j] == 'S' && input[i][j+1] == 'A' && input[i][j+2] == 'M' && input[i][j+3] == 'X' {
		counter += 1
	}

	if i < len(input)-3 && input[i][j] == 'X' && input[i+1][j] == 'M' && input[i+2][j] == 'A' && input[i+3][j] == 'S' {
		counter += 1
	}

	if i < len(input)-3 && input[i][j] == 'S' && input[i+1][j] == 'A' && input[i+2][j] == 'M' && input[i+3][j] == 'X' {
		counter += 1
	}

	if i < len(input)-3 && j < len(input[i])-3 && input[i][j] == 'X' && input[i+1][j+1] == 'M' && input[i+2][j+2] == 'A' && input[i+3][j+3] == 'S' {
		counter += 1
	}

	if i < len(input)-3 && j < len(input[i])-3 && input[i][j] == 'S' && input[i+1][j+1] == 'A' && input[i+2][j+2] == 'M' && input[i+3][j+3] == 'X' {
		counter += 1
	}

	if i < len(input)-3 && j >= 3 && input[i][j] == 'X' && input[i+1][j-1] == 'M' && input[i+2][j-2] == 'A' && input[i+3][j-3] == 'S' {
		counter += 1
	}

	if i < len(input)-3 && j >= 3 && input[i][j] == 'S' && input[i+1][j-1] == 'A' && input[i+2][j-2] == 'M' && input[i+3][j-3] == 'X' {
		counter += 1
	}

	return counter
}

func countXXmas(input []string) int {
	counter := 0
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			counter += isXXmas(input, i, j)
		}
	}

	return counter
}

func isXXmas(input []string, i int, j int) int {
	if input[i][j] != 'A' {
		return 0
	}

	if input[i-1][j-1] == 'M' && input[i-1][j+1] == 'M' && input[i+1][j-1] == 'S' && input[i+1][j+1] == 'S' {
		return 1
	}
	if input[i-1][j-1] == 'M' && input[i-1][j+1] == 'S' && input[i+1][j-1] == 'M' && input[i+1][j+1] == 'S' {
		return 1
	}
	if input[i-1][j-1] == 'S' && input[i-1][j+1] == 'S' && input[i+1][j-1] == 'M' && input[i+1][j+1] == 'M' {
		return 1
	}
	if input[i-1][j-1] == 'S' && input[i-1][j+1] == 'M' && input[i+1][j-1] == 'S' && input[i+1][j+1] == 'M' {
		return 1
	}
	return 0
}

func Answer() {
	input := utils.GetInput(4, false)

	fmt.Println(countXmas(input))

	fmt.Println(countXXmas(input))
}
