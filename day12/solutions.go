package day12

import (
	"aoc2024/utils"
	"fmt"
)

func Answer() {

	input := utils.GetInput(12, false)

	// Part 1
	fmt.Println(calculateFencePrice(input))
	// Part 2
	fmt.Println(calculateSidePrice(input))
}

func calculateFencePrice(input []string) int {
	var marked [][]int

	for i := 0; i < len(input); i++ {
		var row []int
		for j := 0; j < len(input[i]); j++ {
			row = append(row, 0)
		}
		marked = append(marked, row)
	}

	sum := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if marked[i][j] == 0 {
				sum += bfsCircumference(i, j, marked, input)
			}
		}
	}

	return sum
}

func calculateSidePrice(input []string) int {
	var marked [][]int

	for i := 0; i < len(input); i++ {
		var row []int
		for j := 0; j < len(input[i]); j++ {
			row = append(row, 0)
		}
		marked = append(marked, row)
	}

	sum := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if marked[i][j] == 0 {
				sum += bfsSide(i, j, marked, input)
			}
		}
	}

	return sum
}

func bfsSide(i int, j int, marked [][]int, input []string) int {
	marked[i][j] = 1
	queue := []utils.Pair{{i, j}}
	side := 0
	area := 0
	var sides [][]int
	for i := 0; i < len(input); i++ {
		var row []int
		for j := 0; j < len(input[i]); j++ {
			row = append(row, 1)
		}
		sides = append(sides, row)
	}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		area += 1
		if p.X > 0 {
			if input[p.X-1][p.Y] == input[p.X][p.Y] {
				if marked[p.X-1][p.Y] == 0 {
					queue = append(queue, utils.Pair{p.X - 1, p.Y})
					marked[p.X-1][p.Y] = 1
				}
			} else {
				side += checkSide(p, input, sides, 0)
			}
		} else {
			side += checkSide(p, input, sides, 0)
		}
		if p.X < len(input)-1 {
			if input[p.X+1][p.Y] == input[p.X][p.Y] {
				if marked[p.X+1][p.Y] == 0 {
					queue = append(queue, utils.Pair{p.X + 1, p.Y})
					marked[p.X+1][p.Y] = 1
				}
			} else {
				side += checkSide(p, input, sides, 2)
			}
		} else {
			side += checkSide(p, input, sides, 2)
		}
		if p.Y > 0 {
			if input[p.X][p.Y-1] == input[p.X][p.Y] {
				if marked[p.X][p.Y-1] == 0 {
					queue = append(queue, utils.Pair{p.X, p.Y - 1})
					marked[p.X][p.Y-1] = 1
				}
			} else {
				side += checkSide(p, input, sides, 1)
			}
		} else {
			side += checkSide(p, input, sides, 1)
		}
		if p.Y < len(input[0])-1 {
			if input[p.X][p.Y+1] == input[p.X][p.Y] {
				if marked[p.X][p.Y+1] == 0 {
					queue = append(queue, utils.Pair{p.X, p.Y + 1})
					marked[p.X][p.Y+1] = 1
				}
			} else {
				side += checkSide(p, input, sides, 3)
			}
		} else {
			side += checkSide(p, input, sides, 3)
		}
	}
	return area * side
}

func checkSide(p utils.Pair, input []string, sides [][]int, direction int) int {
	directionWeight := []int{2, 3, 5, 7}
	if direction == 0 || direction == 2 {
		if p.Y > 0 && input[p.X][p.Y-1] == input[p.X][p.Y] && sides[p.X][p.Y-1]%directionWeight[direction] == 0 {
			sides[p.X][p.Y] *= directionWeight[direction]
			return 0
		}
		if p.Y < len(input[p.X])-1 && input[p.X][p.Y+1] == input[p.X][p.Y] && sides[p.X][p.Y+1]%directionWeight[direction] == 0 {
			sides[p.X][p.Y] *= directionWeight[direction]
			return 0
		}
		sides[p.X][p.Y] *= directionWeight[direction]
		return 1
	}

	if p.X > 0 && input[p.X-1][p.Y] == input[p.X][p.Y] && sides[p.X-1][p.Y]%directionWeight[direction] == 0 {
		sides[p.X][p.Y] *= directionWeight[direction]
		return 0
	}
	if p.X < len(input)-1 && input[p.X+1][p.Y] == input[p.X][p.Y] && sides[p.X+1][p.Y]%directionWeight[direction] == 0 {
		sides[p.X][p.Y] *= directionWeight[direction]
		return 0
	}
	sides[p.X][p.Y] *= directionWeight[direction]
	return 1
}

func bfsCircumference(i int, j int, marked [][]int, input []string) int {
	marked[i][j] = 1
	queue := []utils.Pair{{i, j}}
	circumference := 0
	area := 0

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		area += 1
		if p.X > 0 {
			if input[p.X-1][p.Y] == input[p.X][p.Y] {
				if marked[p.X-1][p.Y] == 0 {
					queue = append(queue, utils.Pair{p.X - 1, p.Y})
					marked[p.X-1][p.Y] = 1
				}
			} else {
				circumference++
			}
		} else {
			circumference++
		}
		if p.X < len(input)-1 {
			if input[p.X+1][p.Y] == input[p.X][p.Y] {
				if marked[p.X+1][p.Y] == 0 {
					queue = append(queue, utils.Pair{p.X + 1, p.Y})
					marked[p.X+1][p.Y] = 1
				}
			} else {
				circumference++
			}
		} else {
			circumference++
		}
		if p.Y > 0 {
			if input[p.X][p.Y-1] == input[p.X][p.Y] {
				if marked[p.X][p.Y-1] == 0 {
					queue = append(queue, utils.Pair{p.X, p.Y - 1})
					marked[p.X][p.Y-1] = 1
				}
			} else {
				circumference++
			}
		} else {
			circumference++
		}
		if p.Y < len(input[0])-1 {
			if input[p.X][p.Y+1] == input[p.X][p.Y] {
				if marked[p.X][p.Y+1] == 0 {
					queue = append(queue, utils.Pair{p.X, p.Y + 1})
					marked[p.X][p.Y+1] = 1
				}
			} else {
				circumference++
			}
		} else {
			circumference++
		}
	}
	return area * circumference
}
