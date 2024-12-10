package day10

import (
	"aoc2024/utils"
	"fmt"
)

func Answer() {
	input := utils.GetInput(10, false)

	var heights [][]int
	for i := 0; i < len(input); i++ {
		var row = utils.ToCharIntArr(input[i])
		heights = append(heights, row)
	}
	// Part 1
	fmt.Println(trekkingScores(heights))

	// Part 2
	fmt.Println(trekkingRatings(heights))
}

func trekkingRatings(heights [][]int) int {
	sum := 0
	for x := 0; x < len(heights); x++ {
		for y := 0; y < len(heights[x]); y++ {
			if heights[x][y] == 0 {
				sum += trekkingRatingsFrom(heights, x, y)
			}
		}
	}

	return sum
}

func trekkingRatingsFrom(heights [][]int, x int, y int) int {
	var marked [][]int
	for i := 0; i < len(heights); i++ {
		var row []int
		for j := 0; j < len(heights[i]); j++ {
			row = append(row, 0)
		}
		marked = append(marked, row)
	}

	queue := []Pair{{x, y}}
	total := 0
	for len(queue) > 0 {
		p := queue[0]
		marked[p.X][p.Y] += 1
		if heights[p.X][p.Y] == 9 {
			total += 1
		}
		if p.X > 0 && marked[p.X-1][p.Y] == 0 && heights[p.X-1][p.Y] == heights[p.X][p.Y]+1 {
			queue = append(queue, Pair{p.X - 1, p.Y})
		}
		if p.X < len(heights)-1 && marked[p.X+1][p.Y] == 0 && heights[p.X+1][p.Y] == heights[p.X][p.Y]+1 {
			queue = append(queue, Pair{p.X + 1, p.Y})
		}
		if p.Y > 0 && marked[p.X][p.Y-1] == 0 && heights[p.X][p.Y-1] == heights[p.X][p.Y]+1 {
			queue = append(queue, Pair{p.X, p.Y - 1})
		}
		if p.Y < len(heights[p.X])-1 && marked[p.X][p.Y+1] == 0 && heights[p.X][p.Y+1] == heights[p.X][p.Y]+1 {
			queue = append(queue, Pair{p.X, p.Y + 1})
		}
		queue = queue[1:]
	}

	return total
}

type Pair struct {
	X int
	Y int
}

func trekkingScoresFrom(heights [][]int, x int, y int) int {
	var marked [][]int
	for i := 0; i < len(heights); i++ {
		var row []int
		for j := 0; j < len(heights[i]); j++ {
			row = append(row, 0)
		}
		marked = append(marked, row)
	}

	queue := []Pair{{x, y}}
	total := 0
	for len(queue) > 0 {
		p := queue[0]
		if marked[p.X][p.Y] == 0 {
			marked[p.X][p.Y] = 1
			if heights[p.X][p.Y] == 9 {
				total += 1
			}
			if p.X > 0 && marked[p.X-1][p.Y] == 0 && heights[p.X-1][p.Y] == heights[p.X][p.Y]+1 {
				queue = append(queue, Pair{p.X - 1, p.Y})
			}
			if p.X < len(heights)-1 && marked[p.X+1][p.Y] == 0 && heights[p.X+1][p.Y] == heights[p.X][p.Y]+1 {
				queue = append(queue, Pair{p.X + 1, p.Y})
			}
			if p.Y > 0 && marked[p.X][p.Y-1] == 0 && heights[p.X][p.Y-1] == heights[p.X][p.Y]+1 {
				queue = append(queue, Pair{p.X, p.Y - 1})
			}
			if p.Y < len(heights[p.X])-1 && marked[p.X][p.Y+1] == 0 && heights[p.X][p.Y+1] == heights[p.X][p.Y]+1 {
				queue = append(queue, Pair{p.X, p.Y + 1})
			}
		}
		queue = queue[1:]
	}

	return total
}

func trekkingScores(heights [][]int) int {
	sum := 0
	for x := 0; x < len(heights); x++ {
		for y := 0; y < len(heights[x]); y++ {
			if heights[x][y] == 0 {
				sum += trekkingScoresFrom(heights, x, y)
			}
		}
	}

	return sum
}
