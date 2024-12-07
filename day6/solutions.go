package day6

import (
	"aoc2024/utils"
	"fmt"
)

func move(input []string, x int, y int, direction byte) (int, int, byte, bool) {
	if direction == '^' {
		if x == 0 {
			return x, y, direction, true
		} else {
			if input[x-1][y] == '#' {
				return x, y, '>', false
			} else {
				return x - 1, y, direction, false
			}
		}
	} else if direction == 'v' {
		if x == len(input)-1 {
			return x, y, direction, true
		} else {
			if input[x+1][y] == '#' {
				return x, y, '<', false
			} else {
				return x + 1, y, direction, false
			}
		}
	} else if direction == '<' {
		if y == 0 {
			return x, y, direction, true
		} else {
			if input[x][y-1] == '#' {
				return x, y, '^', false
			} else {
				return x, y - 1, direction, false
			}
		}
	} else if direction == '>' {
		if y == len(input[x])-1 {
			return x, y, direction, true
		} else {
			if input[x][y+1] == '#' {
				return x, y, 'v', false
			} else {
				return x, y + 1, direction, false
			}
		}
	}
	return -1, -1, 0, true
}

func countSteps(input []string) int {
	steps := 0
	var marked [][]bool
	for i := 0; i < len(input); i++ {
		var row []bool
		for j := 0; j < len(input[i]); j++ {
			row = append(row, false)
		}
		marked = append(marked, row)
	}
	px, py, direction := findStart(input)

	for stop := false; !stop; {
		if !marked[px][py] {
			marked[px][py] = true
			steps++
		}

		px, py, direction, stop = move(input, px, py, direction)
	}
	return steps
}

func findStart(input []string) (int, int, byte) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == '^' || input[i][j] == 'v' || input[i][j] == '<' || input[i][j] == '>' {
				return i, j, input[i][j]
			}
		}
	}
	return -1, -1, 0
}

func countStepsWithObstacle(input []string) int {
	var marked [][]map[byte]bool
	for i := 0; i < len(input); i++ {
		var row []map[byte]bool
		for j := 0; j < len(input[i]); j++ {
			row = append(row, make(map[byte]bool))
		}
		marked = append(marked, row)
	}
	px, py, direction := findStart(input)

	for stop := false; !stop; {
		visited, _ := marked[px][py][direction]
		if visited {
			return 1
		}
		marked[px][py][direction] = true
		px, py, direction, stop = move(input, px, py, direction)
	}
	return 0
}

func countPossibleObstacles(input []string) int {
	obstacles := 0
	px, py, _ := findStart(input)
	var newInput []string
	for i := 0; i < len(input); i++ {
		newInput = append(newInput, input[i])
	}
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == '.' && (i != px || j != py) {
				newInput[i] = newInput[i][:j] + "#" + newInput[i][j+1:]
				obstacles += countStepsWithObstacle(newInput)
				newInput[i] = input[i]
			}
		}
	}
	return obstacles
}

func Answer() {
	input := utils.GetInput(6, false)

	// part 1
	fmt.Println(countSteps(input))

	// part 2
	fmt.Println(countPossibleObstacles(input))
}
